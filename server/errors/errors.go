package errors

import (
	pb "Downloader_gRPC/server/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrorDownload(message string) (*pb.Stream, error) {
	err := status.Newf(
		codes.InvalidArgument,
		message)
	return nil, err.Err()
}

func ErrorUpload(message string) (*pb.Code, error) {
	code := &pb.Code{Code: pb.StatusCode_Failed}
	err := status.Newf(
		codes.InvalidArgument,
		message)
	return code, err.Err()
}
