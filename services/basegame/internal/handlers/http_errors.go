package handlers

type HandlerErrors struct {}

func (e *HandlerErrors) InvalidRequestBody() map[string]string {
	return map[string]string{"error": "Invalid request body"}
}

func (e *HandlerErrors) MissingGameType() map[string]string {
	return map[string]string{"error": "Missing game type"}
}
