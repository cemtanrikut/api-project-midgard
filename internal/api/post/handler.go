package api

import(
	"time"
)

type PostStruct struct {
	ID string `json:id`
	Title string `json:title`
	Description string `json:description`
	Options OptionStruct[] `json:option`
	IsActive bool `json:IsActive`
	AddDate time.Time `json:add_date`
}

type OptionStruct struct {
	ID string `json:id`
	PostID string `json:post_id`
	Path string `json:path`
	VoteCount int `json:vote_count`
	URL string `json:url`
}

func AddPostHandler(post PostStruct) interface {
    if post == nil {
        response := ResponseHandler(http.StatusBadRequest, "post couldn't fetch", nil)
        return response
    }

	post.IsActive = true
	post.AddDate = time.Now

    response := ResponseHandler(http.StatusOK, "post added", post)
    return response
}

func GetPostHandler(postID string) interface {
	if postID == "" {
		response := ResponseHandler(http.StatusBadRequest, "post couldn't fetch", nil)
        return response
	}

	// TODO will find from DB

	response := ResponseHandler(http.StatusOK, "post is here", post)
    return response
⁡}

func VotePostHandler(optionID string) {
	// TODO Find post and marshall PostStruct model
	if postID == "" {
		response := ResponseHandler(http.StatusBadRequest, "post couldn't fetch", nil)
        return response
	}

	option := OptionStruct{}

	option.VoteCount++

    response := ResponseHandler(http.StatusOK, "post voted", option)
    return response	
}