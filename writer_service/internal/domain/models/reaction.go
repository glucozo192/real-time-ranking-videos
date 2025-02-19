package models

const (
	ReactionImportHeaderReaction       = "Reaction"
	ReactionImportHeaderNumberReaction = "Number reaction"
	ReactionImportHeaderTime           = "Time (s)"
)

var RequiredReactionHeaders = []string{
	ReactionImportHeaderReaction,
	ReactionImportHeaderNumberReaction,
	ReactionImportHeaderTime,
}

// HTTP
type CreateReactionRequest struct {
	VideoID     uint    `json:"video_id"`
	Description string  `json:"description"`
	Name        string  `json:"name"`
	Number      int     `json:"number"`
	TimePoint   float64 `json:"time_point"`
}

type UpdateReactionRequest struct {
	VideoID     uint    `json:"video_id"`
	Description string  `json:"description"`
	Name        string  `json:"name"`
	Number      int     `json:"number"`
	TimePoint   float64 `json:"time_point"`
}
