// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package errorv1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
	runtime "runtime"
	strconv "strconv"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

var ERRORMessageMap = map[int32]string{
	0:   "UNKNOWN 常规",
	1:   "",
	2:   "",
	3:   "",
	4:   "",
	5:   "",
	6:   "",
	7:   "",
	8:   "",
	9:   "",
	10:  "",
	11:  "",
	12:  "",
	13:  "",
	14:  "",
	15:  "",
	16:  "",
	100: "CONTINUE Continue",
	101: "",
	102: "",
	103: "",
	200: "OK OK",
	201: "",
	202: "",
	203: "",
	204: "",
	205: "",
	206: "",
	207: "",
	208: "",
	226: "",
	300: "MULTIPLE_CHOICES MultipleChoices",
	301: "",
	302: "",
	303: "",
	304: "",
	305: "",
	306: "",
	307: "",
	308: "",
	400: "BAD_REQUEST Bad Request",
	401: "",
	402: "",
	403: "",
	404: "",
	405: "",
	406: "",
	407: "",
	408: "",
	409: "",
	410: "",
	411: "",
	412: "",
	413: "",
	414: "",
	415: "",
	416: "",
	417: "",
	418: "",
	421: "",
	422: "",
	423: "",
	424: "",
	425: "",
	426: "",
	428: "",
	429: "",
	431: "",
	451: "",
	500: "INTERNAL_SERVER Internal Server Error",
	501: "",
	502: "",
	503: "",
	504: "",
	505: "",
	506: "",
	507: "",
	508: "",
	510: "",
	511: "",
}

var ERRORHttpCodeMap = map[int32]int{
	0:   404,
	1:   424,
	2:   404,
	3:   409,
	4:   502,
	5:   504,
	6:   500,
	7:   500,
	8:   500,
	9:   500,
	10:  505,
	11:  500,
	12:  500,
	13:  500,
	14:  500,
	15:  500,
	16:  400,
	100: 100,
	101: 101,
	102: 102,
	103: 103,
	200: 200,
	201: 201,
	202: 202,
	203: 203,
	204: 204,
	205: 205,
	206: 206,
	207: 207,
	208: 208,
	226: 226,
	300: 300,
	301: 301,
	302: 302,
	303: 303,
	304: 304,
	305: 305,
	306: 306,
	307: 307,
	308: 308,
	400: 400,
	401: 401,
	402: 402,
	403: 403,
	404: 404,
	405: 405,
	406: 406,
	407: 407,
	408: 408,
	409: 409,
	410: 410,
	411: 411,
	412: 412,
	413: 413,
	414: 414,
	415: 415,
	416: 416,
	417: 417,
	418: 418,
	421: 421,
	422: 422,
	423: 423,
	424: 424,
	425: 425,
	426: 426,
	428: 428,
	429: 429,
	431: 431,
	451: 451,
	500: 500,
	501: 501,
	502: 502,
	503: 503,
	504: 504,
	505: 505,
	506: 506,
	507: 507,
	508: 508,
	510: 510,
	511: 511,
}

func (e ERROR) stackTrace(skipFrame int) string {
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

func (e ERROR) toError(skipFrame int, msgFormat string, args ...interface{}) *errors.Error {
	err := errors.New(ERRORHttpCodeMap[int32(e.Number())], e.String(), fmt.Sprintf(msgFormat, args...))
	innerErr := errors.New(int(err.Code), err.Reason, err.Message).WithMetadata(map[string]string{
		"BizCode":        strconv.Itoa(int(e.Number())),
		"DefaultMessage": ERRORMessageMap[int32(e.Number())],
		"__Stack":        e.stackTrace(skipFrame),
		"__MetaKey":      "__",
	})
	return err.WithCause(innerErr)
}

func (e ERROR) ToError(msgFormat string, args ...interface{}) *errors.Error {
	return e.toError(4, msgFormat, args...)
}

// FromErrorf generate error from err with extra info, if err is nil, mean's everything is fine, return nil
func (e ERROR) FromErrorf(err error, format string, args ...interface{}) *errors.Error {
	if err == nil {
		return nil
	}

	te := e.toError(4, format, args...)
	return te.WithCause(te.Unwrap().(*errors.Error).WithCause(err))
}

// FromError generate error from err with extra info, if err is nil, mean's everything is fine, return nil
func (e ERROR) FromError(err error) *errors.Error {
	if err == nil {
		return nil
	}

	te := e.toError(4, "")
	return te.WithCause(te.Unwrap().(*errors.Error).WithCause(err))
}

// FromOrToError generate error from err with extra info, if err is nil, generate new error
func (e ERROR) FromOrToError(err error) *errors.Error {
	if err == nil {
		return e.toError(4, "")
	}

	te := e.toError(4, "")
	return te.WithCause(te.Unwrap().(*errors.Error).WithCause(err))
}

func (e ERROR) DefaultMessage() string {
	return ERRORMessageMap[int32(e.Number())]
}

func (e ERROR) HttpCode() int {
	return ERRORHttpCodeMap[int32(e.Number())]
}

// UNKNOWN 常规
func IsUnknown(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_UNKNOWN.String() && e.Code == 404
}

// UNKNOWN 常规
func ErrorUnknown(format string, args ...interface{}) *errors.Error {
	return ERROR_UNKNOWN.toError(4, format, args...)
}

func IsRequestFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_REQUEST_FAILED.String() && e.Code == 424
}

