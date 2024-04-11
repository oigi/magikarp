package dao

import (
	"context"
	relationModel "github.com/oigi/Magikarp/app/relation/internal/model"
	"github.com/oigi/Magikarp/grpc/pb/relation"
	"github.com/oigi/Magikarp/pkg/mysql"
	"gorm.io/gorm"
)

type RelationDao struct {
	*gorm.DB
}

func NewRelationDao(ctx context.Context) *RelationDao {
	return &RelationDao{mysql.NewDBClient(ctx)}
}

func (r *RelationDao) AddRelation(req *relation.ActionReq) error {
	result := &relationModel.Relation{
		UserId:         req.UserId,
		FollowedUserId: req.ToUserId,
		Status:         1,
	}
	return r.DB.Create(result).Error
}

func (r *RelationDao) DeleteRelation(req *relation.ActionReq) error {
	return r.DB.Model(&relationModel.Relation{}).
		Where("user_id = ? AND following_id = ?", req.UserId, req.ToUserId).
		Update("status", 0).Error // 0 表示'取消关注'
}

func (r *RelationDao) QueryFollowingUsers(userId int64) ([]int64, error) {
	var userIds []int64
	err := r.DB.Model(&relationModel.Relation{}).
		Where("user_id = ? AND status = 1", userId).
		Pluck("following_id", &userIds).Error
	return userIds, err
}

func (r *RelationDao) QueryFollowers(userId int64) ([]int64, error) {
	var userIds []int64
	err := r.DB.Model(&relationModel.Relation{}).
		Where("following_id = ? AND status = 1", userId).
		Pluck("user_id", &userIds).Error
	return userIds, err
}

func (r *RelationDao) QueryMutualFriends(userId int64) ([]int64, error) {
	var mutualIds []int64
	err := r.DB.Model(&relationModel.Relation{}).
		Where("user_id IN (SELECT following_id FROM relation WHERE user_id = ? AND status = 1) AND following_id = ? AND status = 1", userId, userId).
		Pluck("user_id", &mutualIds).Error

	return mutualIds, err
}
