package base

type Response struct {
	IsSuccess bool   `json:"success"`
	Message   string `json:"message"`
	Error     string `json:"error,omitempty"`
	Status    uint   `json:"status"`
	Data      any    `json:"data"`
}

func CreateSuccessResponse(message string, status uint, data any) Response {
	return Response{
		IsSuccess: true,
		Message:   message,
		Status:    status,
		Data:      data,
	}
}

func CreateFailResponse(msg string, err string, statusCode uint) Response {
	return Response{
		IsSuccess: false,
		Message:   msg,
		Error:     err,
		Status:    statusCode,
		Data:      nil,
	}
}
