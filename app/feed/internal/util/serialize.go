package util

import (
	"encoding/json"
	feedModel "github.com/oigi/Magikarp/app/feed/internal/model"
)

func StringVideoList(list []string) (videos []feedModel.Videos, err error) {
	for _, v := range list {
		var temp feedModel.Videos
		if err := json.Unmarshal([]byte(v), &temp); err != nil {
			return nil, err
		}
		videos = append(videos, temp)
	}

	return videos, nil
}

func VideoListString(list []feedModel.Videos) []string {
	var videos []string

	for _, v := range list {
		marshal, err := json.Marshal(v)
		if err != nil {
			continue
		}
		videos = append(videos, string(marshal))
	}
	return videos
}
