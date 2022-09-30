package main

import (
	"context"
	v1 "kylin-uploader/api/v1"

	"fmt"

	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func main() {
	conn, _ := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("127.0.0.1:9000"),
	)
	client := v1.NewChunkClient(conn)
	reply, _ := client.CreateUpload(context.TODO(), &v1.CreateUploadRequest{})
	fmt.Println(reply.UploadPath)
}
