// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/sms-service/v1/enums/sms.sms.enum.v1.proto

package smsenumv1

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

// Validate checks the field values on SmsInitEnum with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SmsInitEnum) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SmsInitEnum with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SmsInitEnumMultiError, or
// nil if none found.
func (m *SmsInitEnum) ValidateAll() error {
	return m.validate(true)
}

func (m *SmsInitEnum) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SmsInitEnumMultiError(errors)
	}

	return nil
}

// SmsInitEnumMultiError is an error wrapping multiple validation errors
// returned by SmsInitEnum.ValidateAll() if the designated constraints aren't met.
type SmsInitEnumMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SmsInitEnumMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SmsInitEnumMultiError) AllErrors() []error { return m }

// SmsInitEnumValidationError is the validation error returned by
// SmsInitEnum.Validate if the designated constraints aren't met.
type SmsInitEnumValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SmsInitEnumValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SmsInitEnumValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SmsInitEnumValidationError) Cause() error { return e.cause }

func (e SmsInitEnumValidationError) Code() int64 { return 10000000 }

func (e SmsInitEnumValidationError) HttpCode() int64 { return 400 }

// Key function returns key value.
func (e SmsInitEnumValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SmsInitEnumValidationError) ErrorName() string { return "SmsInitEnumValidationError" }

// Error satisfies the builtin error interface
func (e SmsInitEnumValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSmsInitEnum.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SmsInitEnumValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Code() int64
	HttpCode() int64
	Cause() error
	ErrorName() string
} = SmsInitEnumValidationError{}