func ErrorRequestFailed(format string, args ...interface{}) *errors.Error {
	return ERROR_REQUEST_FAILED.toError(4, format, args...)
}

func IsRecordNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_RECORD_NOT_FOUND.String() && e.Code == 404
}

func ErrorRecordNotFound(format string, args ...interface{}) *errors.Error {
	return ERROR_RECORD_NOT_FOUND.toError(4, format, args...)
}

func IsRecordAlreadyExists(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_RECORD_ALREADY_EXISTS.String() && e.Code == 409
}

func ErrorRecordAlreadyExists(format string, args ...interface{}) *errors.Error {
	return ERROR_RECORD_ALREADY_EXISTS.toError(4, format, args...)
}

func IsNetworkError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_NETWORK_ERROR.String() && e.Code == 502
}

func ErrorNetworkError(format string, args ...interface{}) *errors.Error {
	return ERROR_NETWORK_ERROR.toError(4, format, args...)
}

func IsNetworkTimeout(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_NETWORK_TIMEOUT.String() && e.Code == 504
}

func ErrorNetworkTimeout(format string, args ...interface{}) *errors.Error {
	return ERROR_NETWORK_TIMEOUT.toError(4, format, args...)
}

func IsConnection(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_CONNECTION.String() && e.Code == 500
}

func ErrorConnection(format string, args ...interface{}) *errors.Error {
	return ERROR_CONNECTION.toError(4, format, args...)
}

func IsUnimplemented(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_UNIMPLEMENTED.String() && e.Code == 500
}

func ErrorUnimplemented(format string, args ...interface{}) *errors.Error {
	return ERROR_UNIMPLEMENTED.toError(4, format, args...)
}

func IsFatal(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_FATAL.String() && e.Code == 500
}

func ErrorFatal(format string, args ...interface{}) *errors.Error {
	return ERROR_FATAL.toError(4, format, args...)
}

func IsPanic(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_PANIC.String() && e.Code == 500
}

func ErrorPanic(format string, args ...interface{}) *errors.Error {
	return ERROR_PANIC.toError(4, format, args...)
}

func IsRequestNotSupport(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_REQUEST_NOT_SUPPORT.String() && e.Code == 505
}

func ErrorRequestNotSupport(format string, args ...interface{}) *errors.Error {
	return ERROR_REQUEST_NOT_SUPPORT.toError(4, format, args...)
}

func IsDb(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_DB.String() && e.Code == 500
}

func ErrorDb(format string, args ...interface{}) *errors.Error {
	return ERROR_DB.toError(4, format, args...)
}

func IsRedis(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_REDIS.String() && e.Code == 500
}

func ErrorRedis(format string, args ...interface{}) *errors.Error {
	return ERROR_REDIS.toError(4, format, args...)
}

func IsRabbitMq(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_RABBIT_MQ.String() && e.Code == 500
}

func ErrorRabbitMq(format string, args ...interface{}) *errors.Error {
	return ERROR_RABBIT_MQ.toError(4, format, args...)
}

func IsRabbitKafka(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_RABBIT_KAFKA.String() && e.Code == 500
}

func ErrorRabbitKafka(format string, args ...interface{}) *errors.Error {
	return ERROR_RABBIT_KAFKA.toError(4, format, args...)
}

