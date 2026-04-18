package work_usecase

type WorkInput struct {
	Type           string
	ConversationID *string
	Data           map[string]any // payload
}

type WorkOutput struct {
	ID       string
	Status   string
	Response string
}
