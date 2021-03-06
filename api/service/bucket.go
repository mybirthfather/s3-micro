package service

import (
	"github.com/lxlxw/s3-micro/pkg/s3"

	pb "github.com/lxlxw/s3-micro/proto"
)

func CreateBucket(r *pb.CreateBucketRequest) pb.CreateBucketResponse {

	client, _ := s3.New()

	err := client.CreateBucket(r.Bucketname, r.Publicread)
	if err != nil {
		return pb.CreateBucketResponse{Msg: err.Error(), Code: 403}
	}
	return pb.CreateBucketResponse{Msg: "success", Code: 200, Data: ""}
}