func IsMongo(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_MONGO.String() && e.Code == 500
}

func ErrorMongo(format string, args ...interface{}) *errors.Error {
	return ERROR_MONGO.toError(4, format, args...)
}

func IsInvalidParameter(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_INVALID_PARAMETER.String() && e.Code == 400
}

func ErrorInvalidParameter(format string, args ...interface{}) *errors.Error {
	return ERROR_INVALID_PARAMETER.toError(4, format, args...)
}

// CONTINUE Continue
func IsContinue(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_CONTINUE.String() && e.Code == 100
}

// CONTINUE Continue
func ErrorContinue(format string, args ...interface{}) *errors.Error {
	return ERROR_CONTINUE.toError(4, format, args...)
}

func IsSwitchingProtocols(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_SWITCHING_PROTOCOLS.String() && e.Code == 101
}

func ErrorSwitchingProtocols(format string, args ...interface{}) *errors.Error {
	return ERROR_SWITCHING_PROTOCOLS.toError(4, format, args...)
}

func IsProcessing(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_PROCESSING.String() && e.Code == 102
}

func ErrorProcessing(format string, args ...interface{}) *errors.Error {
	return ERROR_PROCESSING.toError(4, format, args...)
}

func IsEarlyHints(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_EARLY_HINTS.String() && e.Code == 103
}

func ErrorEarlyHints(format string, args ...interface{}) *errors.Error {
	return ERROR_EARLY_HINTS.toError(4, format, args...)
}

// OK OK
func IsOk(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_OK.String() && e.Code == 200
}

// OK OK
func ErrorOk(format string, args ...interface{}) *errors.Error {
	return ERROR_OK.toError(4, format, args...)
}

func IsCreated(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_CREATED.String() && e.Code == 201
}

func ErrorCreated(format string, args ...interface{}) *errors.Error {
	return ERROR_CREATED.toError(4, format, args...)
}

func IsAccepted(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_ACCEPTED.String() && e.Code == 202
}

func ErrorAccepted(format string, args ...interface{}) *errors.Error {
	return ERROR_ACCEPTED.toError(4, format, args...)
}

func IsNonAuthoritativeInfo(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_NON_AUTHORITATIVE_INFO.String() && e.Code == 203
}

func ErrorNonAuthoritativeInfo(format string, args ...interface{}) *errors.Error {
	return ERROR_NON_AUTHORITATIVE_INFO.toError(4, format, args...)
}

func IsNoContent(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_NO_CONTENT.String() && e.Code == 204
}

func ErrorNoContent(format string, args ...interface{}) *errors.Error {
	return ERROR_NO_CONTENT.toError(4, format, args...)
}

func IsResetContent(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_RESET_CONTENT.String() && e.Code == 205
}

func ErrorResetContent(format string, args ...interface{}) *errors.Error {
	return ERROR_RESET_CONTENT.toError(4, format, args...)
}

func IsPartialContent(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_PARTIAL_CONTENT.String() && e.Code == 206
}

func ErrorPartialContent(format string, args ...interface{}) *errors.Error {
	return ERROR_PARTIAL_CONTENT.toError(4, format, args...)
}

func IsMultiStatus(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_MULTI_STATUS.String() && e.Code == 207
}

func ErrorMultiStatus(format string, args ...interface{}) *errors.Error {
	return ERROR_MULTI_STATUS.toError(4, format, args...)
}

func IsAlreadyReported(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_ALREADY_REPORTED.String() && e.Code == 208
}

func ErrorAlreadyReported(format string, args ...interface{}) *errors.Error {
	return ERROR_ALREADY_REPORTED.toError(4, format, args...)
}

func IsIMUsed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_I_M_USED.String() && e.Code == 226
}

func ErrorIMUsed(format string, args ...interface{}) *errors.Error {
	return ERROR_I_M_USED.toError(4, format, args...)
}

// MULTIPLE_CHOICES MultipleChoices
func IsMultipleChoices(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_MULTIPLE_CHOICES.String() && e.Code == 300
}

// MULTIPLE_CHOICES MultipleChoices
func ErrorMultipleChoices(format string, args ...interface{}) *errors.Error {
	return ERROR_MULTIPLE_CHOICES.toError(4, format, args...)
}

