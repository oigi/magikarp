package http

import (
	"github.com/gin-gonic/gin"
	"github.com/oigi/Magikarp/app/gateway/model"
	"github.com/oigi/Magikarp/app/gateway/rpc"
	"github.com/oigi/Magikarp/config"
	"github.com/oigi/Magikarp/grpc/pb/publish"
	"github.com/oigi/Magikarp/pkg/consts/e"
	"github.com/oigi/Magikarp/pkg/jwt"
	"go.uber.org/zap"
	"io"
	"net/http"
)

func CreateVideo(ctx *gin.Context) {
	var req publish.CreateVideoRequest
	resp := model.CommonResp{}
	video := model.PublishAction{}
	file, err := ctx.FormFile("data")
	if err != nil {
		// 处理文件上传错误
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 打开上传的文件
	uploadedFile, err := file.Open()
	if err != nil {
		// 处理文件打开错误
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer uploadedFile.Close()

	// 读取文件内容
	fileContent, err := io.ReadAll(uploadedFile)
	if err != nil {
		// 处理文件读取错误
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	video.Data = fileContent
	video.Token = ctx.PostForm("token")
	video.Title = ctx.PostForm("title")
	//data, err := ctx.GetRawData()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//// 把字节流重新放回 body 中
	//ctx.Request.Body = io.NopCloser(bytes.NewBuffer(data))
	//if err := json.NewDecoder(ctx.Request.Body).Decode(&video); err != nil {
	//	resp.StatusCode = e.ERROR
	//	resp.StatusMsg = "解析参数出错"
	//	ctx.JSON(http.StatusOK, resp)
	//	return
	//}

	var userId int64
	authHeader := video.Token
	if authHeader == "" {
		resp = model.CommonResp{
			StatusCode: e.InvalidParams,
			StatusMsg:  "用户不存在",
		}
		ctx.JSON(http.StatusOK, resp)
	} else {
		claims, err := jwt.ParseToken(authHeader)
		if err != nil {
			config.LOG.Error("解析错误", zap.Error(err))
			resp := model.CommonResp{
				StatusCode: e.ERROR,
				StatusMsg:  "解析错误",
			}
			ctx.JSON(http.StatusUnauthorized, resp)
		}
		userId = claims.ID
	}

	req.Data = video.Data
	req.Title = video.Title
	req.ActorId = userId

	//resp = model.CommonResp{ // TODO 临时操作
	//	StatusCode: e.DOUYINSUCCESS,
	//	StatusMsg:  "上传成功",
	//}
	//ctx.JSON(http.StatusOK, resp)

	r, err := rpc.CreateVideo(ctx, &req)
	if err != nil {
		config.LOG.Error("rpc CreateVideo 调用失败", zap.Error(err))
		ctx.JSON(http.StatusOK, r)
	}
	resp.StatusCode = e.DOUYINSUCCESS
	ctx.JSON(http.StatusOK, r)
}

func ListVideo(ctx *gin.Context) {
	var req publish.ListVideoRequest
	resp := model.CommonResp{}
	id := ctx.GetInt64("id")
	if id == 0 {
		resp.StatusCode = e.InvalidParams
		resp.StatusMsg = "用户不存在"
	}
	req.UserId = id

	r, err := rpc.ListVideo(ctx, &req)
	if err != nil {
		resp := model.CommonResp{
			StatusCode: e.ERROR,
			StatusMsg:  "rpc ListVideo 调用失败",
		}
		config.LOG.Error("rpc ListVideo 调用失败", zap.Error(err))
		ctx.JSON(http.StatusOK, resp)
		return
	}
	resp.StatusCode = e.DOUYINSUCCESS
	ctx.JSON(http.StatusOK, r)
}
