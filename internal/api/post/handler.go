package api

import (
	helper "api-project-midgard/helper"
	"net/http"
	"time"
)

type PostStruct struct {
	ID          string       `json:id`
	Title       string       `json:title`
	Description string       `json:description`
	Options     OptionStruct `json:option`
	IsActive    bool         `json:IsActive`
	AddDate     time.Time    `json:add_date`
}

type OptionStruct struct {
	ID        string `json:id`
	PostID    string `json:post_id`
	Path      string `json:path`
	VoteCount int    `json:vote_count`
	URL       string `json:url`
}

func AddPostHandler(post PostStruct) helper.ResponseHelper {
	post.IsActive = true
	post.AddDate = time.Now()

	response := helper.ResponseHandler(http.StatusOK, "post added", post)
	return response
}

// GetPostHandler gets posts
// func GetPostHandler(postID string) ResponseHelper{
func GetPostHandler(postID string) helper.ResponseHelper {
	if postID == "" {
		response := helper.ResponseHandler(http.StatusBadRequest, "post couldn't fetch", nil)
		return response
	}

	// TODO will find from DB

	response := helper.ResponseHandler(http.StatusOK, "post is here", nil)
	return response
}

func VotePostHandler(postID string) helper.ResponseHelper {
	// TODO Find post and marshall PostStruct model
	if postID == "" {
		response := helper.ResponseHandler(http.StatusBadRequest, "post couldn't fetch", nil)
		return response
	}

	option := OptionStruct{}

	option.VoteCount++

	response := helper.ResponseHandler(http.StatusOK, "post voted", option)
	return response
}