func IsMovedPermanently(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_MOVED_PERMANENTLY.String() && e.Code == 301
}

func ErrorMovedPermanently(format string, args ...interface{}) *errors.Error {
	return ERROR_MOVED_PERMANENTLY.toError(4, format, args...)
}

func IsFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_FOUND.String() && e.Code == 302
}

func ErrorFound(format string, args ...interface{}) *errors.Error {
	return ERROR_FOUND.toError(4, format, args...)
}

func IsSeeOther(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_SEE_OTHER.String() && e.Code == 303
}

func ErrorSeeOther(format string, args ...interface{}) *errors.Error {
	return ERROR_SEE_OTHER.toError(4, format, args...)
}

func IsNotModified(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_NOT_MODIFIED.String() && e.Code == 304
}

func ErrorNotModified(format string, args ...interface{}) *errors.Error {
	return ERROR_NOT_MODIFIED.toError(4, format, args...)
}

func IsUseProxy(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_USE_PROXY.String() && e.Code == 305
}

func ErrorUseProxy(format string, args ...interface{}) *errors.Error {
	return ERROR_USE_PROXY.toError(4, format, args...)
}

func IsEmpty306(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_EMPTY306.String() && e.Code == 306
}

func ErrorEmpty306(format string, args ...interface{}) *errors.Error {
	return ERROR_EMPTY306.toError(4, format, args...)
}

func IsTemporaryRedirect(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_TEMPORARY_REDIRECT.String() && e.Code == 307
}

func ErrorTemporaryRedirect(format string, args ...interface{}) *errors.Error {
	return ERROR_TEMPORARY_REDIRECT.toError(4, format, args...)
}

func IsPermanentRedirect(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_PERMANENT_REDIRECT.String() && e.Code == 308
}

func ErrorPermanentRedirect(format string, args ...interface{}) *errors.Error {
	return ERROR_PERMANENT_REDIRECT.toError(4, format, args...)
}

// BAD_REQUEST Bad Request
func IsBadRequest(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_BAD_REQUEST.String() && e.Code == 400
}

// BAD_REQUEST Bad Request
func ErrorBadRequest(format string, args ...interface{}) *errors.Error {
	return ERROR_BAD_REQUEST.toError(4, format, args...)
}

func IsUnauthorized(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_UNAUTHORIZED.String() && e.Code == 401
}

func ErrorUnauthorized(format string, args ...interface{}) *errors.Error {
	return ERROR_UNAUTHORIZED.toError(4, format, args...)
}

func IsPaymentRequired(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_PAYMENT_REQUIRED.String() && e.Code == 402
}

func ErrorPaymentRequired(format string, args ...interface{}) *errors.Error {
	return ERROR_PAYMENT_REQUIRED.toError(4, format, args...)
}

func IsForbidden(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_FORBIDDEN.String() && e.Code == 403
}

func ErrorForbidden(format string, args ...interface{}) *errors.Error {
	return ERROR_FORBIDDEN.toError(4, format, args...)
}

func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_NOT_FOUND.String() && e.Code == 404
}

func ErrorNotFound(format string, args ...interface{}) *errors.Error {
	return ERROR_NOT_FOUND.toError(4, format, args...)
}

func IsMethodNotAllowed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_METHOD_NOT_ALLOWED.String() && e.Code == 405
}

func ErrorMethodNotAllowed(format string, args ...interface{}) *errors.Error {
	return ERROR_METHOD_NOT_ALLOWED.toError(4, format, args...)
}

func IsNotAcceptable(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_NOT_ACCEPTABLE.String() && e.Code == 406
}

func ErrorNotAcceptable(format string, args ...interface{}) *errors.Error {
	return ERROR_NOT_ACCEPTABLE.toError(4, format, args...)
}

func IsProxyAuthRequired(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_PROXY_AUTH_REQUIRED.String() && e.Code == 407
}

func ErrorProxyAuthRequired(format string, args ...interface{}) *errors.Error {
	return ERROR_PROXY_AUTH_REQUIRED.toError(4, format, args...)
}

func IsRequestTimeout(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_REQUEST_TIMEOUT.String() && e.Code == 408
}

func ErrorRequestTimeout(format string, args ...interface{}) *errors.Error {
	return ERROR_REQUEST_TIMEOUT.toError(4, format, args...)
}

func IsConflict(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_CONFLICT.String() && e.Code == 409
}

