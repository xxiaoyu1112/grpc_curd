package main

import (
	"context"
	"update_data/handler"
	"update_data/idl/update_data"
)

type Server struct {
	update_data.UnimplementedUpdateDataManageServer
}

// 添加挂载方法
func (s *Server) UpdateData(c context.Context, req *update_data.UpdateDataRequest) (*update_data.UpdateDataResponse, error) {

	handler := handler.NewUpdateDataHandler(c, req)

	return handler.Response, nil
}
