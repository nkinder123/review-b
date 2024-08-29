package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"review-b/internal/data/model"
)

type BusinessRepo interface {
	CreateReply(context.Context, *model.ReviewReplyInfo) (*model.ReviewReplyInfo, error)
	CreateAppeal(context.Context, *model.ReviewAppealInfo) (*model.ReviewAppealInfo, error)
}

type BusinessUsecase struct {
	repo BusinessRepo
	log  *log.Helper
}

func NewBusinessUsecase(repo BusinessRepo, logger log.Logger) *BusinessUsecase {
	return &BusinessUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *BusinessUsecase) CreateReply(ctx context.Context, info *model.ReviewReplyInfo) (*model.ReviewReplyInfo, error) {
	//参数校验
	//1.1 是否已经有回复了
	//1.2 水平越权
	//review, err := uc.repo.SearchReview(ctx, info.ReviewID)
	//if err != nil {
	//	return nil, err
	//}
	//if review.HasReply == 1 {
	//	return nil, errors.New("the business has reply,deny repeated reply ")
	//}
	//if review.StoreID != info.StoreID {
	//	return nil, errors.New("the replyId  is not the direct business")
	//}
	//reply, err := uc.repo.CreateReply(ctx, info)
	//if err != nil {
	//	return nil, err
	//}
	//return reply, err

	reply, err := uc.repo.CreateReply(ctx, info)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (uc *BusinessUsecase) CreateAppeal(ctx context.Context, info *model.ReviewAppealInfo) (*model.ReviewAppealInfo, error) {
	log.Info("[biz create appealy]")
	appeal, err := uc.repo.CreateAppeal(ctx, info)
	if err != nil {
		return nil, err
	}
	return appeal, nil
}
