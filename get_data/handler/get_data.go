package handler

import (
	"context"
	"get_data/idl/get_data"
)

type GetDataHandler struct {
	Request  *get_data.GetDataRequest
	Response *get_data.GetDataResponse
	Ctx      context.Context
}

func NewGetDataHandler(Ctx context.Context, Request *get_data.GetDataRequest) *GetDataHandler {

	return &GetDataHandler{
		Request: Request,
		Ctx:     Ctx,
		Response: &get_data.GetDataResponse{
			Message: "success",
			Status:  0,
			Data:    nil,
		},
	}

}
