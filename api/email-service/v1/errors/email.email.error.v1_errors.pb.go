// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package emailerrorv1

import (
	fmt "fmt"
	runtime "runtime"
	strconv "strconv"
	strings "strings"

	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

var EMAIL_ERRORMessageMap = map[int32]string{
	0:         "",
	100000001: "",
	100000002: "",
}

var EMAIL_ERRORHttpCodeMap = map[int32]int{
	0:         404,
	100000001: 400,
	100000002: 400,
}

func (e EMAIL_ERROR) stackTrace(skipFrame int) string {
	pc := make([]uintptr, 32)
	n := runtime.Callers(skipFrame, pc)
	pc = pc[:n]
	frames := runtime.CallersFrames(pc)
	msg := make([]string, 0, n)
	for {
		frame, more := frames.Next()
		funcName := frame.Function
		line := frame.Line
		file := frame.File
		msg = append(msg, fmt.Sprintf("\t%s:%d\n\t%s", file, line, funcName))
		if !more {
			break
		}
	}

	return strings.Join(msg, "\n")
}

func (e EMAIL_ERROR) toError(skipFrame int, msgFormat string, args ...interface{}) *errors.Error {
	err := errors.New(EMAIL_ERRORHttpCodeMap[int32(e.Number())], e.String(), fmt.Sprintf(msgFormat, args...))
	innerErr := errors.New(int(err.Code), err.Reason, err.Message).WithMetadata(map[string]string{
		"BizCode":        strconv.Itoa(int(e.Number())),
		"DefaultMessage": EMAIL_ERRORMessageMap[int32(e.Number())],
		"__Stack":        e.stackTrace(skipFrame),
		"__MetaKey":      "__",
	})
	return err.WithCause(innerErr)
}

func (e EMAIL_ERROR) ToError(msgFormat string, args ...interface{}) *errors.Error {
	return e.toError(4, msgFormat, args...)
}

// FromErrorf generate error from err with extra info, if err is nil, mean's everything is fine, return nil
func (e EMAIL_ERROR) FromErrorf(err error, format string, args ...interface{}) *errors.Error {
	if err == nil {
		return nil
	}

	te := e.toError(4, format, args...)
	return te.WithCause(te.Unwrap().(*errors.Error).WithCause(err))
}

// FromError generate error from err with extra info, if err is nil, mean's everything is fine, return nil
func (e EMAIL_ERROR) FromError(err error) *errors.Error {
	if err == nil {
		return nil
	}

	te := e.toError(4, "")
	return te.WithCause(te.Unwrap().(*errors.Error).WithCause(err))
}

// FromOrToError generate error from err with extra info, if err is nil, generate new error
func (e EMAIL_ERROR) FromOrToError(err error) *errors.Error {
	if err == nil {
		return e.toError(4, "")
	}

	te := e.toError(4, "")
	return te.WithCause(te.Unwrap().(*errors.Error).WithCause(err))
}

func (e EMAIL_ERROR) DefaultMessage() string {
	return EMAIL_ERRORMessageMap[int32(e.Number())]
}

func (e EMAIL_ERROR) HttpCode() int {
	return EMAIL_ERRORHttpCodeMap[int32(e.Number())]
}

func IsEmailUnknown(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == EMAIL_ERROR_EMAIL_UNKNOWN.String() && e.Code == 404
}

func ErrorEmailUnknown(format string, args ...interface{}) *errors.Error {
	return EMAIL_ERROR_EMAIL_UNKNOWN.toError(4, format, args...)
}

func IsEmailContentMissing(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == EMAIL_ERROR_EMAIL_CONTENT_MISSING.String() && e.Code == 400
}

func ErrorEmailContentMissing(format string, args ...interface{}) *errors.Error {
	return EMAIL_ERROR_EMAIL_CONTENT_MISSING.toError(4, format, args...)
}

func IsEmailContentError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == EMAIL_ERROR_EMAIL_CONTENT_ERROR.String() && e.Code == 400
}

func ErrorEmailContentError(format string, args ...interface{}) *errors.Error {
	return EMAIL_ERROR_EMAIL_CONTENT_ERROR.toError(4, format, args...)
}