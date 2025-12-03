package models

type GetGreetingRequest struct {
	Thing string `json:"thing" validate:"min=1,max=32"`
}

type GetGreetingResponse struct {
	Greeting string `json:"greeting"`
}
