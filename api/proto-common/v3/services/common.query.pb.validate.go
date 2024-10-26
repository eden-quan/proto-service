// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/proto-common/v3/services/common.query.proto

package common_services

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

// Validate checks the field values on QueryTypeEnum with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *QueryTypeEnum) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QueryTypeEnum with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in QueryTypeEnumMultiError, or
// nil if none found.
func (m *QueryTypeEnum) ValidateAll() error {
	return m.validate(true)
}

func (m *QueryTypeEnum) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return QueryTypeEnumMultiError(errors)
	}

	return nil
}

// QueryTypeEnumMultiError is an error wrapping multiple validation errors
// returned by QueryTypeEnum.ValidateAll() if the designated constraints
// aren't met.
type QueryTypeEnumMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QueryTypeEnumMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QueryTypeEnumMultiError) AllErrors() []error { return m }

// QueryTypeEnumValidationError is the validation error returned by
// QueryTypeEnum.Validate if the designated constraints aren't met.
type QueryTypeEnumValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QueryTypeEnumValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QueryTypeEnumValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QueryTypeEnumValidationError) Cause() error { return e.cause }

func (e QueryTypeEnumValidationError) Code() int64 { return 10000000 }

func (e QueryTypeEnumValidationError) HttpCode() int64 { return 400 }

// Key function returns key value.
func (e QueryTypeEnumValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QueryTypeEnumValidationError) ErrorName() string { return "QueryTypeEnumValidationError" }

// Error satisfies the builtin error interface
func (e QueryTypeEnumValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQueryTypeEnum.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QueryTypeEnumValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Code() int64
	HttpCode() int64
	Cause() error
	ErrorName() string
} = QueryTypeEnumValidationError{}

// Validate checks the field values on DataQuery with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DataQuery) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DataQuery with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DataQueryMultiError, or nil
// if none found.
func (m *DataQuery) ValidateAll() error {
	return m.validate(true)
}

func (m *DataQuery) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetQuery()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DataQueryValidationError{
					field:  "Query",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DataQueryValidationError{
					field:  "Query",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetQuery()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DataQueryValidationError{
				field:  "Query",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DataQueryMultiError(errors)
	}

	return nil
}

// DataQueryMultiError is an error wrapping multiple validation errors returned
// by DataQuery.ValidateAll() if the designated constraints aren't met.
type DataQueryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DataQueryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DataQueryMultiError) AllErrors() []error { return m }

// DataQueryValidationError is the validation error returned by
// DataQuery.Validate if the designated constraints aren't met.
type DataQueryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DataQueryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DataQueryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DataQueryValidationError) Cause() error { return e.cause }

func (e DataQueryValidationError) Code() int64 { return 10000000 }

func (e DataQueryValidationError) HttpCode() int64 { return 400 }

// Key function returns key value.
func (e DataQueryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DataQueryValidationError) ErrorName() string { return "DataQueryValidationError" }

// Error satisfies the builtin error interface
func (e DataQueryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDataQuery.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DataQueryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Code() int64
	HttpCode() int64
	Cause() error
	ErrorName() string
} = DataQueryValidationError{}

// Validate checks the field values on QueryChain with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *QueryChain) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QueryChain with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in QueryChainMultiError, or
// nil if none found.
func (m *QueryChain) ValidateAll() error {
	return m.validate(true)
}

func (m *QueryChain) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetQuery() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, QueryChainValidationError{
						field:  fmt.Sprintf("Query[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, QueryChainValidationError{
						field:  fmt.Sprintf("Query[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return QueryChainValidationError{
					field:  fmt.Sprintf("Query[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return QueryChainMultiError(errors)
	}

	return nil
}

// QueryChainMultiError is an error wrapping multiple validation errors
// returned by QueryChain.ValidateAll() if the designated constraints aren't met.
type QueryChainMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QueryChainMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QueryChainMultiError) AllErrors() []error { return m }

// QueryChainValidationError is the validation error returned by
// QueryChain.Validate if the designated constraints aren't met.
type QueryChainValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QueryChainValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QueryChainValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QueryChainValidationError) Cause() error { return e.cause }

func (e QueryChainValidationError) Code() int64 { return 10000000 }

func (e QueryChainValidationError) HttpCode() int64 { return 400 }

// Key function returns key value.
func (e QueryChainValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QueryChainValidationError) ErrorName() string { return "QueryChainValidationError" }

