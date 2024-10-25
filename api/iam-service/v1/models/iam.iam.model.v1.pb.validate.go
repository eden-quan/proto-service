// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/iam-service/v1/models/iam.iam.model.v1.proto

package iammodelv1

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

// Validate checks the field values on IamDemoModel with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *IamDemoModel) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IamDemoModel with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in IamDemoModelMultiError, or
// nil if none found.
func (m *IamDemoModel) ValidateAll() error {
	return m.validate(true)
}

func (m *IamDemoModel) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Age

	// no validation rules for Id

	// no validation rules for IsDeleted

	if all {
		switch v := interface{}(m.GetCreateTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, IamDemoModelValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IamDemoModelValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IamDemoModelValidationError{
				field:  "CreateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdateTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, IamDemoModelValidationError{
					field:  "UpdateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IamDemoModelValidationError{
					field:  "UpdateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IamDemoModelValidationError{
				field:  "UpdateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return IamDemoModelMultiError(errors)
	}

	return nil
}

// IamDemoModelMultiError is an error wrapping multiple validation errors
// returned by IamDemoModel.ValidateAll() if the designated constraints aren't met.
type IamDemoModelMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IamDemoModelMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IamDemoModelMultiError) AllErrors() []error { return m }

// IamDemoModelValidationError is the validation error returned by
// IamDemoModel.Validate if the designated constraints aren't met.
type IamDemoModelValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IamDemoModelValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IamDemoModelValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IamDemoModelValidationError) Cause() error { return e.cause }

func (e IamDemoModelValidationError) Code() int64 { return 10000000 }

func (e IamDemoModelValidationError) HttpCode() int64 { return 400 }

// Key function returns key value.
func (e IamDemoModelValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IamDemoModelValidationError) ErrorName() string { return "IamDemoModelValidationError" }

// Error satisfies the builtin error interface
func (e IamDemoModelValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIamDemoModel.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IamDemoModelValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Code() int64
	HttpCode() int64
	Cause() error
	ErrorName() string
} = IamDemoModelValidationError{}