package helper

type ResponseHelper struct {
	StatusCode int         `json:status_code`
	Message    string      `json:message`
	Model      interface{} `json:model`
}

func ResponseHandler(statusCode int, message string, model interface{}) ResponseHelper {
	// Logger will write to here

	helper := ResponseHelper{
		StatusCode: statusCode,
		Message:    message,
		Model:      model,
	}

	return helper
}
