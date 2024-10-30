// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package companyerrorv1

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

var COMPANY_ERRORMessageMap = map[int32]string{
	0:         "",
	100001:    "params error",
	100002:    "internal error",
	100003:    "authorization token error",
	100004:    "token invalid",
	100001401: "user is not exist",
	100001402: "company is not exist",
	100001403: "user not in company",
	100001404: "user not admin",
	100001405: "password is Decrypted failed",
	100001406: "upload data error",
	100001407: "privilege not good",
	100001408: "member is full",
	100001409: "use time is over",
	100001410: "passwd is empty",
	100001411: "passwd is same",
	100001412: "can not delete yourself",
	100001413: "this phone exist",
	100001414: "passwd length is not good",
	100001415: "business license is empty",
	100001416: "production license is empty",
	100001417: "privilege type error",
	100001418: "email regex error",
	100001419: "phone regex error",
	100001420: "role permission not enough",
	100001421: "user not registed",
	100001422: "invite self not allow",
	100001423: "forbid error",
	100001424: "user in company",
	100001425: "company not found",
	100001426: "invite record used",
	100001427: "create invite record error",
	100001428: "not support add member",
	100001429: "has not invite record",
	100001430: "cannot delete company",
	100001431: "not support delete member",
	100001432: "switch company is current company",
}

var COMPANY_ERRORHttpCodeMap = map[int32]int{
	0:         404,
	100001:    400,
	100002:    400,
	100003:    400,
	100004:    400,
	100001401: 400,
	100001402: 400,
	100001403: 400,
	100001404: 400,
	100001405: 400,
	100001406: 400,
	100001407: 400,
	100001408: 400,
	100001409: 400,
	100001410: 400,
	100001411: 400,
	100001412: 400,
	100001413: 400,
	100001414: 400,
	100001415: 400,
	100001416: 400,
	100001417: 400,
	100001418: 400,
	100001419: 400,
	100001420: 400,
	100001421: 400,
	100001422: 400,
	100001423: 400,
	100001424: 400,
	100001425: 400,
	100001426: 400,
	100001427: 400,
	100001428: 400,
	100001429: 400,
	100001430: 400,
	100001431: 400,
	100001432: 400,
}

