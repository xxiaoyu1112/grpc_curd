package handler

import (
	"context"
	"update_data/idl/update_data"
)

type UpdateDataHandler struct {
	Request  *update_data.UpdateDataRequest
	Response *update_data.UpdateDataResponse
	Ctx      context.Context
}

func NewUpdateDataHandler(ctx context.Context, request *update_data.UpdateDataRequest) *UpdateDataHandler {
	return &UpdateDataHandler{
		Request: request,
		Ctx:     ctx,
		Response: &update_data.UpdateDataResponse{
			Message:              "success",
			Status:               0,
			UpdateDataTaskStatus: nil,
		},
	}
}
