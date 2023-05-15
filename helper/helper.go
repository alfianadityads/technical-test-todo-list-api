package helper

type ResponseWithData struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseWithoutData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func APIResponseWithData(status string, message string, data interface{}) ResponseWithData {
	res := ResponseWithData{
		Message: message,
		Status:  status,
		Data:    data,
	}
	return res
}

func APIResponseWitouthData(status, message string) ResponseWithoutData {
	res := ResponseWithoutData{
		Message: message,
		Status:  status,
	}

	return res
}
