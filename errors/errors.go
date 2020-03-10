package errors

import (
	"fmt"
	"github.com/pkg/errors"
)

//
//const (
//	NoType = ErrorType(iota)
//	Internal1SError // при выполнении команды обнаружены ошибки в данных.
//	BadRequest
//	CheckError
//	TimeoutError
//	NotFound
//	BadDatabase = 101 // при выполнении команды обнаружены ошибки в данных.
//)

var (
	_ error = (*Error)(nil)
)

type Kind uint

const (
	Other         Kind = iota // Unclassified error. This value is not printed in the error message.
	Invalid                   // Invalid operation for this type of item.
	Permission                // Permission denied.
	IO                        // External I/O error such as network failure.
	Exist                     // Item already exists.
	NotExist                  // Item does not exist.
	IsDir                     // Item is a directory.
	NotDir                    // Item is not a directory.
	NotEmpty                  // Directory not empty.
	Private                   // Information withheld.
	Internal                  // Internal error or inconsistency.
	CannotDecrypt             // No wrapped key for user with read access.
	Transient                 // A transient error.
	BrokenLink                // Link target does not exist.
	Timeout                   // Link target does not exist.
	Database
	BadRequest
	BadConnectString
	BadCommand
	Runtime
)

type ErrorType uint

func (k Kind) String() string {
	switch k {
	case Other:
		return "other error"
	case Invalid:
		return "invalid operation"
	case Permission:
		return "permission denied"
	case IO:
		return "I/O error"
	case Exist:
		return "item already exists"
	case NotExist:
		return "item does not exist"
	case BrokenLink:
		return "link target does not exist"
	case IsDir:
		return "item is a directory"
	case NotDir:
		return "item is not a directory"
	case NotEmpty:
		return "directory not empty"
	case Private:
		return "information withheld"
	case Internal:
		return "internal error"
	case CannotDecrypt:
		return "cannot decrypt"
	case Transient:
		return "transient error"
	case Timeout:
		return "timeout error"
	case Database:
		return "database data invalid"
	case BadRequest:
		return "bad request"
	}
	return "unknown error kind"
}

type Error struct {
	kind        Kind
	err         error
	contextInfo errorContext
}

type errorContext struct {
	Field   string
	Message string
}

// Error returns the mssage of a Error
func (e Error) Error() string {
	return e.err.Error()
}

// AddErrorContext adds a context to an error
func (e Error) WithContext(field, message string) error {

	context := errorContext{Field: field, Message: message}
	return Error{kind: e.kind, err: e.err, contextInfo: context}

}

func (e *Error) IsZero() bool {
	return e.err == nil && e.kind == 0
}

// New creates a new Error
func (e Kind) New(msg string) Error {
	return Error{kind: e, err: errors.New(msg)}
}

// New creates a new Error with formatted message
func (e Kind) Newf(msg string, args ...interface{}) Error {
	err := fmt.Errorf(msg, args...)

	return Error{kind: e, err: err}
}

// Wrap creates a new wrapped error
func (e Kind) Wrap(err error, msg string) Error {
	return e.Wrapf(err, msg)
}

// Wrap creates a new wrapped error with formatted message
func (e Kind) Wrapf(err error, msg string, args ...interface{}) Error {
	newErr := errors.Wrapf(err, msg, args...)

	return Error{kind: e, err: newErr}
}

// Cause gives the original error
func Cause(err error) error {
	return errors.Cause(err)
}

// Wrapf wraps an error with format string
func Wrapf(err error, msg string, args ...interface{}) error {

	if err == nil {
		return err
	}

	wrappedError := errors.Wrapf(err, msg, args...)
	if customErr, ok := err.(Error); ok {
		return Error{
			kind:        customErr.kind,
			err:         wrappedError,
			contextInfo: customErr.contextInfo,
		}
	}

	return Error{kind: Other, err: wrappedError}
}

// AddErrorContext adds a context to an error
func AddErrorContext(err error, field, message string) error {
	context := errorContext{Field: field, Message: message}
	if customErr, ok := err.(Error); ok {
		return Error{kind: customErr.kind, err: customErr.err, contextInfo: context}
	}

	return Error{kind: Other, err: err, contextInfo: context}
}

// GetErrorContext returns the error context
func GetErrorContext(err error) map[string]string {
	emptyContext := errorContext{}
	if customErr, ok := err.(Error); ok || customErr.contextInfo != emptyContext {

		return map[string]string{"field": customErr.contextInfo.Field, "message": customErr.contextInfo.Message}
	}

	return nil
}

// GetType returns the error type
func GetType(err error) Kind {
	if customErr, ok := err.(Error); ok {
		return customErr.kind
	}

	return Other
}

// Is reports whether err is an *Error of the given Kind.
// If err is nil then Is returns false.
func Is(kind Kind, err error) bool {
	e, ok := err.(*Error)
	if !ok {
		return false
	}
	if e.kind != Other {
		return e.kind == kind
	}
	if e.err != nil {
		return Is(kind, e.err)
	}
	return false
}
