package dto

type HttpResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ImportDataVideo struct {
	Unit        string        `json:"unit"`
	LinkVideo   string        `json:"link_video"`
	DataObjects []DataObjects `json:"data_objects"`
}

type DataObjects struct {
	MarkerName  string `json:"marker_name"`
	TimeStart   int    `json:"time_start"`
	TimeEnd     int    `json:"time_end"`
	TouchVector string `json:"touch_vector"`
}
