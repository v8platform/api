package sshclient

import (
	"encoding/json"
	"github.com/Khorevaa/go-v8runner/agent/client/errors"
	"strings"
)

type RespondType string
type RespondErrorType string

const (
	LogType           RespondType = "log"            // информационное сообщение.
	SuccessType                   = "success"        // операция успешно завершена.
	ErrorType                     = "error"          // операция завершена с ошибкой.
	CanceledType                  = "canceled"       // операция отменена
	QuestionType                  = "question"       // вопрос пользователю
	DbstruType                    = "dbstru"         // информация о процессе реструктуризации
	LoadingIssueType              = "loading-issue"  // ошибки и предупреждения, накопленные за время загрузки конфигурации из файлов
	ProgressType                  = "progress"       // информация о прогрессе выполнения команды
	ExtensionInfoType             = "extension-info" // информация о расширении, которое находится в информационной базе
	UnknownType                   = "unknown"
)

var UnknownRespond = Respond{Type: UnknownType}

type Respond struct {
	Type      RespondType      `json:"type, not_null"`
	ErrorType RespondErrorType `json:"error-type"`
	Message   string           `json:"message"`
	Body      json.RawMessage  `json:"body"`
}

func (res Respond) IsSuccess() bool {

	return res.Type == SuccessType

}

func (res Respond) IsError() bool {

	return res.Type == ErrorType

}

func (res Respond) IsProgress() bool {

	return res.Type == ProgressType

}

func (res Respond) Error() error {

	if res.Type != ErrorType {
		return nil
	}

	errType := res.ErrorType

	switch errType {

	case "UnknownError":
		return errors.UnknownError.New(res.Message)
	case "DesignerNotConnectedToInfoBase":
		return errors.DesignerNotConnectedToInfoBase.New(res.Message)
	case "DesignerAlreadyConnectedToInfoBase":
		return errors.DesignerAlreadyConnectedToInfoBase.New(res.Message)
	case "CommandFormatError":
		return errors.CommandFormatError.New(res.Message)
	case "DBRestructInfo":
		return errors.DBRestructInfo.New(res.Message)
	case "InfoBaseNotFound":
		return errors.InfoBaseNotFound.New(res.Message)
	case "AdministrationAccessRightRequired":
		return errors.AdministrationAccessRightRequired.New(res.Message)
	case "ConfigFilesError":
		return errors.ConfigFilesError.New(res.Message)
	case "DesignerAlreadyStarted":
		return errors.DesignerAlreadyStarted.New(res.Message)
	case "InfoBaseExclusiveLockRequired":
		return errors.InfoBaseExclusiveLockRequired.New(res.Message)
	case "LanguageNotFound":
		return errors.LanguageNotFound.New(res.Message)
	case "ExtensionWithDataIsActive":
		return errors.ExtensionWithDataIsActive.New(res.Message)
	case "ExtensionNotFound":
		return errors.ExtensionNotFound.New(res.Message)
	}

	return nil
}

func (res Respond) ReadBody(t interface{}) error {

	err := json.Unmarshal(res.Body, &t)
	if err != nil {
		return err
	}

	return nil

}

type ProgressInfo struct {
	Message string `json:"message"`
	Percent int    `json:"percent"`
}

type LoadingIssueLevel string

const (
	LoadingIssueLevelWarning LoadingIssueLevel = "warning"
	LoadingIssueLevelError                     = "error"
)

type LoadingIssue struct {
	Message string            `json:"message"`
	Level   LoadingIssueLevel `json:"level"`
}

func readRespond(data []byte) ([]Respond, error) {

	// json data
	var r []Respond

	// unmarshall it
	err := json.Unmarshal(data, &r)
	if err != nil {
		return r, err
	}

	return r, nil
}

func readRespondString(str string) (r []Respond, err error) {

	str = strings.ReplaceAll(str, "][", ",")

	return readRespond([]byte(str))

}
