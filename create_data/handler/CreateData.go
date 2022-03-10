package handler

import (
	"context"
	"rpc/create_data/idl/create_data"
)

type CreateDataHandler struct {
	Request  *create_data.CreateDataRequest
	Response *create_data.CreateDataResponse
	Ctx      context.Context
}

func NewCreateDataHandler(Ctx context.Context, Request *create_data.CreateDataRequest) *CreateDataHandler {
	return &CreateDataHandler{
		Request: Request,
		Ctx:     Ctx,
		Response: &create_data.CreateDataResponse{
			Message: "success",
			Status:  0,
			Data:    nil,
		},
	}
}
