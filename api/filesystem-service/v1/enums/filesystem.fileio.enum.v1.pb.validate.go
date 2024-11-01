// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/filesystem-service/v1/enums/filesystem.fileio.enum.v1.proto

package filesystemenumv1

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

// Validate checks the field values on FileioInitEnum with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FileioInitEnum) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FileioInitEnum with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FileioInitEnumMultiError,
// or nil if none found.
func (m *FileioInitEnum) ValidateAll() error {
	return m.validate(true)
}

func (m *FileioInitEnum) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return FileioInitEnumMultiError(errors)
	}

	return nil
}

// FileioInitEnumMultiError is an error wrapping multiple validation errors
// returned by FileioInitEnum.ValidateAll() if the designated constraints
// aren't met.
type FileioInitEnumMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FileioInitEnumMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FileioInitEnumMultiError) AllErrors() []error { return m }

// FileioInitEnumValidationError is the validation error returned by
// FileioInitEnum.Validate if the designated constraints aren't met.
type FileioInitEnumValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FileioInitEnumValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FileioInitEnumValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FileioInitEnumValidationError) Cause() error { return e.cause }

func (e FileioInitEnumValidationError) Code() int64 { return 10000000 }

func (e FileioInitEnumValidationError) HttpCode() int64 { return 400 }

// Key function returns key value.
func (e FileioInitEnumValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FileioInitEnumValidationError) ErrorName() string { return "FileioInitEnumValidationError" }

// Error satisfies the builtin error interface
func (e FileioInitEnumValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFileioInitEnum.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FileioInitEnumValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Code() int64
	HttpCode() int64
	Cause() error
	ErrorName() string
} = FileioInitEnumValidationError{}
