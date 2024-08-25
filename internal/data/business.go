package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"review-b/internal/biz"
	"review-b/internal/data/model"
	"review-b/internal/data/query"
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
	err := br.data.query.Transaction(func(tx *query.Query) error {
		if err := tx.ReviewReplyInfo.WithContext(ctx).Save(replyInfo); err != nil {
			return err
		}
		if _, err := tx.ReviewInfo.WithContext(ctx).Where(tx.ReviewInfo.ReviewID.Eq(replyInfo.ReviewID)).Update(tx.ReviewInfo.HasReply, 0); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return replyInfo, err
}

func (br *businessrepo) SearchReview(ctx context.Context, review_id int64) (*model.ReviewInfo, error) {
	first, err := br.data.query.ReviewInfo.WithContext(ctx).Where(br.data.query.ReviewInfo.ReviewID.Eq(review_id)).First()
	if err != nil {
		return nil, errors.New("[data]find the review item has error")
	}
	return first, nil
}
