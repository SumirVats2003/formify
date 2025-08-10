package models

type Form struct {
	Id                string   `json:"id"`
	Title             string   `json:"title"`
	CreatorId         string   `json:"creator_id"`
	QuestionIds       []string `json:"question_ids"`
	AttachedSheet     string   `json:"sheet_url"`
	ValidityTimestamp int64    `json:"validity_timestamp"`
}

type FormRequest struct {
	Title             string            `json:"title"`
	CreatorId         string            `json:"creator_id"`
	Questions         []QuestionRequest `json:"questions"`
	ValidityTimestamp int64             `json:"validity_timestamp"`
}

type FormResponse struct {
	Id                string     `json:"id"`
	Title             string     `json:"title"`
	CreatorId         string     `json:"creator_id"`
	Questions         []Question `json:"questions"`
	AttachedSheet     string     `json:"sheet_url"`
	ValidityTimestamp int64      `json:"validity_timestamp"`
}

type Question struct {
	Id         string   `json:"id"`
	Title      string   `json:"title"`
	AnswerType string   `json:"answer_type"`
	Required   bool     `json:"required"`
	Options    []string `json:"options"`
}

type QuestionRequest struct {
	Title      string   `json:"title"`
	AnswerType string   `json:"answer_type"`
	Required   bool     `json:"required"`
	Options    []string `json:"options"`
}
