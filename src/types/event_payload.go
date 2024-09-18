package types

type EventPayload struct {
	// TODO : You can define the lambda event payload here
	Type    string `json:"type"`
	Message string `json:"message"`
}
