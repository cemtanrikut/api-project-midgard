package helper

type ResponseHelper struct{
    StatusCode int
    Message string
    Model interface
}

func ResponseHandler(statusCode int, message string, model interface) ResponseHandler{
	// Logger will write to here

	helper := ResponseHandler {
		StatusCode: statusCode,
		Message: message,
		Model: model
	}

	return helper
}