package main

import (
	"Downloader_gRPC/server/api"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io/ioutil"
	"os"
)

func main() {
	connection, err := grpc.Dial(":9000", grpc.WithInsecure())
	client := api.NewFileDownloaderClient(connection)
	var path string
	fmt.Println("Input path for upload file")
	fmt.Scanf("%s\n", &path)
	b, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	var name string
	fmt.Println("Input name and file extension. For example text.txt")
	fmt.Scanf("%s\n", &name)
	response, err := client.Upload(context.Background(), &api.Stream{Stream: b, FileName: name})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response.GetCode())

	fmt.Println("Input path for download file")
	fmt.Scanf("%s\n", &path)
	finalPath := fmt.Sprintf("%s/%s", path, name)
	req, err := client.Download(context.Background(), &api.Stream{FileName: name})
	err = ioutil.WriteFile(finalPath, req.GetStream(), 0777)
	if err != nil {
		fmt.Println(err)
	}
}
