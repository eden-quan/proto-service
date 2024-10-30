// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/messagecenter-service/v1/enums/messagecenter.messagecenter.enum.v1.proto

package messagecenterenumv1

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

// Validate checks the field values on MsgStatusEnum with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MsgStatusEnum) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MsgStatusEnum with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MsgStatusEnumMultiError, or
// nil if none found.
func (m *MsgStatusEnum) ValidateAll() error {
	return m.validate(true)
}

func (m *MsgStatusEnum) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return MsgStatusEnumMultiError(errors)
	}

	return nil
}

// MsgStatusEnumMultiError is an error wrapping multiple validation errors
// returned by MsgStatusEnum.ValidateAll() if the designated constraints
// aren't met.
type MsgStatusEnumMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MsgStatusEnumMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MsgStatusEnumMultiError) AllErrors() []error { return m }

// MsgStatusEnumValidationError is the validation error returned by
// MsgStatusEnum.Validate if the designated constraints aren't met.
type MsgStatusEnumValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MsgStatusEnumValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MsgStatusEnumValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MsgStatusEnumValidationError) Cause() error { return e.cause }

func (e MsgStatusEnumValidationError) Code() int64 { return 10000000 }

func (e MsgStatusEnumValidationError) HttpCode() int64 { return 400 }

// Key function returns key value.
func (e MsgStatusEnumValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MsgStatusEnumValidationError) ErrorName() string { return "MsgStatusEnumValidationError" }

// Error satisfies the builtin error interface
func (e MsgStatusEnumValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMsgStatusEnum.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MsgStatusEnumValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Code() int64
	HttpCode() int64
	Cause() error
	ErrorName() string
} = MsgStatusEnumValidationError{}

// Validate checks the field values on MsgDisplayTypeEnum with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *MsgDisplayTypeEnum) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MsgDisplayTypeEnum with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MsgDisplayTypeEnumMultiError, or nil if none found.
func (m *MsgDisplayTypeEnum) ValidateAll() error {
	return m.validate(true)
}

func (m *MsgDisplayTypeEnum) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return MsgDisplayTypeEnumMultiError(errors)
	}

	return nil
}

// MsgDisplayTypeEnumMultiError is an error wrapping multiple validation errors
// returned by MsgDisplayTypeEnum.ValidateAll() if the designated constraints
// aren't met.
type MsgDisplayTypeEnumMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MsgDisplayTypeEnumMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MsgDisplayTypeEnumMultiError) AllErrors() []error { return m }

// MsgDisplayTypeEnumValidationError is the validation error returned by
// MsgDisplayTypeEnum.Validate if the designated constraints aren't met.
type MsgDisplayTypeEnumValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MsgDisplayTypeEnumValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MsgDisplayTypeEnumValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MsgDisplayTypeEnumValidationError) Cause() error { return e.cause }

func (e MsgDisplayTypeEnumValidationError) Code() int64 { return 10000000 }

func (e MsgDisplayTypeEnumValidationError) HttpCode() int64 { return 400 }

// Key function returns key value.
func (e MsgDisplayTypeEnumValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MsgDisplayTypeEnumValidationError) ErrorName() string {
	return "MsgDisplayTypeEnumValidationError"
}

// Error satisfies the builtin error interface
func (e MsgDisplayTypeEnumValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMsgDisplayTypeEnum.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MsgDisplayTypeEnumValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Code() int64
	HttpCode() int64
	Cause() error
	ErrorName() string
} = MsgDisplayTypeEnumValidationError{}

// Validate checks the field values on MsgTypeEnum with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MsgTypeEnum) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MsgTypeEnum with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MsgTypeEnumMultiError, or
// nil if none found.
func (m *MsgTypeEnum) ValidateAll() error {
	return m.validate(true)
}

func (m *MsgTypeEnum) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return MsgTypeEnumMultiError(errors)
	}

	return nil
}

