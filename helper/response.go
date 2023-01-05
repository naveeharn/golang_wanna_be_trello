package helper

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

func CreateResponse(status bool, message string, data interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
}

func CreateErrorResponse(messgae string, err string, data interface{}) Response {
	return Response{
		Status:  false,
		Message: messgae,
		Errors:  err,
		Data:    data,
	}
}
