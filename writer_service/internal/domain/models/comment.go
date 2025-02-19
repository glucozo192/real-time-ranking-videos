package models

const (
	CommentImportHeaderName    = "Name"
	CommentImportHeaderAvatar  = "Avatar"
	CommentImportHeaderComment = "Comment"
	CommentImportHeaderTime    = "Time (s)"
)

var RequiredCommentHeaders = []string{
	CommentImportHeaderName,
	CommentImportHeaderAvatar,
	CommentImportHeaderComment,
	CommentImportHeaderTime,
}

// HTTP
type CreateCommentRequest struct {
	VideoID     uint    `json:"video_id"`
	Description string  `json:"description"`
	Comment     string  `json:"comment"`
	UserName    string  `json:"user_name"`
	Avatar      string  `json:"avatar"`
	TimePoint   float64 `json:"time_point"`
}

type UpdateCommentRequest struct {
	VideoID     uint    `json:"video_id"`
	Description string  `json:"description"`
	Comment     string  `json:"comment"`
	UserName    string  `json:"user_name"`
	Avatar      string  `json:"avatar"`
	TimePoint   float64 `json:"time_point"`
}
