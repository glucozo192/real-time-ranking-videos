package models

const (
	ViewerImportHeaderNumberView = "Number View"
	ViewerImportHeaderTime       = "Time (s)"
)

var RequiredViewerHeaders = []string{
	ViewerImportHeaderNumberView,
	ViewerImportHeaderTime,
}

type CreateViewerRequest struct {
	VideoID   uint    `json:"video_id"`
	Number    int     `json:"number"`
	TimePoint float64 `json:"time_point"`
}

type UpdateViewerRequest struct {
	VideoID   uint    `json:"video_id"`
	Number    int     `json:"number"`
	TimePoint float64 `json:"time_point"`
}
