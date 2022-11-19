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

type UserResponse struct{
    StatusCode int
    Message string
    Model interface
}

//? Creates user
func CreateUserHandler(userStruct UserStruct) UserResponse {
    if userStruct == nil {
        response := UserResponse{
            StatusCode: http.StatusBadRequest,
            Message: "user couldn't fetch",
            Model: nil
        }

        return response
    }
    user := UserStruct{}

    response := UserResponse{
        StatusCode: http.StatusOK,
        Message: "Success",
        Model: user
    }

    return response
}

//? Gets user
func GetUserHandler(userID) UserResponse {
    if userID == ""{
        response := UserResponse{
            StatusCode: http.StatusBadRequest,
            Message: "user_id couldn't fetch",
            Model: nil
        }

        return response
    }

    user := UserStruct{}
    
    response := UserResponse{
        StatusCode: http.StatusOK,
        Message: "Success",
        Model: user
    }

    return response
}

func UpdateUserHandler(userStruct UserStruct) {
    if userStruct == nil {
        response := UserResponse{
            StatusCode: http.StatusBadRequest,
            Message: "user couldn't fetch",
            Model: nil
        }

        return response
    }

    response := UserResponse{
        StatusCode: http.StatusOK,
        Message: "user updated",
        Model: userStruct
    }

    return response
}

func DeleteUserHandler(userStruct UserStruct) {
    if userStruct == nil {
        response := UserResponse{
            StatusCode: http.StatusBadRequest,
            Message: "user couldn't fetch",
            Model: nil
        }

        return response
    }

    userStruct.IsActive = false

    response := UserResponse{
        StatusCode: http.StatusBadRequest,
        Message: "user deleted",
        Model: nil
    }

    return response
}