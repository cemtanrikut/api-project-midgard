package api

import (
	helper "api-project-midgard/helper"
	"net/http"
)

type UserStruct struct {
	Name      string `json:name`
	Surname   string `json:surname`
	Email     string `json:email`
	Password  string `json:password`
	Username  string `json:username`
	Birthdate string `json:birthdate`
	IsActive  bool   `json:is_active`
}

// ? Creates user
func CreateUserHandler(userStruct UserStruct) helper.ResponseHelper {
	user := UserStruct{}

	response := helper.ResponseHandler(http.StatusOK, "user created", user)
	return response
}

// ? Gets user
func GetUserHandler(userID string) helper.ResponseHelper {
	if userID == "" {
		response := helper.ResponseHandler(http.StatusBadRequest, "user_id couldn't fetch", nil)
		return response
	}

	user := UserStruct{}

	response := helper.ResponseHandler(http.StatusOK, "user created", user)
	return response
}

func UpdateUserHandler(userStruct UserStruct) helper.ResponseHelper {
	user := userStruct
	response := helper.ResponseHandler(http.StatusOK, "user updated", user)
	return response
}

func DeleteUserHandler(userStruct UserStruct) helper.ResponseHelper {
	user := userStruct
	user.IsActive = false

	response := helper.ResponseHandler(http.StatusOK, "user deleted", user)
	return response
}
