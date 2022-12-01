package notion

import (
	"fmt"
	"net/http"
	"notionboy/db/ent"
	"notionboy/db/ent/account"
	"notionboy/internal/pkg/config"
	"notionboy/internal/pkg/db/dao"
	"notionboy/internal/pkg/logger"
	"strings"

	"github.com/gin-gonic/gin"

	"golang.org/x/oauth2"
)

type OauthInterface interface {
	OAuthURL(state string) string
	OAuthProcess(g *gin.Context)
	OAuthCallback(g *gin.Context)
}

type oauthManager struct{}

func GetOauthManager() OauthInterface {
	return &oauthManager{}
}

func getOauthConf() *oauth2.Config {
	logger.SugaredLogger.Infof("oauthConf: %#v", config.GetConfig().NotionOauth)
	return &oauth2.Config{
		ClientID:     config.GetConfig().NotionOauth.ClientID,
		ClientSecret: config.GetConfig().NotionOauth.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  config.GetConfig().NotionOauth.AuthURL,
			TokenURL: config.GetConfig().NotionOauth.AuthToken,
		},
	}
}

func (o *oauthManager) OAuthURL(state string) string {
	// url := "https://notionboy-test.theboys.tech/notion/oauth?state=" + state
	url := fmt.Sprintf("%s/notion/oauth?state=%s", config.GetConfig().Service.URL, state)
	logger.SugaredLogger.Debugf("Visit the OAuthURL: %v", url)
	return url
}

func (o *oauthManager) OAuthProcess(g *gin.Context) {
	state := g.Query("state")
	oauthConf := getOauthConf()
	url := oauthConf.AuthCodeURL(state, oauth2.AccessTypeOffline)
	logger.SugaredLogger.Debugf("Visit the URL for the auth dialog: %v", url)
	g.Redirect(302, url)
}

func (o *oauthManager) OAuthCallback(g *gin.Context) {
	code := g.Query("code")
	state := g.Query("state")
	if code == "" || state == "" {
		logger.SugaredLogger.Error("code or state is empty")
		return
	}
	states := strings.Split(state, ":")
	userType := states[0]
	userID := strings.Join(states[1:], "")
	oauthConf := getOauthConf()
	tok, err := oauthConf.Exchange(g, code)
	logger.SugaredLogger.Debugf("tok: %#v", tok)
	if err != nil {
		logger.SugaredLogger.Errorf("oauthConf.Exchange() failed with %v, code is %s", err, code)
		return
	}

	// oAuthInfo
	token := tok.AccessToken

	databaseID, err := bindNotion(g, token)
	if err != nil {
		logger.SugaredLogger.Errorf("GetDatabaseID() failed with %v", err)
		g.Data(http.StatusOK, "text/html; charset=utf-8", []byte("绑定失败，错误详情："+err.Error()))
		return
	}

	acc := &ent.Account{
		UserID:         userID,
		UserType:       account.UserType(userType),
		AccessToken:    token,
		DatabaseID:     databaseID,
		IsLatestSchema: true,
	}

	if err := dao.SaveAccount(g, acc); err != nil {
		logger.SugaredLogger.Errorw("Save account failed", "err", err, "account", acc)
		g.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte(err.Error()))
	} else {
		g.Data(http.StatusOK, "text/html; charset=utf-8", []byte(config.MSG_BIND_SUCCESS))
	}
}
