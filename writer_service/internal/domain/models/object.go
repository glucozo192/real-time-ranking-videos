package models

// HTTP
type CreateObjectRequest struct {
	VideoID     uint    `json:"video_id"`
	Description string  `json:"description"`
	CoordinateX int     `json:"coordinate_x"`
	CoordinateY int     `json:"coordinate_y"`
	Length      int     `json:"length"`
	Width       int     `json:"width"`
	Order       int     `json:"order"`
	TimeStart   float64 `json:"time_start"`
	TimeEnd     float64 `json:"time_end"`
}

type UpdateObjectRequest struct {
	VideoID     uint    `json:"video_id"`
	Description string  `json:"description"`
	CoordinateX int     `json:"coordinate_x"`
	CoordinateY int     `json:"coordinate_y"`
	Length      int     `json:"length"`
	Width       int     `json:"width"`
	Order       int     `json:"order"`
	TimeStart   float64 `json:"time_start"`
	TimeEnd     float64 `json:"time_end"`
	MarkerName  string  `json:"marker_name"`
	TouchVector string  `json:"touch_vector"`
}
