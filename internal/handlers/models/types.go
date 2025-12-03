package models

type GetGreetingRequest struct {
	Thing string `json:"thing"`
}

type GetGreetingResponse struct {
	Greeting string `json:"greeting"`
}
