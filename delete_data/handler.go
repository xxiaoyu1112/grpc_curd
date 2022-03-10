package main

import (
	"context"
	"delete_data/handler"
	"delete_data/idl/delete_data"
)

type Server struct {
	delete_data.UnimplementedDeleteDataManageServer
}

func (s *Server) DeleteData(c context.Context, req *delete_data.DeleteDataRequest) (*delete_data.DeleteDataResponse, error) {
	resp := handler.NewDeleteDataHandler(c, req)
	return resp.Response, nil
}
