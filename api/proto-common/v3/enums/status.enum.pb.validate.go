// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/proto-common/v3/enums/status.enum.proto

package common_enums

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

// Validate checks the field values on StatusCodeEnum with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StatusCodeEnum) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StatusCodeEnum with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StatusCodeEnumMultiError,
// or nil if none found.
func (m *StatusCodeEnum) ValidateAll() error {
	return m.validate(true)
}

func (m *StatusCodeEnum) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return StatusCodeEnumMultiError(errors)
	}

	return nil
}

// StatusCodeEnumMultiError is an error wrapping multiple validation errors
// returned by StatusCodeEnum.ValidateAll() if the designated constraints
// aren't met.
type StatusCodeEnumMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StatusCodeEnumMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StatusCodeEnumMultiError) AllErrors() []error { return m }

// StatusCodeEnumValidationError is the validation error returned by
// StatusCodeEnum.Validate if the designated constraints aren't met.
type StatusCodeEnumValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StatusCodeEnumValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StatusCodeEnumValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StatusCodeEnumValidationError) Cause() error { return e.cause }

func (e StatusCodeEnumValidationError) Code() int64 { return 10000000 }

func (e StatusCodeEnumValidationError) HttpCode() int64 { return 400 }

// Key function returns key value.
func (e StatusCodeEnumValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StatusCodeEnumValidationError) ErrorName() string { return "StatusCodeEnumValidationError" }

// Error satisfies the builtin error interface
func (e StatusCodeEnumValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStatusCodeEnum.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StatusCodeEnumValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Code() int64
	HttpCode() int64
	Cause() error
	ErrorName() string
} = StatusCodeEnumValidationError{}
