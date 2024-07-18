package dao

import (
	"context"
	CommentModel "github.com/oigi/Magikarp/app/comment/internal/model"
	"github.com/oigi/Magikarp/grpc/pb/comment"
	"github.com/oigi/Magikarp/grpc/pb/user"
	"github.com/oigi/Magikarp/pkg/mysql"
	"gorm.io/gorm"
)

type CommentDao struct {
	*gorm.DB
}

func NewCommentDao(ctx context.Context) *CommentDao {
	return &CommentDao{mysql.NewDBClient(ctx)}
}

func (c *CommentDao) AddComment(req *comment.CommentActionReq) error {
	comment := &CommentModel.Comment{
		UserID:  req.UserId,
		VideoID: req.VideoId,
		Content: req.CommentText,
	}
	return c.DB.Create(comment).Error
}

func (c *CommentDao) DelComment(req *comment.CommentActionReq) error {
	var comment CommentModel.Comment
	return c.DB.Where("user_id = ? AND video_id = ?", req.UserId, req.VideoId).Delete(comment).Error
}

func (c *CommentDao) GetCommentCount(req *comment.CommentCountReq) (resp *comment.CommentCountResp, err error) {
	// 初始化 CommentCount 字段
	resp = &comment.CommentCountResp{
		CommentCount: make(map[int64]int64),
	}
	var comment *CommentModel.Comment
	// 遍历视频ID列表并获取评论数
	for _, videoID := range req.VideoId {
		var count int64

		// 查询符合条件的评论数量
		if err := c.Model(&comment).Where("video_id = ?", videoID).Count(&count).Error; err != nil {
			return nil, err
		}

		// 将评论数存储到响应对象中
		resp.CommentCount[videoID] = count
	}
	return
}

func (c *CommentDao) GetCommentList(req *comment.CommentListReq) (resp *comment.CommentListResp, err error) {
	var comments []*CommentModel.Comment
	if err := c.Model(&CommentModel.Comment{}).Where("video_id = ?", req.VideoId).Find(&comments).Error; err != nil {
		return nil, err
	}

	// 初始化
	resp = &comment.CommentListResp{
		StatusCode:  0,
		StatusMsg:   "",
		CommentList: make([]*comment.Comment, len(comments)),
	}

	for k, v := range comments {
		// 获取评论作者的用户信息
		userInfo, err := c.GetUserInfo(v.UserID)
		if err != nil {
			return nil, err
		}

		// 填充评论信息到响应
		resp.CommentList[k] = &comment.Comment{
			Id:         v.ID,
			User:       convertUserModel(userInfo),
			Content:    v.Content,
			CreateDate: v.CreatedAt.String(),
		}
	}
	return
}

func (c *CommentDao) GetUserInfo(userID int64) (*CommentModel.User, error) {
	var user CommentModel.User
	err := c.Model(&CommentModel.User{}).Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func convertUserModel(userModel *CommentModel.User) *user.User {
	if userModel == nil {
		return nil
	}
	return &user.User{
		Id:   userModel.ID,
		Name: userModel.NickName,
	}
}
