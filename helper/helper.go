package helper

import (
	"errors"
	"log"
	"strings"

	"github.com/go-playground/validator"
)

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

var validate *validator.Validate

func Validation(data interface{}) error {
	validate = validator.New()
	err := validate.Struct(data)
	if err != nil {
		log.Println(err)
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println(err)
		}
		msg := ""
		if strings.Contains(err.Error(), "required") {
			msg = "Body request cannot be blank"
		} else if strings.Contains(err.Error(), "title") {
			msg = "title cannot be null"
		} else if strings.Contains(err.Error(), "email") {
			msg = "email cannot be null"
		}
		return errors.New(msg)
	}
	
	return nil
}
