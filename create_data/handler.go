package main

import (
	"context"
	"create_data/handler"
	"create_data/idl/create_data"
)

type Server struct {
	create_data.UnimplementedCreateDataCollectServer
}

func (s *Server) CreateData(ctx context.Context, req *create_data.CreateDataRequest) (*create_data.CreateDataResponse, error) {
	handler := handler.NewCreateDataHandler(ctx, req)
	return handler.Response, nil
}
