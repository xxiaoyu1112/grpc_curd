package handler

import (
	"context"
	"delete_data/idl/delete_data"
)

type DeleteDataHandler struct {
	Request  *delete_data.DeleteDataRequest
	Response *delete_data.DeleteDataResponse
	Ctx      context.Context
}

func NewDeleteDataHandler(ctx context.Context, request *delete_data.DeleteDataRequest) *DeleteDataHandler {
	return &DeleteDataHandler{
		Request: request,
		Ctx:     ctx,
		Response: &delete_data.DeleteDataResponse{
			Message:              "success",
			Status:               0,
			DeleteDataTaskStatus: nil,
		},
	}
}
