package upload

import (
	"context"
	"io"
	"os"
	"time"

	"google.golang.org/grpc"

	uploadpb "upload-file-client/proto"
)

type Client struct {
	client uploadpb.FileUploadServiceClient
}

func NewClient(conn grpc.ClientConnInterface) Client {
	return Client{
		client: uploadpb.NewFileUploadServiceClient(conn),
	}
}

func (c Client) Upload(ctx context.Context, file string) (string, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(10*time.Second))
	defer cancel()

	stream, err := c.client.Upload(ctx)
	if err != nil {
		return "", err
	}

	fil, err := os.Open(file)
	if err != nil {
		return "", err
	}

	// Maximum 1KB size per stream.
	buf := make([]byte, 1024)

	for {
		num, err := fil.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}

		uploadReq := uploadpb.UploadRequest{
			Data: &uploadpb.UploadRequest_File{
				File: buf[:num],
			},
		}
		if err := stream.Send(&uploadReq); err != nil {
			return "", err
		}
	}

	uploadReq := uploadpb.UploadRequest{
		Data: &uploadpb.UploadRequest_FileInfo{
			FileInfo: &uploadpb.FileInfo{
				Name: file,
				Type: "image",
			},
		},
	}

	if err := stream.Send(&uploadReq); err != nil {
		return "", err
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return "", err
	}

	return res.GetMsg(), nil
}
