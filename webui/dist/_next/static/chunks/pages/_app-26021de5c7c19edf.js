(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[888],{705:function(){},6952:function(e,t,n){"use strict";n.d(t,{F:function(){return a},f:function(){return c}});var r=n(959);let i=["light","dark"],s="(prefers-color-scheme: dark)",o="undefined"==typeof window,u=(0,r.createContext)(void 0),l={setTheme:e=>{},themes:[]},a=()=>{var e;return null!==(e=(0,r.useContext)(u))&&void 0!==e?e:l},c=e=>(0,r.useContext)(u)?r.createElement(r.Fragment,null,e.children):r.createElement(h,e),d=["light","dark"],h=({forcedTheme:e,disableTransitionOnChange:t=!1,enableSystem:n=!0,enableColorScheme:o=!0,storageKey:l="theme",themes:a=d,defaultTheme:c=n?"system":"light",attribute:h="data-theme",value:y,children:_,nonce:S})=>{let[g,b]=(0,r.useState)(()=>f(l,c)),[T,w]=(0,r.useState)(()=>f(l)),C=y?Object.values(y):a,F=(0,r.useCallback)(e=>{let r=e;if(!r)return;"system"===e&&n&&(r=v());let s=y?y[r]:r,u=t?p():null,l=document.documentElement;if("class"===h?(l.classList.remove(...C),s&&l.classList.add(s)):s?l.setAttribute(h,s):l.removeAttribute(h),o){let e=i.includes(c)?c:null,t=i.includes(r)?r:e;l.style.colorScheme=t}null==u||u()},[]),x=(0,r.useCallback)(e=>{b(e);try{localStorage.setItem(l,e)}catch(e){}},[e]),R=(0,r.useCallback)(t=>{let r=v(t);w(r),"system"===g&&n&&!e&&F("system")},[g,e]);(0,r.useEffect)(()=>{let e=window.matchMedia(s);return e.addListener(R),R(e),()=>e.removeListener(R)},[R]),(0,r.useEffect)(()=>{let e=e=>{e.key===l&&x(e.newValue||c)};return window.addEventListener("storage",e),()=>window.removeEventListener("storage",e)},[x]),(0,r.useEffect)(()=>{F(null!=e?e:g)},[e,g]);let k=(0,r.useMemo)(()=>({theme:g,setTheme:x,forcedTheme:e,resolvedTheme:"system"===g?T:g,themes:n?[...a,"system"]:a,systemTheme:n?T:void 0}),[g,x,e,T,n,a]);return r.createElement(u.Provider,{value:k},r.createElement(m,{forcedTheme:e,disableTransitionOnChange:t,enableSystem:n,enableColorScheme:o,storageKey:l,themes:a,defaultTheme:c,attribute:h,value:y,children:_,attrs:C,nonce:S}),_)},m=(0,r.memo)(({forcedTheme:e,storageKey:t,attribute:n,enableSystem:o,enableColorScheme:u,defaultTheme:l,value:a,attrs:c,nonce:d})=>{let h="system"===l,m="class"===n?`var d=document.documentElement,c=d.classList;c.remove(${c.map(e=>`'${e}'`).join(",")});`:`var d=document.documentElement,n='${n}',s='setAttribute';`,f=u?i.includes(l)&&l?`if(e==='light'||e==='dark'||!e)d.style.colorScheme=e||'${l}'`:"if(e==='light'||e==='dark')d.style.colorScheme=e":"",p=(e,t=!1,r=!0)=>{let s=a?a[e]:e,o=t?e+"|| ''":`'${s}'`,l="";return u&&r&&!t&&i.includes(e)&&(l+=`d.style.colorScheme = '${e}';`),"class"===n?l+=t||s?`c.add(${o})`:"null":s&&(l+=`d[s](n,${o})`),l},v=e?`!function(){${m}${p(e)}}()`:o?`!function(){try{${m}var e=localStorage.getItem('${t}');if('system'===e||(!e&&${h})){var t='${s}',m=window.matchMedia(t);if(m.media!==t||m.matches){${p("dark")}}else{${p("light")}}}else if(e){${a?`var x=${JSON.stringify(a)};`:""}${p(a?"x[e]":"e",!0)}}${h?"":"else{"+p(l,!1,!1)+"}"}${f}}catch(e){}}()`:`!function(){try{${m}var e=localStorage.getItem('${t}');if(e){${a?`var x=${JSON.stringify(a)};`:""}${p(a?"x[e]":"e",!0)}}else{${p(l,!1,!1)};}${f}}catch(t){}}();`;return r.createElement("script",{nonce:d,dangerouslySetInnerHTML:{__html:v}})},()=>!0),f=(e,t)=>{let n;if(!o){try{n=localStorage.getItem(e)||void 0}catch(e){}return n||t}},p=()=>{let e=document.createElement("style");return e.appendChild(document.createTextNode("*{-webkit-transition:none!important;-moz-transition:none!important;-o-transition:none!important;-ms-transition:none!important;transition:none!important}")),document.head.appendChild(e),()=>{window.getComputedStyle(document.body),setTimeout(()=>{document.head.removeChild(e)},1)}},v=e=>(e||(e=window.matchMedia(s)),e.matches?"dark":"light")},9462:function(e,t,n){"use strict";var r,i;e.exports=(null==(r=n.g.process)?void 0:r.env)&&"object"==typeof(null==(i=n.g.process)?void 0:i.env)?n.g.process:n(7287)},3282:function(e,t,n){(window.__NEXT_P=window.__NEXT_P||[]).push(["/_app",function(){return n(5248)}])},1137:function(e,t,n){var r=n(9462);n(705);var i=n(959),s=i&&"object"==typeof i&&"default"in i?i:{default:i};function o(e,t){for(var n=0;n<t.length;n++){var r=t[n];r.enumerable=r.enumerable||!1,r.configurable=!0,"value"in r&&(r.writable=!0),Object.defineProperty(e,r.key,r)}}var u=void 0!==r&&r.env&&!0,l=function(e){return"[object String]"===Object.prototype.toString.call(e)},a=function(){function e(e){var t=void 0===e?{}:e,n=t.name,r=void 0===n?"stylesheet":n,i=t.optimizeForSpeed,s=void 0===i?u:i;c(l(r),"`name` must be a string"),this._name=r,this._deletedRulePlaceholder="#"+r+"-deleted-rule____{}",c("boolean"==typeof s,"`optimizeForSpeed` must be a boolean"),this._optimizeForSpeed=s,this._serverSheet=void 0,this._tags=[],this._injected=!1,this._rulesCount=0;var o=document.querySelector('meta[property="csp-nonce"]');this._nonce=o?o.getAttribute("content"):null}var t,n=e.prototype;return n.setOptimizeForSpeed=function(e){c("boolean"==typeof e,"`setOptimizeForSpeed` accepts a boolean"),c(0===this._rulesCount,"optimizeForSpeed cannot be when rules have already been inserted"),this.flush(),this._optimizeForSpeed=e,this.inject()},n.isOptimizeForSpeed=function(){return this._optimizeForSpeed},n.inject=function(){var e=this;if(c(!this._injected,"sheet already injected"),this._injected=!0,this._optimizeForSpeed){this._tags[0]=this.makeStyleTag(this._name),this._optimizeForSpeed="insertRule"in this.getSheet(),this._optimizeForSpeed||(u||console.warn("StyleSheet: optimizeForSpeed mode not supported falling back to standard mode."),this.flush(),this._injected=!0);return}this._serverSheet={cssRules:[],insertRule:function(t,n){return"number"==typeof n?e._serverSheet.cssRules[n]={cssText:t}:e._serverSheet.cssRules.push({cssText:t}),n},deleteRule:function(t){e._serverSheet.cssRules[t]=null}}},n.getSheetForTag=function(e){if(e.sheet)return e.sheet;for(var t=0;t<document.styleSheets.length;t++)if(document.styleSheets[t].ownerNode===e)return document.styleSheets[t]},n.getSheet=function(){return this.getSheetForTag(this._tags[this._tags.length-1])},n.insertRule=function(e,t){if(c(l(e),"`insertRule` accepts only strings"),this._optimizeForSpeed){var n=this.getSheet();"number"!=typeof t&&(t=n.cssRules.length);try{n.insertRule(e,t)}catch(t){return u||console.warn("StyleSheet: illegal rule: \n\n"+e+"\n\nSee https://stackoverflow.com/q/20007992 for more info"),-1}}else{var r=this._tags[t];this._tags.push(this.makeStyleTag(this._name,e,r))}return this._rulesCount++},n.replaceRule=function(e,t){if(this._optimizeForSpeed){var n=this.getSheet();if(t.trim()||(t=this._deletedRulePlaceholder),!n.cssRules[e])return e;n.deleteRule(e);try{n.insertRule(t,e)}catch(r){u||console.warn("StyleSheet: illegal rule: \n\n"+t+"\n\nSee https://stackoverflow.com/q/20007992 for more info"),n.insertRule(this._deletedRulePlaceholder,e)}}else{var r=this._tags[e];c(r,"old rule at index `"+e+"` not found"),r.textContent=t}return e},n.deleteRule=function(e){if(this._optimizeForSpeed)this.replaceRule(e,"");else{var t=this._tags[e];c(t,"rule at index `"+e+"` not found"),t.parentNode.removeChild(t),this._tags[e]=null}},n.flush=function(){this._injected=!1,this._rulesCount=0,this._tags.forEach(function(e){return e&&e.parentNode.removeChild(e)}),this._tags=[]},n.cssRules=function(){var e=this;return this._tags.reduce(function(t,n){return n?t=t.concat(Array.prototype.map.call(e.getSheetForTag(n).cssRules,function(t){return t.cssText===e._deletedRulePlaceholder?null:t})):t.push(null),t},[])},n.makeStyleTag=function(e,t,n){t&&c(l(t),"makeStyleTag accepts only strings as second parameter");var r=document.createElement("style");this._nonce&&r.setAttribute("nonce",this._nonce),r.type="text/css",r.setAttribute("data-"+e,""),t&&r.appendChild(document.createTextNode(t));var i=document.head||document.getElementsByTagName("head")[0];return n?i.insertBefore(r,n):i.appendChild(r),r},o(e.prototype,[{key:"length",get:function(){return this._rulesCount}}]),t&&o(e,t),e}();function c(e,t){if(!e)throw Error("StyleSheet: "+t+".")}var d=function(e){for(var t=5381,n=e.length;n;)t=33*t^e.charCodeAt(--n);return t>>>0},h={};function m(e,t){if(!t)return"jsx-"+e;var n=String(t),r=e+n;return h[r]||(h[r]="jsx-"+d(e+"-"+n)),h[r]}function f(e,t){var n=e+t;return h[n]||(h[n]=t.replace(/__jsx-style-dynamic-selector/g,e)),h[n]}var p=function(){function e(e){var t=void 0===e?{}:e,n=t.styleSheet,r=void 0===n?null:n,i=t.optimizeForSpeed,s=void 0!==i&&i;this._sheet=r||new a({name:"styled-jsx",optimizeForSpeed:s}),this._sheet.inject(),r&&"boolean"==typeof s&&(this._sheet.setOptimizeForSpeed(s),this._optimizeForSpeed=this._sheet.isOptimizeForSpeed()),this._fromServer=void 0,this._indices={},this._instancesCounts={}}var t=e.prototype;return t.add=function(e){var t=this;void 0===this._optimizeForSpeed&&(this._optimizeForSpeed=Array.isArray(e.children),this._sheet.setOptimizeForSpeed(this._optimizeForSpeed),this._optimizeForSpeed=this._sheet.isOptimizeForSpeed()),this._fromServer||(this._fromServer=this.selectFromServer(),this._instancesCounts=Object.keys(this._fromServer).reduce(function(e,t){return e[t]=0,e},{}));var n=this.getIdAndRules(e),r=n.styleId,i=n.rules;if(r in this._instancesCounts){this._instancesCounts[r]+=1;return}var s=i.map(function(e){return t._sheet.insertRule(e)}).filter(function(e){return -1!==e});this._indices[r]=s,this._instancesCounts[r]=1},t.remove=function(e){var t=this,n=this.getIdAndRules(e).styleId;if(function(e,t){if(!e)throw Error("StyleSheetRegistry: "+t+".")}(n in this._instancesCounts,"styleId: `"+n+"` not found"),this._instancesCounts[n]-=1,this._instancesCounts[n]<1){var r=this._fromServer&&this._fromServer[n];r?(r.parentNode.removeChild(r),delete this._fromServer[n]):(this._indices[n].forEach(function(e){return t._sheet.deleteRule(e)}),delete this._indices[n]),delete this._instancesCounts[n]}},t.update=function(e,t){this.add(t),this.remove(e)},t.flush=function(){this._sheet.flush(),this._sheet.inject(),this._fromServer=void 0,this._indices={},this._instancesCounts={}},t.cssRules=function(){var e=this,t=this._fromServer?Object.keys(this._fromServer).map(function(t){return[t,e._fromServer[t]]}):[],n=this._sheet.cssRules();return t.concat(Object.keys(this._indices).map(function(t){return[t,e._indices[t].map(function(e){return n[e].cssText}).join(e._optimizeForSpeed?"":"\n")]}).filter(function(e){return Boolean(e[1])}))},t.styles=function(e){var t,n;return t=this.cssRules(),void 0===(n=e)&&(n={}),t.map(function(e){var t=e[0],r=e[1];return s.default.createElement("style",{id:"__"+t,key:"__"+t,nonce:n.nonce?n.nonce:void 0,dangerouslySetInnerHTML:{__html:r}})})},t.getIdAndRules=function(e){var t=e.children,n=e.dynamic,r=e.id;if(n){var i=m(r,n);return{styleId:i,rules:Array.isArray(t)?t.map(function(e){return f(i,e)}):[f(i,t)]}}return{styleId:m(r),rules:Array.isArray(t)?t:[t]}},t.selectFromServer=function(){return Array.prototype.slice.call(document.querySelectorAll('[id^="__jsx-"]')).reduce(function(e,t){return e[t.id.slice(2)]=t,e},{})},e}(),v=i.createContext(null);v.displayName="StyleSheetContext";var y=s.default.useInsertionEffect||s.default.useLayoutEffect,_=new p;function S(e){var t=_||i.useContext(v);return t&&y(function(){return t.add(e),function(){t.remove(e)}},[e.id,String(e.dynamic)]),null}S.dynamic=function(e){return e.map(function(e){return m(e[0],e[1])}).join(" ")},t.style=S},2311:function(e,t,n){"use strict";e.exports=n(1137).style},5248:function(e,t,n){"use strict";n.r(t),n.d(t,{default:function(){return a}});var r=n(1527),i=n(2897),s=n.n(i),o=n(2311),u=n.n(o),l=n(6952);function a(e){let{Component:t,pageProps:n}=e;return(0,r.jsxs)(r.Fragment,{children:[(0,r.jsx)(u(),{id:"508c1ed201c666fb",dynamic:[s().style.fontFamily],children:":root{--font-sans:".concat(s().style.fontFamily,"}}")}),(0,r.jsx)(l.f,{attribute:"class",defaultTheme:"system",enableSystem:!0,children:(0,r.jsx)(t,{...n,className:u().dynamic([["508c1ed201c666fb",[s().style.fontFamily]]])+" "+(n&&null!=n.className&&n.className||"")})})]})}n(1511)},1511:function(){},2897:function(e){e.exports={style:{fontFamily:"'__Inter_79ed20', '__Inter_Fallback_79ed20'",fontStyle:"normal"},className:"__className_79ed20",variable:"__variable_79ed20"}},7287:function(e){!function(){var t={229:function(e){var t,n,r,i=e.exports={};function s(){throw Error("setTimeout has not been defined")}function o(){throw Error("clearTimeout has not been defined")}function u(e){if(t===setTimeout)return setTimeout(e,0);if((t===s||!t)&&setTimeout)return t=setTimeout,setTimeout(e,0);try{return t(e,0)}catch(n){try{return t.call(null,e,0)}catch(n){return t.call(this,e,0)}}}!function(){try{t="function"==typeof setTimeout?setTimeout:s}catch(e){t=s}try{n="function"==typeof clearTimeout?clearTimeout:o}catch(e){n=o}}();var l=[],a=!1,c=-1;function d(){a&&r&&(a=!1,r.length?l=r.concat(l):c=-1,l.length&&h())}function h(){if(!a){var e=u(d);a=!0;for(var t=l.length;t;){for(r=l,l=[];++c<t;)r&&r[c].run();c=-1,t=l.length}r=null,a=!1,function(e){if(n===clearTimeout)return clearTimeout(e);if((n===o||!n)&&clearTimeout)return n=clearTimeout,clearTimeout(e);try{n(e)}catch(t){try{return n.call(null,e)}catch(t){return n.call(this,e)}}}(e)}}function m(e,t){this.fun=e,this.array=t}function f(){}i.nextTick=function(e){var t=Array(arguments.length-1);if(arguments.length>1)for(var n=1;n<arguments.length;n++)t[n-1]=arguments[n];l.push(new m(e,t)),1!==l.length||a||u(h)},m.prototype.run=function(){this.fun.apply(null,this.array)},i.title="browser",i.browser=!0,i.env={},i.argv=[],i.version="",i.versions={},i.on=f,i.addListener=f,i.once=f,i.off=f,i.removeListener=f,i.removeAllListeners=f,i.emit=f,i.prependListener=f,i.prependOnceListener=f,i.listeners=function(e){return[]},i.binding=function(e){throw Error("process.binding is not supported")},i.cwd=function(){return"/"},i.chdir=function(e){throw Error("process.chdir is not supported")},i.umask=function(){return 0}}},n={};function r(e){var i=n[e];if(void 0!==i)return i.exports;var s=n[e]={exports:{}},o=!0;try{t[e](s,s.exports,r),o=!1}finally{o&&delete n[e]}return s.exports}r.ab="//";var i=r(229);e.exports=i}()}},function(e){var t=function(t){return e(e.s=t)};e.O(0,[774,179],function(){return t(3282),t(2970)}),_N_E=e.O()}]);
