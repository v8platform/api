package errors

import (
	"fmt"
	"github.com/pkg/errors"
)

var (
	_ error = (*Error)(nil)
)

type Kind uint

const (
	NullError                          Kind = iota
	UnknownError                            // NullError ‑ неизвестная ошибка.
	DesignerNotConnectedToInfoBase          // DesignerNotConnectedToInfoBase ‑ соединение с информационной базой не установлено.
	DesignerAlreadyConnectedToInfoBase      // DesignerAlreadyConnectedToInfoBase ‑ соединение с информационной базой уже установлено.
	CommandFormatError                      // CommandFormatError ‑ неверный формат команды.
	DBRestructInfo                          // DBRestructInfo ‑ ошибка реструктуризации базы данных.
	InfoBaseNotFound                        // InfoBaseNotFound ‑ информационная база не найдена.
	AdministrationAccessRightRequired       // AdministrationAccessRightRequired ‑ для выполнения операции требуются административные права.
	ConfigFilesError                        // ConfigFilesError ‑ ошибки в процессе загрузки/выгрузки конфигурации из/в файла.
	DesignerAlreadyStarted                  // DesignerAlreadyStarted ‑ обнаружен запущенный конфигуратор.
	InfoBaseExclusiveLockRequired           // InfoBaseExclusiveLockRequired ‑ требуется исключительная блокировка информационной базы.
	LanguageNotFound                        // LanguageNotFound ‑ язык не обнаружен.
	ExtensionWithDataIsActive               // ExtensionWithDataIsActive ‑ расширение конфигурации активно и содержит данные.
	ExtensionNotFound                       // ExtensionNotFound ‑ расширение не обнаружено.
)

//
//func (e Kind) String() string {
//	switch e {
//	case NullError:
//		return "пустая ошибка"
//	case UnknownError:
//		return "неизвестная ошибка"
//	case DesignerNotConnectedToInfoBase:
//		return "соединение с информационной базой не установлено"
//	case DesignerAlreadyConnectedToInfoBase:
//		return "соединение с информационной базой уже установлено"
//	case CommandFormatError:
//		return "неверный формат команды"
//	case DBRestructInfo:
//		return "ошибка реструктуризации базы данных"
//	case InfoBaseNotFound:
//		return "информационная база не найдена"
//	case AdministrationAccessRightRequired:
//		return "для выполнения операции требуются административные права"
//	case ConfigFilesError:
//		return "ошибки в процессе загрузки/выгрузки конфигурации из/в файла"
//	case DesignerAlreadyStarted:
//		return "обнаружен запущенный конфигуратор"
//	case InfoBaseExclusiveLockRequired:
//		return "требуется исключительная блокировка информационной базы"
//	case LanguageNotFound:
//		return "язык не обнаружен"
//	case ExtensionWithDataIsActive:
//		return "расширение конфигурации активно и содержит данные"
//	case ExtensionNotFound:
//		return "расширение не обнаружено"
//	}
//	return "unknown error kind"
//}

type Error struct {
	kind        Kind
	err         error
	contextInfo errorContext
}

type errorContext struct {
	Field   string
	Message string
}

// Error returns the message of a Error
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
	err := Error{kind: e, err: errors.New(msg)}
	return err
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

	return Error{kind: NullError, err: wrappedError}
}

// AddErrorContext adds a context to an error
func AddErrorContext(err error, field, message string) error {
	context := errorContext{Field: field, Message: message}
	if customErr, ok := err.(Error); ok {
		return Error{kind: customErr.kind, err: customErr.err, contextInfo: context}
	}

	return Error{kind: NullError, err: err, contextInfo: context}
}

// GetErrorContext returns the error context
func GetErrorContext(err error) map[string]string {
	emptyContext := errorContext{}
	if customErr, ok := err.(Error); ok || customErr.contextInfo != emptyContext {

		return map[string]string{"field": customErr.contextInfo.Field, "message": customErr.contextInfo.Message}
	}

	return nil
}

func New(msg string) error {
	return NullError.New(msg)
}

// GetType returns the error type
func GetType(err error) Kind {
	if customErr, ok := err.(Error); ok {
		return customErr.kind
	}

	return NullError
}

// Is reports whether err is an *Error of the given Kind.
// If err is nil then Is returns false.
func Is(kind Kind, err error) bool {

	var e Error
	switch err.(type) {
	case *Error:
		e1, _ := err.(*Error)
		e = *e1
	case Error:
		e, _ = err.(Error)
	}

	if e.kind != NullError {
		return e.kind == kind
	}
	if e.err != nil {
		return Is(kind, e.err)
	}
	return false
}
