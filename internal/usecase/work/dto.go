package work_usecase

type WorkInput struct {
	Type           string
	ConversationID *string
	Data           map[string]any // payload
}

type WorkOutput struct {
	ID       string `json:"id,omitempty"`
	Status   string `json:"status,omitempty"`
	Response string `json:"response,omitempty"`
}
