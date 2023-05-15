package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"todolistapi/features/activity"
	"todolistapi/helper"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type activHandler struct {
	srv activity.ActivityService
}

func New(srv activity.ActivityService) activity.ActivityHandler {
	return &activHandler{
		srv: srv,
	}
}

// Create implements activity.ActivityHandler
func (ah *activHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := CreateActivityReq{}
		if err := c.Bind(&input); err != nil {
			response := helper.APIResponseWithoutData("Bad Request", "Bad Request")
			return c.JSON(http.StatusBadRequest, response)
		}

		validate := validator.New()
		if err := validate.Struct(input); err != nil {
			msg := ""
			fmt.Println(err.Error())
			if strings.Contains(err.Error(), "Title") {
				msg = "title cannot be null"
			} else if strings.Contains(err.Error(), "Email") {
				msg = "email cannot be null"
			} else {
				msg = "request body cannot be null"
			}

			response := helper.APIResponseWithoutData("Bad Request", msg)
			return c.JSON(http.StatusBadRequest, response)
		}

		res, err := ah.srv.Create(*ReqToCore(input))
		if err != nil {
			response := helper.APIResponseWithoutData("Error", "Error")
			return c.JSON(http.StatusInternalServerError, response)
		}

		response := helper.APIResponseWithData("Success", "Success", CoreToResp(res))
		return c.JSON(http.StatusCreated, response)
	}
}

// Delete implements activity.ActivityHandler
func (ah *activHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response := helper.APIResponseWithoutData("Error", "Error")
			return c.JSON(http.StatusBadRequest, response)
		}

		err = ah.srv.Delete(uint(id))
		if err != nil {
			msg := fmt.Sprintf("Activity with ID %d Not Found", id)
			response := helper.APIResponseWithoutData("Not Found", msg)
			return c.JSON(http.StatusNotFound, response)
		}

		response := helper.APIResponseWithData("Success", "Success", helper.NoData{})
		return c.JSON(http.StatusOK, response)
	}
}

// GetAll implements activity.ActivityHandler
func (ah *activHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ah.srv.GetAll()
		if err != nil {
			response := helper.APIResponseWithoutData("Error", "Error")
			return c.JSON(http.StatusInternalServerError, response)
		}

		response := helper.APIResponseWithData("Success", "Success", CoreToRespArr(res))
		return c.JSON(http.StatusOK, response)
	}
}

// GetOne implements activity.ActivityHandler
func (ah *activHandler) GetOne() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response := helper.APIResponseWithoutData("Error", "Error")
			return c.JSON(http.StatusBadRequest, response)
		}

		res, err := ah.srv.GetOne(uint(id))
		if err != nil {
			msg := fmt.Sprintf("Activity with ID %d Not Found", id)
			response := helper.APIResponseWithoutData("Not Found", msg)
			return c.JSON(http.StatusNotFound, response)
		}

		response := helper.APIResponseWithData("Success", "Success", CoreToResp(res))
		return c.JSON(http.StatusOK, response)
	}
}

// Update implements activity.ActivityHandler
func (ah *activHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response := helper.APIResponseWithoutData("Error", "Error")
			return c.JSON(http.StatusNotFound, response)
		}

		input := UpdateActivityReq{}

		if err := c.Bind(&input); err != nil {
			response := helper.APIResponseWithoutData("Bad Request", "Bad request")
			return c.JSON(http.StatusBadRequest, response)
		}

		validate := validator.New()
		if err := validate.Struct(input); err != nil {
			response := helper.APIResponseWithoutData("Bad Request", "title cannot be null")
			return c.JSON(http.StatusBadRequest, response)
		}

		res, err := ah.srv.Update(uint(id), *ReqToCore(input))
		if err != nil {
			msg := fmt.Sprintf("Activity with ID %d Not Found", id)
			response := helper.APIResponseWithoutData("Not Found", msg)
			return c.JSON(http.StatusNotFound, response)
		}

		response := helper.APIResponseWithData("Success", "Success", CoreToResp(res))
		return c.JSON(http.StatusOK, response)
	}
}
