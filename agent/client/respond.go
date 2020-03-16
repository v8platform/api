package sshclient

import (
	"encoding/json"
)

type RespondType string
type RespondErrorType string

const (
	LogType           RespondType = "log"     // ‑ информационное сообщение.
	SuccessType                   = "success" // ‑ операция успешно завершена.
	ErrorType                     = "error "  // ‑ операция завершена с ошибкой.
	CanceledType                  = "canceled"
	QuestionType                  = "question"
	DbstruType                    = "dbstru"
	LoadingIssueType              = "loading-issue"
	ProgressType                  = "progress"
	ExtensionInfoType             = "extension-info"
)

type Respond struct {
	Type      RespondType      `json:"type, not_null"`
	ErrorType RespondErrorType `json:"error-type"`
	Message   string           `json:"message"`
	Body      string           `json:"body, raw"`
}

func (res Respond) IsSuccess() bool {

	return res.Type == SuccessType

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

	return readRespond([]byte(str))

}
