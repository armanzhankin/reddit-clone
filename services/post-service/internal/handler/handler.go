package handler

import (
	"context"
	post_service "github.com/armanzhankin/reddit-clone/services/post-service/proto/post-service"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	post_service.UnimplementedPostServiceServer
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h Handler) Create(ctx context.Context, req *post_service.CreatePostRequest) (*post_service.CreatePostResponse, error) {
	return &post_service.CreatePostResponse{}, nil
}

func (h Handler) Get(ctx context.Context, req *post_service.GetPostRequest) (*post_service.GetPostResponse, error) {
	return &post_service.GetPostResponse{}, nil
}

func (h Handler) GetList(ctx context.Context, req *post_service.GetListRequest) (*post_service.GetListResponse, error) {
	return &post_service.GetListResponse{}, nil
}

func (h Handler) Update(ctx context.Context, req *post_service.UpdatePostRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (h Handler) Delete(ctx context.Context, req *post_service.DeletePostRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