func (e COMPANY_ERROR) stackTrace(skipFrame int) string {
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

func (e COMPANY_ERROR) toError(skipFrame int, msgFormat string, args ...interface{}) *errors.Error {
	err := errors.New(COMPANY_ERRORHttpCodeMap[int32(e.Number())], e.String(), fmt.Sprintf(msgFormat, args...))
	innerErr := errors.New(int(err.Code), err.Reason, err.Message).WithMetadata(map[string]string{
		"BizCode":        strconv.Itoa(int(e.Number())),
		"DefaultMessage": COMPANY_ERRORMessageMap[int32(e.Number())],
		"__Stack":        e.stackTrace(skipFrame),
		"__MetaKey":      "__",
	})
	return err.WithCause(innerErr)
}

func (e COMPANY_ERROR) ToError(msgFormat string, args ...interface{}) *errors.Error {
	return e.toError(4, msgFormat, args...)
}

// FromErrorf generate error from err with extra info, if err is nil, mean's everything is fine, return nil
func (e COMPANY_ERROR) FromErrorf(err error, format string, args ...interface{}) *errors.Error {
	if err == nil {
		return nil
	}

	te := e.toError(4, format, args...)
	return te.WithCause(te.Unwrap().(*errors.Error).WithCause(err))
}

// FromError generate error from err with extra info, if err is nil, mean's everything is fine, return nil
func (e COMPANY_ERROR) FromError(err error) *errors.Error {
	if err == nil {
		return nil
	}

	te := e.toError(4, "")
	return te.WithCause(te.Unwrap().(*errors.Error).WithCause(err))
}

// FromOrToError generate error from err with extra info, if err is nil, generate new error
func (e COMPANY_ERROR) FromOrToError(err error) *errors.Error {
	if err == nil {
		return e.toError(4, "")
	}

	te := e.toError(4, "")
	return te.WithCause(te.Unwrap().(*errors.Error).WithCause(err))
}

func (e COMPANY_ERROR) DefaultMessage() string {
	return COMPANY_ERRORMessageMap[int32(e.Number())]
}

func (e COMPANY_ERROR) HttpCode() int {
	return COMPANY_ERRORHttpCodeMap[int32(e.Number())]
}

func IsCompanyUnknown(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_COMPANY_UNKNOWN.String() && e.Code == 404
}

func ErrorCompanyUnknown(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_COMPANY_UNKNOWN.toError(4, format, args...)
}

// params error
func IsParam(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_PARAM.String() && e.Code == 400
}

// params error
func ErrorParam(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_PARAM.toError(4, format, args...)
}

// internal error
func IsInternal(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_INTERNAL.String() && e.Code == 400
}

// internal error
func ErrorInternal(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_INTERNAL.toError(4, format, args...)
}

// authorization token error
func IsAuthorizationToken(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_AUTHORIZATION_TOKEN.String() && e.Code == 400
}

// authorization token error
func ErrorAuthorizationToken(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_AUTHORIZATION_TOKEN.toError(4, format, args...)
}

// token invalid
func IsTokenInvalid(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_TOKEN_INVALID.String() && e.Code == 400
}

// token invalid
func ErrorTokenInvalid(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_TOKEN_INVALID.toError(4, format, args...)
}

// user is not exist
func IsUserNotExit(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_USER_NOT_EXIT.String() && e.Code == 400
}

// user is not exist
func ErrorUserNotExit(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_USER_NOT_EXIT.toError(4, format, args...)
}

// company is not exist
func IsCompanyNotExit(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_COMPANY_NOT_EXIT.String() && e.Code == 400
}

// company is not exist
func ErrorCompanyNotExit(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_COMPANY_NOT_EXIT.toError(4, format, args...)
}

// user not in company
func IsUserNotInCompany(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_USER_NOT_IN_COMPANY.String() && e.Code == 400
}

// user not in company
func ErrorUserNotInCompany(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_USER_NOT_IN_COMPANY.toError(4, format, args...)
}

// user not admin
func IsUserNotAdmin(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_USER_NOT_ADMIN.String() && e.Code == 400
}

// user not admin
func ErrorUserNotAdmin(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_USER_NOT_ADMIN.toError(4, format, args...)
}

// password is Decrypted failed
func IsPwdDecryptPassword(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_PWD_DecryptPassword.String() && e.Code == 400
}

// password is Decrypted failed
func ErrorPwdDecryptPassword(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_PWD_DecryptPassword.toError(4, format, args...)
}

// upload data error
func IsUploadData(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_UPLOAD_DATA.String() && e.Code == 400
}

// upload data error
func ErrorUploadData(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_UPLOAD_DATA.toError(4, format, args...)
}

// privilege not good
func IsPrivilegeNotEnough(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_PRIVILEGE_NOT_ENOUGH.String() && e.Code == 400
}

// privilege not good
func ErrorPrivilegeNotEnough(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_PRIVILEGE_NOT_ENOUGH.toError(4, format, args...)
}

// member is full
func IsMemberIsFull(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_MEMBER_IS_FULL.String() && e.Code == 400
}

// member is full
func ErrorMemberIsFull(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_MEMBER_IS_FULL.toError(4, format, args...)
}

// use time is over
func IsUsetimeIsFull(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_USETIME_IS_FULL.String() && e.Code == 400
}

// use time is over
func ErrorUsetimeIsFull(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_USETIME_IS_FULL.toError(4, format, args...)
}

// passwd is empty
func IsPwdEmpty(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_PWD_EMPTY.String() && e.Code == 400
}

// passwd is empty
func ErrorPwdEmpty(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_PWD_EMPTY.toError(4, format, args...)
}

// passwd is same
func IsPwdSame(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_PWD_SAME.String() && e.Code == 400
}

// passwd is same
func ErrorPwdSame(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_PWD_SAME.toError(4, format, args...)
}

// can not delete yourself
func IsDeleteYourself(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_DELETE_YOURSELF.String() && e.Code == 400
}

// can not delete yourself
func ErrorDeleteYourself(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_DELETE_YOURSELF.toError(4, format, args...)
}

// this phone exist
func IsPhoneExist(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_PHONE_EXIST.String() && e.Code == 400
}

// this phone exist
func ErrorPhoneExist(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_PHONE_EXIST.toError(4, format, args...)
}

// passwd length is not good
func IsPwdLengthNotGood(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_PWD_LENGTH_NOT_GOOD.String() && e.Code == 400
}

// passwd length is not good
func ErrorPwdLengthNotGood(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_PWD_LENGTH_NOT_GOOD.toError(4, format, args...)
}

// business license is empty
func IsBusinessLicenseEmpty(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_BUSINESS_LICENSE_EMPTY.String() && e.Code == 400
}

// business license is empty
func ErrorBusinessLicenseEmpty(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_BUSINESS_LICENSE_EMPTY.toError(4, format, args...)
}

// production license is empty
func IsProductionLicenseEmpty(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_PRODUCTION_LICENSE_EMPTY.String() && e.Code == 400
}

// production license is empty
func ErrorProductionLicenseEmpty(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_PRODUCTION_LICENSE_EMPTY.toError(4, format, args...)
}

// privilege type error
func IsPrivilegeTypeError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_PRIVILEGE_TYPE_ERROR.String() && e.Code == 400
}

