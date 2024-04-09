package dao

import (
    "context"
    "fmt"
    "github.com/oigi/Magikarp/app/favorite/internal/consts"
    "github.com/oigi/Magikarp/pkg/redis"
    "strconv"
)

var RedisClient map[string]*redis.ClientRedis

func WriteFavoriteInCache(ctx context.Context, userId int64, videoId int64, isFavorite bool) error {
    rdb := RedisClient[consts.FavCache]

    var status string
    if isFavorite {
        status = "1" // 已点赞
    } else {
        status = "0" // 未点赞
    }
    return rdb.Client.HSet(ctx, fmt.Sprint(userId), fmt.Sprint(videoId), status).Err()
}

func UpdateCacheCount(ctx context.Context, videoId int64, isFavorite bool) error {
    rdb := RedisClient[consts.FavCache]
    var status int
    if isFavorite {
        status = 1
    } else {
        status = -1
    }
    result, err := rdb.Client.Get(ctx, fmt.Sprint(videoId)).Result()
    if err != nil {
        return rdb.Client.Set(ctx, fmt.Sprint(videoId), "1", -1).Err()
    }
    atoi, _ := strconv.Atoi(result)
    atoi += status

    return rdb.Client.Set(context.Background(), fmt.Sprint(videoId), fmt.Sprint(atoi), -1).Err()
}

// QueryFavoriteListInCache 查询用户点赞列表
func QueryFavoriteListInCache(ctx context.Context, userId int64) ([]int64, error) {
    rdb := RedisClient[consts.FavCache]
    m, err := rdb.Client.HGetAll(ctx, fmt.Sprint(userId)).Result()
    if err != nil {
        return nil, err
    }
    var result []int64
    for k, v := range m {
        videoId, err := strconv.ParseInt(k, 10, 64)
        if err != nil {
            return nil, err
        }
        if v == "1" {
            result = append(result, videoId)
        }
    }
    return result, nil
}

func QueryFavoriteCount(videoIds []int64) (map[int64]int64, error) {
    rdb := RedisClient[consts.FavCache]
    counts := make(map[int64]int64)

    // 查询点赞数量
    for _, videoId := range videoIds {
        countStr, err := rdb.Client.Get(context.Background(), strconv.FormatInt(videoId, 10)).Result()
        if err != nil {
            return nil, err
        }
        count, err := strconv.ParseInt(countStr, 10, 64)
        if err != nil {
            return nil, err
        }
        counts[videoId] = count
    }

    return counts, nil
}
