package service

import (
	"github.com/oigi/Magikarp/grpc/pb/feed"
	"sync"
)

var FeedServeOnce sync.Once
var FeedServeIns *FeedServe

type FeedServe struct {
	feed.UnimplementedFeedServer
}

func GetFeedServe() *FeedServe {
	FeedServeOnce.Do(func() {
		FeedServeIns = &FeedServe{}
	})
	return FeedServeIns
}
