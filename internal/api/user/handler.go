package api

type UserStruct struct{
    Name string
    Surname string
    Email string
    Password string
    Nickname string
    Birthdate string
    Gender string
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
}