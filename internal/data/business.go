package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"review-b/api/review/v1"
	"review-b/internal/biz"
	"review-b/internal/data/model"
)

type businessrepo struct {
	data *Data
	log  *log.Helper
}

func NewBusinessRepo(data *Data, logger log.Logger) biz.BusinessRepo {
	return &businessrepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// 实现rpc调用。
func (br *businessrepo) CreateReply(ctx context.Context, replyInfo *model.ReviewReplyInfo) (*model.ReviewReplyInfo, error) {
	request := &v1.ReplyReviewReq{
		ReviewId:  replyInfo.ReviewID,
		StoreId:   replyInfo.StoreID,
		Content:   replyInfo.Content,
		PicInfo:   replyInfo.PicInfo,
		VideoInfo: replyInfo.VideoInfo,
	}
	review, err := br.data.rc.ReplyReview(ctx, request)
	if err != nil {
		br.log.Errorf("rpc create reply has error")
		fmt.Printf("err:", err)
		return nil, err
	}
	return &model.ReviewReplyInfo{ReplyID: review.ReplyId}, nil
}