func ErrorConflict(format string, args ...interface{}) *errors.Error {
	return ERROR_CONFLICT.toError(4, format, args...)
}

func IsGone(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_GONE.String() && e.Code == 410
}

func ErrorGone(format string, args ...interface{}) *errors.Error {
	return ERROR_GONE.toError(4, format, args...)
}

func IsLengthRequired(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_LENGTH_REQUIRED.String() && e.Code == 411
}

func ErrorLengthRequired(format string, args ...interface{}) *errors.Error {
	return ERROR_LENGTH_REQUIRED.toError(4, format, args...)
}

func IsPreconditionFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_PRECONDITION_FAILED.String() && e.Code == 412
}

func ErrorPreconditionFailed(format string, args ...interface{}) *errors.Error {
	return ERROR_PRECONDITION_FAILED.toError(4, format, args...)
}

func IsRequestEntityTooLarge(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_REQUEST_ENTITY_TOO_LARGE.String() && e.Code == 413
}

func ErrorRequestEntityTooLarge(format string, args ...interface{}) *errors.Error {
	return ERROR_REQUEST_ENTITY_TOO_LARGE.toError(4, format, args...)
}

func IsRequestUriTooLong(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_REQUEST_URI_TOO_LONG.String() && e.Code == 414
}

func ErrorRequestUriTooLong(format string, args ...interface{}) *errors.Error {
	return ERROR_REQUEST_URI_TOO_LONG.toError(4, format, args...)
}

func IsUnsupportedMediaType(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_UNSUPPORTED_MEDIA_TYPE.String() && e.Code == 415
}

func ErrorUnsupportedMediaType(format string, args ...interface{}) *errors.Error {
	return ERROR_UNSUPPORTED_MEDIA_TYPE.toError(4, format, args...)
}

func IsRequestedRangeNotSatisfiable(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_REQUESTED_RANGE_NOT_SATISFIABLE.String() && e.Code == 416
}

func ErrorRequestedRangeNotSatisfiable(format string, args ...interface{}) *errors.Error {
	return ERROR_REQUESTED_RANGE_NOT_SATISFIABLE.toError(4, format, args...)
}

func IsExpectationFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_EXPECTATION_FAILED.String() && e.Code == 417
}

func ErrorExpectationFailed(format string, args ...interface{}) *errors.Error {
	return ERROR_EXPECTATION_FAILED.toError(4, format, args...)
}

func IsTeapot(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_TEAPOT.String() && e.Code == 418
}

func ErrorTeapot(format string, args ...interface{}) *errors.Error {
	return ERROR_TEAPOT.toError(4, format, args...)
}

func IsMisdirectedRequest(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_MISDIRECTED_REQUEST.String() && e.Code == 421
}

func ErrorMisdirectedRequest(format string, args ...interface{}) *errors.Error {
	return ERROR_MISDIRECTED_REQUEST.toError(4, format, args...)
}

func IsUnprocessableEntity(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_UNPROCESSABLE_ENTITY.String() && e.Code == 422
}

func ErrorUnprocessableEntity(format string, args ...interface{}) *errors.Error {
	return ERROR_UNPROCESSABLE_ENTITY.toError(4, format, args...)
}

func IsLocked(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_LOCKED.String() && e.Code == 423
}

func ErrorLocked(format string, args ...interface{}) *errors.Error {
	return ERROR_LOCKED.toError(4, format, args...)
}

func IsFailedDependency(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_FAILED_DEPENDENCY.String() && e.Code == 424
}

func ErrorFailedDependency(format string, args ...interface{}) *errors.Error {
	return ERROR_FAILED_DEPENDENCY.toError(4, format, args...)
}

func IsTooEarly(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_TOO_EARLY.String() && e.Code == 425
}

func ErrorTooEarly(format string, args ...interface{}) *errors.Error {
	return ERROR_TOO_EARLY.toError(4, format, args...)
}

func IsUpgradeRequired(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_UPGRADE_REQUIRED.String() && e.Code == 426
}

func ErrorUpgradeRequired(format string, args ...interface{}) *errors.Error {
	return ERROR_UPGRADE_REQUIRED.toError(4, format, args...)
}

func IsPreconditionRequired(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_PRECONDITION_REQUIRED.String() && e.Code == 428
}

