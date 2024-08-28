package service

import (
	"context"
	"review-b/internal/biz"
	"review-b/internal/data/model"

	pb "review-b/api/business/v1"
)

type BusinessService struct {
	pb.UnimplementedBusinessServer
	uc *biz.BusinessUsecase
}

func NewBusinessService(uc *biz.BusinessUsecase) *BusinessService {
	return &BusinessService{uc: uc}
}

func (s *BusinessService) GreateReply(ctx context.Context, req *pb.CreateReplyReq) (*pb.CreateReplyResp, error) {

	reqs := &model.ReviewReplyInfo{
		ReviewID:  req.ReviewId,
		StoreID:   req.StoreId,
		Content:   req.Content,
		PicInfo:   req.PicInfo,
		VideoInfo: req.VideoInfo,
	}
	reply, err := s.uc.CreateReply(ctx, reqs)
	if err != nil {
		return nil, err
	}
	return &pb.CreateReplyResp{ReplyId: reply.ReplyID}, nil
}

// 创建appeal
func (s *BusinessService) CreateAppeal(ctx context.Context, req *pb.CreateAppealRequest) (*pb.CreateAppealReply, error) {
	reqs := &model.ReviewAppealInfo{
		AppealID:  req.AppealId,
		ReviewID:  req.ReviewId,
		StoreID:   req.StoreId,
		Status:    10,
		Reason:    req.Reason,
		Content:   req.Content,
		PicInfo:   req.PicInfo,
		VideoInfo: req.VideoInfo,
	}
	_, err := s.uc.CreateAppeal(ctx, reqs)
	if err != nil {
		return nil, err
	}
	return &pb.CreateAppealReply{}, nil
}
