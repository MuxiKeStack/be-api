// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package evaluationv1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsInvalidInput(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == EvaluationErrorReason_INVALID_INPUT.String() && e.Code == 400
}

func ErrorInvalidInput(format string, args ...interface{}) *errors.Error {
	return errors.New(400, EvaluationErrorReason_INVALID_INPUT.String(), fmt.Sprintf(format, args...))
}

func IsEvaluationNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == EvaluationErrorReason_EVALUATION_NOT_FOUND.String() && e.Code == 404
}

func ErrorEvaluationNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, EvaluationErrorReason_EVALUATION_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsCanNotEvaluateUnattendedCourse(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == EvaluationErrorReason_CAN_NOT_EVALUATE_UNATTENDED_COURSE.String() && e.Code == 403
}

func ErrorCanNotEvaluateUnattendedCourse(format string, args ...interface{}) *errors.Error {
	return errors.New(403, EvaluationErrorReason_CAN_NOT_EVALUATE_UNATTENDED_COURSE.String(), fmt.Sprintf(format, args...))
}

func IsGormTooManyRequest(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == EvaluationErrorReason_GORM_TOO_MANY_REQUEST.String() && e.Code == 429
}

func ErrorGormTooManyRequest(format string, args ...interface{}) *errors.Error {
	return errors.New(429, EvaluationErrorReason_GORM_TOO_MANY_REQUEST.String(), fmt.Sprintf(format, args...))
}