// MsgTypeEnumMultiError is an error wrapping multiple validation errors
// returned by MsgTypeEnum.ValidateAll() if the designated constraints aren't met.
type MsgTypeEnumMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MsgTypeEnumMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MsgTypeEnumMultiError) AllErrors() []error { return m }

// MsgTypeEnumValidationError is the validation error returned by
// MsgTypeEnum.Validate if the designated constraints aren't met.
type MsgTypeEnumValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MsgTypeEnumValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MsgTypeEnumValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MsgTypeEnumValidationError) Cause() error { return e.cause }

func (e MsgTypeEnumValidationError) Code() int64 { return 10000000 }

func (e MsgTypeEnumValidationError) HttpCode() int64 { return 400 }

// Key function returns key value.
func (e MsgTypeEnumValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MsgTypeEnumValidationError) ErrorName() string { return "MsgTypeEnumValidationError" }

// Error satisfies the builtin error interface
func (e MsgTypeEnumValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMsgTypeEnum.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MsgTypeEnumValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Code() int64
	HttpCode() int64
	Cause() error
	ErrorName() string
} = MsgTypeEnumValidationError{}

// Validate checks the field values on MsgButtonColorEnum with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *MsgButtonColorEnum) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MsgButtonColorEnum with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MsgButtonColorEnumMultiError, or nil if none found.
func (m *MsgButtonColorEnum) ValidateAll() error {
	return m.validate(true)
}

func (m *MsgButtonColorEnum) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return MsgButtonColorEnumMultiError(errors)
	}

	return nil
}

// MsgButtonColorEnumMultiError is an error wrapping multiple validation errors
// returned by MsgButtonColorEnum.ValidateAll() if the designated constraints
// aren't met.
type MsgButtonColorEnumMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MsgButtonColorEnumMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MsgButtonColorEnumMultiError) AllErrors() []error { return m }

// MsgButtonColorEnumValidationError is the validation error returned by
// MsgButtonColorEnum.Validate if the designated constraints aren't met.
type MsgButtonColorEnumValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MsgButtonColorEnumValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MsgButtonColorEnumValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MsgButtonColorEnumValidationError) Cause() error { return e.cause }

func (e MsgButtonColorEnumValidationError) Code() int64 { return 10000000 }

func (e MsgButtonColorEnumValidationError) HttpCode() int64 { return 400 }

// Key function returns key value.
func (e MsgButtonColorEnumValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MsgButtonColorEnumValidationError) ErrorName() string {
	return "MsgButtonColorEnumValidationError"
}

// Error satisfies the builtin error interface
func (e MsgButtonColorEnumValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMsgButtonColorEnum.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MsgButtonColorEnumValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Code() int64
	HttpCode() int64
	Cause() error
	ErrorName() string
} = MsgButtonColorEnumValidationError{}

// Validate checks the field values on MsgButtonActionTypeEnum with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *MsgButtonActionTypeEnum) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MsgButtonActionTypeEnum with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MsgButtonActionTypeEnumMultiError, or nil if none found.
func (m *MsgButtonActionTypeEnum) ValidateAll() error {
	return m.validate(true)
}

func (m *MsgButtonActionTypeEnum) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return MsgButtonActionTypeEnumMultiError(errors)
	}

	return nil
}

// MsgButtonActionTypeEnumMultiError is an error wrapping multiple validation
// errors returned by MsgButtonActionTypeEnum.ValidateAll() if the designated
// constraints aren't met.
type MsgButtonActionTypeEnumMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MsgButtonActionTypeEnumMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MsgButtonActionTypeEnumMultiError) AllErrors() []error { return m }

// MsgButtonActionTypeEnumValidationError is the validation error returned by
// MsgButtonActionTypeEnum.Validate if the designated constraints aren't met.
type MsgButtonActionTypeEnumValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MsgButtonActionTypeEnumValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MsgButtonActionTypeEnumValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MsgButtonActionTypeEnumValidationError) Cause() error { return e.cause }