// Error satisfies the builtin error interface
func (e QueryChainValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQueryChain.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QueryChainValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Code() int64
	HttpCode() int64
	Cause() error
	ErrorName() string
} = QueryChainValidationError{}

// Validate checks the field values on DataMapping with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DataMapping) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DataMapping with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DataMappingMultiError, or
// nil if none found.
func (m *DataMapping) ValidateAll() error {
	return m.validate(true)
}

func (m *DataMapping) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	// no validation rules for InjectName

	// no validation rules for Query

	for idx, item := range m.GetArgs() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DataMappingValidationError{
						field:  fmt.Sprintf("Args[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DataMappingValidationError{
						field:  fmt.Sprintf("Args[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DataMappingValidationError{
					field:  fmt.Sprintf("Args[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetResp() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DataMappingValidationError{
						field:  fmt.Sprintf("Resp[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DataMappingValidationError{
						field:  fmt.Sprintf("Resp[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DataMappingValidationError{
					field:  fmt.Sprintf("Resp[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return DataMappingMultiError(errors)
	}

	return nil
}

// DataMappingMultiError is an error wrapping multiple validation errors
// returned by DataMapping.ValidateAll() if the designated constraints aren't met.
type DataMappingMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DataMappingMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DataMappingMultiError) AllErrors() []error { return m }

// DataMappingValidationError is the validation error returned by
// DataMapping.Validate if the designated constraints aren't met.
type DataMappingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DataMappingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DataMappingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DataMappingValidationError) Cause() error { return e.cause }

func (e DataMappingValidationError) Code() int64 { return 10000000 }

func (e DataMappingValidationError) HttpCode() int64 { return 400 }

// Key function returns key value.
func (e DataMappingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DataMappingValidationError) ErrorName() string { return "DataMappingValidationError" }

// Error satisfies the builtin error interface
func (e DataMappingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDataMapping.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DataMappingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Code() int64
	HttpCode() int64
	Cause() error
	ErrorName() string
} = DataMappingValidationError{}

// Validate checks the field values on DataBinding with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DataBinding) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DataBinding with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DataBindingMultiError, or
// nil if none found.
func (m *DataBinding) ValidateAll() error {
	return m.validate(true)
}

func (m *DataBinding) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Type

	switch v := m.PattenFrom.(type) {
	case *DataBinding_FromArg:
		if v == nil {
			err := DataBindingValidationError{
				field:  "PattenFrom",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for FromArg
	case *DataBinding_FromResp:
		if v == nil {
			err := DataBindingValidationError{
				field:  "PattenFrom",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for FromResp
	case *DataBinding_FromContext:
		if v == nil {
			err := DataBindingValidationError{
				field:  "PattenFrom",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for FromContext
	case *DataBinding_FromQuery:
		if v == nil {
			err := DataBindingValidationError{
				field:  "PattenFrom",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for FromQuery
	default:
		_ = v // ensures v is used
	}
	switch v := m.PattenTo.(type) {
	case *DataBinding_ToArg:
		if v == nil {
			err := DataBindingValidationError{
				field:  "PattenTo",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for ToArg
	case *DataBinding_ToResp:
		if v == nil {
			err := DataBindingValidationError{
				field:  "PattenTo",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for ToResp
	case *DataBinding_ToContext:
		if v == nil {
			err := DataBindingValidationError{
				field:  "PattenTo",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for ToContext
	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return DataBindingMultiError(errors)
	}

	return nil
}

// DataBindingMultiError is an error wrapping multiple validation errors
// returned by DataBinding.ValidateAll() if the designated constraints aren't met.
type DataBindingMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DataBindingMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DataBindingMultiError) AllErrors() []error { return m }

// DataBindingValidationError is the validation error returned by
// DataBinding.Validate if the designated constraints aren't met.
type DataBindingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DataBindingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DataBindingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DataBindingValidationError) Cause() error { return e.cause }

func (e DataBindingValidationError) Code() int64 { return 10000000 }

func (e DataBindingValidationError) HttpCode() int64 { return 400 }

// Key function returns key value.
func (e DataBindingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DataBindingValidationError) ErrorName() string { return "DataBindingValidationError" }

// Error satisfies the builtin error interface
func (e DataBindingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDataBinding.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DataBindingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Code() int64
	HttpCode() int64
	Cause() error
	ErrorName() string
} = DataBindingValidationError{}

// Validate checks the field values on TableDefine with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TableDefine) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TableDefine with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TableDefineMultiError, or
// nil if none found.
func (m *TableDefine) ValidateAll() error {
	return m.validate(true)
}

func (m *TableDefine) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	for idx, item := range m.GetIndex() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TableDefineValidationError{
						field:  fmt.Sprintf("Index[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TableDefineValidationError{
						field:  fmt.Sprintf("Index[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TableDefineValidationError{
					field:  fmt.Sprintf("Index[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return TableDefineMultiError(errors)
	}

	return nil
}

// TableDefineMultiError is an error wrapping multiple validation errors
// returned by TableDefine.ValidateAll() if the designated constraints aren't met.
type TableDefineMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TableDefineMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TableDefineMultiError) AllErrors() []error { return m }

// TableDefineValidationError is the validation error returned by
// TableDefine.Validate if the designated constraints aren't met.
type TableDefineValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TableDefineValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TableDefineValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TableDefineValidationError) Cause() error { return e.cause }

func (e TableDefineValidationError) Code() int64 { return 10000000 }

func (e TableDefineValidationError) HttpCode() int64 { return 400 }

// Key function returns key value.
func (e TableDefineValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TableDefineValidationError) ErrorName() string { return "TableDefineValidationError" }

// Error satisfies the builtin error interface
func (e TableDefineValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTableDefine.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TableDefineValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Code() int64
	HttpCode() int64
	Cause() error
	ErrorName() string
} = TableDefineValidationError{}

// Validate checks the field values on TableFindDefine with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *TableFindDefine) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TableFindDefine with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TableFindDefineMultiError, or nil if none found.
func (m *TableFindDefine) ValidateAll() error {
	return m.validate(true)
}

func (m *TableFindDefine) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Where

	if all {
		switch v := interface{}(m.GetPage()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TableFindDefineValidationError{
					field:  "Page",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TableFindDefineValidationError{
					field:  "Page",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TableFindDefineValidationError{
				field:  "Page",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Query

	if len(errors) > 0 {
		return TableFindDefineMultiError(errors)
	}

	return nil
}

// TableFindDefineMultiError is an error wrapping multiple validation errors
// returned by TableFindDefine.ValidateAll() if the designated constraints
// aren't met.
type TableFindDefineMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TableFindDefineMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TableFindDefineMultiError) AllErrors() []error { return m }

// TableFindDefineValidationError is the validation error returned by
// TableFindDefine.Validate if the designated constraints aren't met.
type TableFindDefineValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TableFindDefineValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TableFindDefineValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TableFindDefineValidationError) Cause() error { return e.cause }

func (e TableFindDefineValidationError) Code() int64 { return 10000000 }

func (e TableFindDefineValidationError) HttpCode() int64 { return 400 }

// Key function returns key value.
func (e TableFindDefineValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TableFindDefineValidationError) ErrorName() string { return "TableFindDefineValidationError" }

// Error satisfies the builtin error interface
func (e TableFindDefineValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTableFindDefine.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TableFindDefineValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Code() int64
	HttpCode() int64
	Cause() error
	ErrorName() string
} = TableFindDefineValidationError{}

// Validate checks the field values on TableFindPagingDefine with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *TableFindPagingDefine) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TableFindPagingDefine with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TableFindPagingDefineMultiError, or nil if none found.
func (m *TableFindPagingDefine) ValidateAll() error {
	return m.validate(true)
}

func (m *TableFindPagingDefine) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Page

	// no validation rules for Size

	if len(errors) > 0 {
		return TableFindPagingDefineMultiError(errors)
	}

	return nil
}

// TableFindPagingDefineMultiError is an error wrapping multiple validation
// errors returned by TableFindPagingDefine.ValidateAll() if the designated
// constraints aren't met.
type TableFindPagingDefineMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TableFindPagingDefineMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TableFindPagingDefineMultiError) AllErrors() []error { return m }

// TableFindPagingDefineValidationError is the validation error returned by
// TableFindPagingDefine.Validate if the designated constraints aren't met.
type TableFindPagingDefineValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TableFindPagingDefineValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TableFindPagingDefineValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TableFindPagingDefineValidationError) Cause() error { return e.cause }

func (e TableFindPagingDefineValidationError) Code() int64 { return 10000000 }

func (e TableFindPagingDefineValidationError) HttpCode() int64 { return 400 }

// Key function returns key value.
func (e TableFindPagingDefineValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TableFindPagingDefineValidationError) ErrorName() string {
	return "TableFindPagingDefineValidationError"
}

// Error satisfies the builtin error interface
func (e TableFindPagingDefineValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTableFindPagingDefine.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TableFindPagingDefineValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Code() int64
	HttpCode() int64
	Cause() error
	ErrorName() string
} = TableFindPagingDefineValidationError{}

// Validate checks the field values on IndexDefine with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *IndexDefine) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IndexDefine with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in IndexDefineMultiError, or
// nil if none found.
func (m *IndexDefine) ValidateAll() error {
	return m.validate(true)
}

func (m *IndexDefine) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return IndexDefineMultiError(errors)
	}

	return nil
}

// IndexDefineMultiError is an error wrapping multiple validation errors
// returned by IndexDefine.ValidateAll() if the designated constraints aren't met.
type IndexDefineMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IndexDefineMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IndexDefineMultiError) AllErrors() []error { return m }

// IndexDefineValidationError is the validation error returned by
// IndexDefine.Validate if the designated constraints aren't met.
type IndexDefineValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IndexDefineValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IndexDefineValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IndexDefineValidationError) Cause() error { return e.cause }

func (e IndexDefineValidationError) Code() int64 { return 10000000 }

func (e IndexDefineValidationError) HttpCode() int64 { return 400 }

// Key function returns key value.
func (e IndexDefineValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IndexDefineValidationError) ErrorName() string { return "IndexDefineValidationError" }

// Error satisfies the builtin error interface
func (e IndexDefineValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIndexDefine.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IndexDefineValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Code() int64
	HttpCode() int64
	Cause() error
	ErrorName() string
} = IndexDefineValidationError{}

// Validate checks the field values on BindDefine with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BindDefine) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BindDefine with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BindDefineMultiError, or
// nil if none found.
func (m *BindDefine) ValidateAll() error {
	return m.validate(true)
}

func (m *BindDefine) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Bind

	if len(errors) > 0 {
		return BindDefineMultiError(errors)
	}

	return nil
}

// BindDefineMultiError is an error wrapping multiple validation errors
// returned by BindDefine.ValidateAll() if the designated constraints aren't met.
type BindDefineMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BindDefineMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BindDefineMultiError) AllErrors() []error { return m }

// BindDefineValidationError is the validation error returned by
// BindDefine.Validate if the designated constraints aren't met.
type BindDefineValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BindDefineValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BindDefineValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BindDefineValidationError) Cause() error { return e.cause }

func (e BindDefineValidationError) Code() int64 { return 10000000 }

func (e BindDefineValidationError) HttpCode() int64 { return 400 }

// Key function returns key value.
func (e BindDefineValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BindDefineValidationError) ErrorName() string { return "BindDefineValidationError" }

// Error satisfies the builtin error interface
func (e BindDefineValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBindDefine.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BindDefineValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Code() int64
	HttpCode() int64
	Cause() error
	ErrorName() string
} = BindDefineValidationError{}

// Validate checks the field values on BindEnum with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BindEnum) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BindEnum with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BindEnumMultiError, or nil
// if none found.
func (m *BindEnum) ValidateAll() error {
	return m.validate(true)
}

func (m *BindEnum) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return BindEnumMultiError(errors)
	}

	return nil
}

// BindEnumMultiError is an error wrapping multiple validation errors returned
// by BindEnum.ValidateAll() if the designated constraints aren't met.
type BindEnumMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BindEnumMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BindEnumMultiError) AllErrors() []error { return m }

// BindEnumValidationError is the validation error returned by
// BindEnum.Validate if the designated constraints aren't met.
type BindEnumValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BindEnumValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BindEnumValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BindEnumValidationError) Cause() error { return e.cause }

func (e BindEnumValidationError) Code() int64 { return 10000000 }

func (e BindEnumValidationError) HttpCode() int64 { return 400 }

// Key function returns key value.
func (e BindEnumValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BindEnumValidationError) ErrorName() string { return "BindEnumValidationError" }

// Error satisfies the builtin error interface
func (e BindEnumValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBindEnum.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BindEnumValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Code() int64
	HttpCode() int64
	Cause() error
	ErrorName() string
} = BindEnumValidationError{}