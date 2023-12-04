package models

type Response struct {
	Timestamp string      `json:"timestamp"`
	Data      interface{} `json:"data"`
}

type Err struct {
	StatusCode  int    `json:"status"`
	StatusName  string `json:"name"`
	Message     string `json:"message"`     //the displayed error text for the user
	Description string `json:"description"` //the error description for the developer, aka err.Error()
}

func (e Err) Error() string {
	return e.Message
}
