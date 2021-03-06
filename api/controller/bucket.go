package controller

import (
	"golang.org/x/net/context"

	"github.com/lxlxw/s3-micro/api/service"

	pb "github.com/lxlxw/s3-micro/proto"
)

func (h ApiService) CreateBucket(ctx context.Context, r *pb.CreateBucketRequest) (*pb.CreateBucketResponse, error) {
	res := service.CreateBucket(r)
	return &res, nil
}
