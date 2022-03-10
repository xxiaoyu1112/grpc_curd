package main

import (
	"context"
	"get_data/handler"
	"get_data/idl/get_data"
)

type Server struct {
	get_data.UnimplementedGetDataCollectServer
}

// 挂载方法： 处理请求信息
func (s *Server) GetData(c context.Context, req *get_data.GetDataRequest) (*get_data.GetDataResponse, error) {

	handler1 := handler.NewGetDataHandler(c, req)
	return handler1.Response, nil

}
