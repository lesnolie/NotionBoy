// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: model/common.proto

package model

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on ErrorResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ErrorResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ErrorResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ErrorResponseMultiError, or
// nil if none found.
func (m *ErrorResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ErrorResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for CorrelationId

	if all {
		switch v := interface{}(m.GetError()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ErrorResponseValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ErrorResponseValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ErrorResponseValidationError{
				field:  "Error",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ErrorResponseMultiError(errors)
	}

	return nil
}

// ErrorResponseMultiError is an error wrapping multiple validation errors
// returned by ErrorResponse.ValidateAll() if the designated constraints
// aren't met.
type ErrorResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ErrorResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ErrorResponseMultiError) AllErrors() []error { return m }

// ErrorResponseValidationError is the validation error returned by
// ErrorResponse.Validate if the designated constraints aren't met.
type ErrorResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ErrorResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ErrorResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ErrorResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ErrorResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ErrorResponseValidationError) ErrorName() string { return "ErrorResponseValidationError" }

// Error satisfies the builtin error interface
func (e ErrorResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sErrorResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ErrorResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ErrorResponseValidationError{}

// Validate checks the field values on ErrorObject with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ErrorObject) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ErrorObject with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ErrorObjectMultiError, or
// nil if none found.
func (m *ErrorObject) ValidateAll() error {
	return m.validate(true)
}

func (m *ErrorObject) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	// no validation rules for Message

	if len(errors) > 0 {
		return ErrorObjectMultiError(errors)
	}

	return nil
}

// ErrorObjectMultiError is an error wrapping multiple validation errors
// returned by ErrorObject.ValidateAll() if the designated constraints aren't met.
type ErrorObjectMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ErrorObjectMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ErrorObjectMultiError) AllErrors() []error { return m }

// ErrorObjectValidationError is the validation error returned by
// ErrorObject.Validate if the designated constraints aren't met.
type ErrorObjectValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ErrorObjectValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ErrorObjectValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ErrorObjectValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ErrorObjectValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ErrorObjectValidationError) ErrorName() string { return "ErrorObjectValidationError" }

// Error satisfies the builtin error interface
func (e ErrorObjectValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sErrorObject.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ErrorObjectValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ErrorObjectValidationError{}

// Validate checks the field values on GenrateTokenResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GenrateTokenResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GenrateTokenResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GenrateTokenResponseMultiError, or nil if none found.
func (m *GenrateTokenResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GenrateTokenResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	// no validation rules for Type

	// no validation rules for Expiry

	if len(errors) > 0 {
		return GenrateTokenResponseMultiError(errors)
	}

	return nil
}

// GenrateTokenResponseMultiError is an error wrapping multiple validation
// errors returned by GenrateTokenResponse.ValidateAll() if the designated
// constraints aren't met.
type GenrateTokenResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GenrateTokenResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GenrateTokenResponseMultiError) AllErrors() []error { return m }

// GenrateTokenResponseValidationError is the validation error returned by
// GenrateTokenResponse.Validate if the designated constraints aren't met.
type GenrateTokenResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GenrateTokenResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GenrateTokenResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GenrateTokenResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GenrateTokenResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GenrateTokenResponseValidationError) ErrorName() string {
	return "GenrateTokenResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GenrateTokenResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGenrateTokenResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GenrateTokenResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GenrateTokenResponseValidationError{}

// Validate checks the field values on GenerateApiKeyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GenerateApiKeyResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GenerateApiKeyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GenerateApiKeyResponseMultiError, or nil if none found.
func (m *GenerateApiKeyResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GenerateApiKeyResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ApiKey

	if len(errors) > 0 {
		return GenerateApiKeyResponseMultiError(errors)
	}

	return nil
}

// GenerateApiKeyResponseMultiError is an error wrapping multiple validation
// errors returned by GenerateApiKeyResponse.ValidateAll() if the designated
// constraints aren't met.
type GenerateApiKeyResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GenerateApiKeyResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GenerateApiKeyResponseMultiError) AllErrors() []error { return m }

// GenerateApiKeyResponseValidationError is the validation error returned by
// GenerateApiKeyResponse.Validate if the designated constraints aren't met.
type GenerateApiKeyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GenerateApiKeyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GenerateApiKeyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GenerateApiKeyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GenerateApiKeyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GenerateApiKeyResponseValidationError) ErrorName() string {
	return "GenerateApiKeyResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GenerateApiKeyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGenerateApiKeyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GenerateApiKeyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GenerateApiKeyResponseValidationError{}
