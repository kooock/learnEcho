package errors


type ErrorMsg struct {
	Msg string `json:"error"`
}

func NewErrorMsg(msg string) *ErrorMsg{
	return &ErrorMsg{Msg:msg}
}