func ErrorPreconditionRequired(format string, args ...interface{}) *errors.Error {
	return ERROR_PRECONDITION_REQUIRED.toError(4, format, args...)
}

func IsTooManyRequests(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_TOO_MANY_REQUESTS.String() && e.Code == 429
}

func ErrorTooManyRequests(format string, args ...interface{}) *errors.Error {
	return ERROR_TOO_MANY_REQUESTS.toError(4, format, args...)
}

func IsRequestHeaderFieldsTooLarge(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_REQUEST_HEADER_FIELDS_TOO_LARGE.String() && e.Code == 431
}

func ErrorRequestHeaderFieldsTooLarge(format string, args ...interface{}) *errors.Error {
	return ERROR_REQUEST_HEADER_FIELDS_TOO_LARGE.toError(4, format, args...)
}

func IsUnavailableForLegalReasons(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_UNAVAILABLE_FOR_LEGAL_REASONS.String() && e.Code == 451
}

func ErrorUnavailableForLegalReasons(format string, args ...interface{}) *errors.Error {
	return ERROR_UNAVAILABLE_FOR_LEGAL_REASONS.toError(4, format, args...)
}

// INTERNAL_SERVER Internal Server Error
func IsInternalServer(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_INTERNAL_SERVER.String() && e.Code == 500
}

// INTERNAL_SERVER Internal Server Error
func ErrorInternalServer(format string, args ...interface{}) *errors.Error {
	return ERROR_INTERNAL_SERVER.toError(4, format, args...)
}

func IsNotImplemented(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_NOT_IMPLEMENTED.String() && e.Code == 501
}

func ErrorNotImplemented(format string, args ...interface{}) *errors.Error {
	return ERROR_NOT_IMPLEMENTED.toError(4, format, args...)
}

func IsBadGateway(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_BAD_GATEWAY.String() && e.Code == 502
}

func ErrorBadGateway(format string, args ...interface{}) *errors.Error {
	return ERROR_BAD_GATEWAY.toError(4, format, args...)
}

func IsServiceUnavailable(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_SERVICE_UNAVAILABLE.String() && e.Code == 503
}

func ErrorServiceUnavailable(format string, args ...interface{}) *errors.Error {
	return ERROR_SERVICE_UNAVAILABLE.toError(4, format, args...)
}

func IsGatewayTimeout(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_GATEWAY_TIMEOUT.String() && e.Code == 504
}

func ErrorGatewayTimeout(format string, args ...interface{}) *errors.Error {
	return ERROR_GATEWAY_TIMEOUT.toError(4, format, args...)
}

func IsHttpVersionNotSupported(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_HTTP_VERSION_NOT_SUPPORTED.String() && e.Code == 505
}

func ErrorHttpVersionNotSupported(format string, args ...interface{}) *errors.Error {
	return ERROR_HTTP_VERSION_NOT_SUPPORTED.toError(4, format, args...)
}

func IsVariantAlsoNegotiates(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_VARIANT_ALSO_NEGOTIATES.String() && e.Code == 506
}

func ErrorVariantAlsoNegotiates(format string, args ...interface{}) *errors.Error {
	return ERROR_VARIANT_ALSO_NEGOTIATES.toError(4, format, args...)
}

func IsInsufficientStorage(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_INSUFFICIENT_STORAGE.String() && e.Code == 507
}

func ErrorInsufficientStorage(format string, args ...interface{}) *errors.Error {
	return ERROR_INSUFFICIENT_STORAGE.toError(4, format, args...)
}

func IsLoopDetected(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_LOOP_DETECTED.String() && e.Code == 508
}

func ErrorLoopDetected(format string, args ...interface{}) *errors.Error {
	return ERROR_LOOP_DETECTED.toError(4, format, args...)
}

func IsNotExtended(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_NOT_EXTENDED.String() && e.Code == 510
}

func ErrorNotExtended(format string, args ...interface{}) *errors.Error {
	return ERROR_NOT_EXTENDED.toError(4, format, args...)
}

func IsNetworkAuthenticationRequired(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ERROR_NETWORK_AUTHENTICATION_REQUIRED.String() && e.Code == 511
}

func ErrorNetworkAuthenticationRequired(format string, args ...interface{}) *errors.Error {
	return ERROR_NETWORK_AUTHENTICATION_REQUIRED.toError(4, format, args...)
}