func (e MsgButtonActionTypeEnumValidationError) Code() int64 { return 10000000 }

func (e MsgButtonActionTypeEnumValidationError) HttpCode() int64 { return 400 }

// Key function returns key value.
func (e MsgButtonActionTypeEnumValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MsgButtonActionTypeEnumValidationError) ErrorName() string {
	return "MsgButtonActionTypeEnumValidationError"
}

// Error satisfies the builtin error interface
func (e MsgButtonActionTypeEnumValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMsgButtonActionTypeEnum.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MsgButtonActionTypeEnumValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Code() int64
	HttpCode() int64
	Cause() error
	ErrorName() string
} = MsgButtonActionTypeEnumValidationError{}

// Validate checks the field values on MsgButtonStyleEnum with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *MsgButtonStyleEnum) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MsgButtonStyleEnum with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MsgButtonStyleEnumMultiError, or nil if none found.
func (m *MsgButtonStyleEnum) ValidateAll() error {
	return m.validate(true)
}

func (m *MsgButtonStyleEnum) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return MsgButtonStyleEnumMultiError(errors)
	}

	return nil
}

// MsgButtonStyleEnumMultiError is an error wrapping multiple validation errors
// returned by MsgButtonStyleEnum.ValidateAll() if the designated constraints
// aren't met.
type MsgButtonStyleEnumMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MsgButtonStyleEnumMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MsgButtonStyleEnumMultiError) AllErrors() []error { return m }

// MsgButtonStyleEnumValidationError is the validation error returned by
// MsgButtonStyleEnum.Validate if the designated constraints aren't met.
type MsgButtonStyleEnumValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MsgButtonStyleEnumValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MsgButtonStyleEnumValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MsgButtonStyleEnumValidationError) Cause() error { return e.cause }

func (e MsgButtonStyleEnumValidationError) Code() int64 { return 10000000 }

func (e MsgButtonStyleEnumValidationError) HttpCode() int64 { return 400 }

// Key function returns key value.
func (e MsgButtonStyleEnumValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MsgButtonStyleEnumValidationError) ErrorName() string {
	return "MsgButtonStyleEnumValidationError"
}

// Error satisfies the builtin error interface
func (e MsgButtonStyleEnumValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMsgButtonStyleEnum.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MsgButtonStyleEnumValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Code() int64
	HttpCode() int64
	Cause() error
	ErrorName() string
} = MsgButtonStyleEnumValidationError{}

// Validate checks the field values on PlatformEnum with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PlatformEnum) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PlatformEnum with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PlatformEnumMultiError, or
// nil if none found.
func (m *PlatformEnum) ValidateAll() error {
	return m.validate(true)
}

func (m *PlatformEnum) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return PlatformEnumMultiError(errors)
	}

	return nil
}

// PlatformEnumMultiError is an error wrapping multiple validation errors
// returned by PlatformEnum.ValidateAll() if the designated constraints aren't met.
type PlatformEnumMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PlatformEnumMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PlatformEnumMultiError) AllErrors() []error { return m }

// PlatformEnumValidationError is the validation error returned by
// PlatformEnum.Validate if the designated constraints aren't met.
type PlatformEnumValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlatformEnumValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlatformEnumValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlatformEnumValidationError) Cause() error { return e.cause }

func (e PlatformEnumValidationError) Code() int64 { return 10000000 }

func (e PlatformEnumValidationError) HttpCode() int64 { return 400 }

// Key function returns key value.
func (e PlatformEnumValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlatformEnumValidationError) ErrorName() string { return "PlatformEnumValidationError" }

// Error satisfies the builtin error interface
func (e PlatformEnumValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlatformEnum.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlatformEnumValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Code() int64
	HttpCode() int64
	Cause() error
	ErrorName() string
} = PlatformEnumValidationError{}