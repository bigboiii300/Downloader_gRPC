package main

import (
	"Downloader_gRPC/server/api"
	"Downloader_gRPC/server/errors"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type DownloadServer struct {
	api.FileDownloaderServer
	m map[string][]byte
}

func (s *DownloadServer) Download(_ context.Context, file *api.Stream) (*api.Stream, error) {
	content, ok := s.m[file.GetFileName()]
	if !ok {
		return errors.ErrorDownload("There is no file with this name")
	}
	p := &api.Stream{FileName: file.GetFileName(), Stream: content}
	return p, nil
}

func (s *DownloadServer) Upload(_ context.Context, file *api.Stream) (*api.Code, error) {
	_, ok := s.m[file.GetFileName()]
	if ok {
		return errors.ErrorUpload("A file with that name already exists")
	}
	s.m[file.GetFileName()] = file.GetStream()
	code := &api.Code{Code: api.StatusCode_Ok}
	return code, nil
}

func main() {
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln(err)
	}
	server := &DownloadServer{}
	server.m = make(map[string][]byte)
	api.RegisterFileDownloaderServer(s, server)
	err = s.Serve(lis)
	if err != nil {
		log.Fatalln(err)
	}
}
