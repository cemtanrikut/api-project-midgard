package api

type UserStruct struct{
    Name string `json:name`
    Surname string `json:surname`
    Email string `json:email`
    Password string `json:password`
    Username string `json:username`
    Birthdate string `json:birthdate`
    IsActive bool `json:is_active`
}


//? Creates user
func CreateUserHandler(userStruct UserStruct) UserResponse {
    if userStruct == nil {
        response := ResponseHandler(http.StatusBadRequest, "user couldn't fetch", nil)
        return response
    }
    user := UserStruct{}
    
    response := ResponseHandler(http.StatusOK, "user created", user)
    return response
}

//? Gets user
func GetUserHandler(userID) UserResponse {
    if userID == ""{
        response := ResponseHandler(http.StatusBadRequest, "user_id couldn't fetch", nil)
        return response
    }

    user := UserStruct{}
    
    response := ResponseHandler(http.StatusOK, "user created", user)
    return response
}

func UpdateUserHandler(userStruct UserStruct) {
    if userStruct == nil {
        response := ResponseHandler(http.StatusBadRequest, "user couldn't fetch", nil)
        return response
    }

    response := ResponseHandler(http.StatusOK, "user updated", user)
    return response
}

func DeleteUserHandler(userStruct UserStruct) {
    if userStruct == nil {
        response := ResponseHandler(http.StatusBadRequest, "user couldn't fetch", nil)
        return response
    }

    userStruct.IsActive = false

    response := ResponseHandler(http.StatusOK, "user deleted", user)
    return response
}