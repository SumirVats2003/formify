package models

type Form struct {
	Id                string   `json:"id"`
	Title             string   `json:"title"`
	CreatorId         string   `json:"creatorId"`
	QuestionIds       []string `json:"questions"`
	AttachedSheet     string   `json:"sheetUrl"`
	ValidityTimestamp int64    `json:"validityTimestamp"`
}

type FormRequest struct {
	Title             string            `json:"title"`
	CreatorId         string            `json:"creatorId"`
	Questions         []QuestionRequest `json:"questions"`
	ValidityTimestamp int64             `json:"validityTimestamp"`
}

type Question struct {
	Id            string   `json:"id"`
	QuestionTitle string   `json:"title"`
	AnswerType    string   `json:"answerType"`
	Required      bool     `json:"required"`
	Options       []string `json:"options"`
}

type QuestionRequest struct {
	QuestionTitle string   `json:"title"`
	AnswerType    string   `json:"answerType"`
	Required      bool     `json:"required"`
	Options       []string `json:"options"`
}
