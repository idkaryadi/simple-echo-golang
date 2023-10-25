package main

// Request
type ProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Response
type ProductResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type BaseResponse struct {
	Status  string      `json:"status"`
	Payload interface{} `json:"payload"`
}
