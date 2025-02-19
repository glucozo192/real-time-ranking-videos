package models

import (
	"mime/multipart"
	"strconv"
)

const (
	VideoImportHeaderName     = "Name"
	VideoImportHeaderVideoURL = "Link Video"
	VideoImportHeaderAuthor   = "Author"
	FolderImage               = "images"
	FolderVideo               = "video_k5_test"
	FolderResource            = "video"
	FolderZip                 = "App/zip"
	HD                        = "hd"
	CodeHD                    = 2
	HDR                       = "hdr"
	CodeHDR                   = 4
	FolderHd                  = "hd"
	FolderHdr                 = "hdr"
)

var DeviceTypes = map[string]string{
	HD:  strconv.Itoa(CodeHD),
	HDR: strconv.Itoa(CodeHDR),
}

var RequiredVideoHeaders = []string{
	VideoImportHeaderName,
	VideoImportHeaderVideoURL,
	VideoImportHeaderAuthor,
}

// HTTP
type CreateVideoRequest struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	VideoURL     string `json:"video_url"`
	Config       string `json:"config"`
	PathResource string `json:"path_resource"`
	LevelSystem  string `json:"level_system"`
	Status       string `json:"status"`
	Note         string `json:"note"`
	Assign       string `json:"assign"`
	Author       string `json:"author"`
}

type UpdateVideoRequest struct {
	ID           uint   `json:"id"`
	Name         string `json:"name" json:"name,omitempty"`
	Description  string `json:"description" json:"description,omitempty"`
	VideoURL     string `json:"video_url" json:"video_url,omitempty"`
	Config       string `json:"config" json:"config,omitempty"`
	PathResource string `json:"path_resource" json:"path_resource,omitempty"`
	LevelSystem  string `json:"level_system" json:"level_system,omitempty"`
	Status       string `json:"status" json:"status,omitempty"`
	Note         string `json:"note" json:"note,omitempty"`
	Assign       string `json:"assign" json:"assign,omitempty"`
	Author       string `json:"author" json:"author,omitempty"`
}

type VideoItem struct {
	VideoID      uint           `json:"video_id"`
	Name         string         `json:"name"`
	VideoURL     string         `json:"video_url"`
	PathResource string         `json:"path_resource"`
	Comments     []CommentItem  `json:"comment"`
	Viewers      []ViewerItem   `json:"viewer"`
	Reactions    []ReactionItem `json:"reaction"`
	Objects      []ObjectItem   `json:"object"`
}

type CommentItem struct {
	CommentID uint    `json:"comment_id"`
	Comment   string  `json:"comment"`
	UserName  string  `json:"user_name"`
	Avatar    string  `json:"avatar"`
	TimePoint float64 `json:"time_point"`
}

type ViewerItem struct {
	ViewerID  uint    `json:"viewer_id"`
	Number    int     `json:"number"`
	TimePoint float64 `json:"time_point"`
}

type ReactionItem struct {
	ReactionID uint    `json:"reaction_id"`
	Name       string  `json:"name"`
	Number     int     `json:"number"`
	TimePoint  float64 `json:"time_point"`
}

type ObjectItem struct {
	ObjectID    uint    `json:"object_id"`
	CoordinateX int     `json:"coordinate_x"`
	CoordinateY int     `json:"coordinate_y"`
	Length      int     `json:"length"`
	Width       int     `json:"width"`
	Order       int     `json:"order"`
	TimeStart   float64 `json:"time_start"`
	TimeEnd     float64 `json:"time_end"`
}

type UpdateVideoByExcel struct {
	File *multipart.FileHeader `form:"file" json:"file"`
}
