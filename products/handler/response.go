package handler

import (
	"simple-echo-backend/products/core"
)

// Request
type ProductRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

// Response
type SuccessPayload struct {
	Status string `json:"status"`
}

func NewSuccessPayload() SuccessPayload {
	return SuccessPayload{Status: "success"}
}

type ProductPayloadResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ToProductPayload(product core.Products) ProductPayloadResponse {
	return ProductPayloadResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
	}
}

type GetProductResponse struct {
	Status  string                 `json:"status"`
	Payload ProductPayloadResponse `json:"payload"`
}

func NewGetProductResponse(result core.Products) GetProductResponse {
	productPayload := ToProductPayload(result)
	return GetProductResponse{
		Status:  "Success",
		Payload: productPayload,
	}
}

type GetProductListResponse struct {
	Status  string                   `json:"status"`
	Payload []ProductPayloadResponse `json:"payload"`
}

func NewGetProductListResponse(result []core.Products) GetProductListResponse {
	listProductResponse := make([]ProductPayloadResponse, len(result))
	for i, v := range result {
		listProductResponse[i] = ToProductPayload(v)
	}
	return GetProductListResponse{
		Status:  "success",
		Payload: listProductResponse,
	}
}

// status: success, error
type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func NewErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Status:  "error",
		Message: message,
	}
}