// privilege type error
func ErrorPrivilegeTypeError(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_PRIVILEGE_TYPE_ERROR.toError(4, format, args...)
}

// email regex error
func IsEmailNotRegexMatch(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_EMAIL_NOT_REGEX_MATCH.String() && e.Code == 400
}

// email regex error
func ErrorEmailNotRegexMatch(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_EMAIL_NOT_REGEX_MATCH.toError(4, format, args...)
}

// phone regex error
func IsPhoneNotRegexMatch(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_PHONE_NOT_REGEX_MATCH.String() && e.Code == 400
}

// phone regex error
func ErrorPhoneNotRegexMatch(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_PHONE_NOT_REGEX_MATCH.toError(4, format, args...)
}

// role permission not enough
func IsRolePermissionNotEnough(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_ROLE_PERMISSION_NOT_ENOUGH.String() && e.Code == 400
}

// role permission not enough
func ErrorRolePermissionNotEnough(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_ROLE_PERMISSION_NOT_ENOUGH.toError(4, format, args...)
}

// user not registed
func IsUserNotRegisted(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_USER_NOT_REGISTED.String() && e.Code == 400
}

// user not registed
func ErrorUserNotRegisted(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_USER_NOT_REGISTED.toError(4, format, args...)
}

// invite self not allow
func IsInviteSelf(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_INVITE_SELF.String() && e.Code == 400
}

// invite self not allow
func ErrorInviteSelf(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_INVITE_SELF.toError(4, format, args...)
}

// forbid error
func IsForbid(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_FORBID.String() && e.Code == 400
}

// forbid error
func ErrorForbid(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_FORBID.toError(4, format, args...)
}

// user in company
func IsUserInCompany(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_USER_IN_COMPANY.String() && e.Code == 400
}

// user in company
func ErrorUserInCompany(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_USER_IN_COMPANY.toError(4, format, args...)
}

// company not found
func IsCompanyNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_COMPANY_NOT_FOUND.String() && e.Code == 400
}

// company not found
func ErrorCompanyNotFound(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_COMPANY_NOT_FOUND.toError(4, format, args...)
}

// invite record used
func IsInviteRecordUsed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_INVITE_RECORD_USED.String() && e.Code == 400
}

// invite record used
func ErrorInviteRecordUsed(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_INVITE_RECORD_USED.toError(4, format, args...)
}

// create invite record error
func IsCreateInviteRecord(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_CREATE_INVITE_RECORD.String() && e.Code == 400
}

// create invite record error
func ErrorCreateInviteRecord(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_CREATE_INVITE_RECORD.toError(4, format, args...)
}

// not support add member
func IsNotSupportAddMember(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_NOT_SUPPORT_ADD_MEMBER.String() && e.Code == 400
}

// not support add member
func ErrorNotSupportAddMember(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_NOT_SUPPORT_ADD_MEMBER.toError(4, format, args...)
}

// has not invite record
func IsHasNotInviteRecord(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_HAS_NOT_INVITE_RECORD.String() && e.Code == 400
}

// has not invite record
func ErrorHasNotInviteRecord(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_HAS_NOT_INVITE_RECORD.toError(4, format, args...)
}

// cannot delete company
func IsCannotDeletePersonOrg(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_CANNOT_DELETE_PERSON_ORG.String() && e.Code == 400
}

// cannot delete company
func ErrorCannotDeletePersonOrg(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_CANNOT_DELETE_PERSON_ORG.toError(4, format, args...)
}

// not support delete member
func IsNotSupportDelMember(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_NOT_SUPPORT_DEL_MEMBER.String() && e.Code == 400
}

// not support delete member
func ErrorNotSupportDelMember(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_NOT_SUPPORT_DEL_MEMBER.toError(4, format, args...)
}

// switch company is current company
func IsSwitchIsCurrentCompany(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == COMPANY_ERROR_SWITCH_IS_CURRENT_COMPANY.String() && e.Code == 400
}

// switch company is current company
func ErrorSwitchIsCurrentCompany(format string, args ...interface{}) *errors.Error {
	return COMPANY_ERROR_SWITCH_IS_CURRENT_COMPANY.toError(4, format, args...)
}