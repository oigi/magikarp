package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/oigi/Magikarp/app/publish/internal/dao"
	"github.com/oigi/Magikarp/config"
	"github.com/oigi/Magikarp/grpc/pb/publish"
	"time"
)

func SavePublish(ctx context.Context, req *publish.CreateVideoRequest) error {
	timestamp := time.Now().Unix()
	hash := md5.New()
	hash.Write([]byte(fmt.Sprint(req.ActorId) + req.Title + req.Category + req.Label))
	filename := hex.EncodeToString(hash.Sum(nil))

	// 存本地
	// 使用 FFmpeg 转换视频格式
	// 生成视频封面缩略图

	// 存oss

	playUrl := "https://" + config.CONFIG.Oss.BucketName + "." + config.CONFIG.Oss.Endpoint + "/video/" + filename
	coverUrl := playUrl + "?x-oss-process=video/snapshot,t_30000,f_jpg,w_0,h_0,m_fast,ar_auto"
	// 2. 写入数据到MySQl
	id, err := dao.NewPublishDao(ctx).InsertVideo(req.ActorId, req.Title, playUrl, coverUrl, req.Label, req.Category)
	if err != nil {
		return err
	}
	// 3. 写入数据到Mongo
	err = dao.NewMongoClient(ctx).InsertVideoInMongo(id, req.ActorId, req.Title, playUrl, coverUrl, req.Label, req.Category, fmt.Sprint(timestamp))
	if err != nil {
		return err
	}
	return nil
}
