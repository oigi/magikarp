package util

import "github.com/oigi/Magikarp/grpc/pb/publish"

func Check(req *publish.CreateVideoRequest) (bool, string) {
	if len(req.Data) == 0 {
		return true, "缺少视频数据"
	}

	if len(req.Title) == 0 {
		return true, "缺少标题"
	}

	if len(req.Label) == 0 {
		return true, "缺少标签"
	}
	if len(req.Category) == 0 {
		return true, "缺少分类"
	}
	return false, ""
}
