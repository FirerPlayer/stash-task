package web

type ResponseType string

const (
	Error ResponseType = "Error"
	//Success ResponseType = "Success"
	//Info ResponseType = "Info"
)

type SystemResponse struct {
	Type    ResponseType `json:"type"`
	Message string       `json:"message"`
}

func ErrorResponse(message string) *SystemResponse {
	return &SystemResponse{
		Type:    Error,
		Message: message,
	}
}
