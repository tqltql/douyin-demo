package controller

import (
	"douyin-demo/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FeedUser struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}
type FeedVideo struct {
	Id            int64    `json:"id"`
	Author        FeedUser `json:"author"`
	PlayUrl       string   `json:"play_url"`
	CoverUrl      string   `json:"cover_url"`
	FavoriteCount int64    `json:"favorite_count"`
	CommentCount  int64    `json:"comment_count"`
	IsFavorite    bool     `json:"is_favorite"`
	Title         string   `json:"title"`
}
type FeedResponse struct {
	repository.Response
	VideoList []FeedVideo `json:"video_list"`
	NextTime  int64       `json:"next_time"`
}

func Feed(c *gin.Context) {
	//查询参数
	sLatestTime := c.Query("latest_time")
	latestTime, err := strconv.ParseInt(sLatestTime, 10, 64)
	//类型转换失败
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response: repository.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	//token := c.Query("token")

	//feed服务
	feedResponse, err := FeedService(latestTime)

	//feed报错返回
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response: repository.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	//成功则返回视频列表和next_time
	if len(feedResponse.VideoList) > 0 {
		c.JSON(http.StatusOK, FeedResponse{
			Response:  repository.Response{StatusCode: 0},
			VideoList: feedResponse.VideoList,
			NextTime:  feedResponse.NextTime,
		})
	} else {
		c.JSON(http.StatusOK, FeedResponse{
			Response:  repository.Response{StatusCode: 0},
			VideoList: feedResponse.VideoList,
			NextTime:  0,
		})
	}

	return
}
func FeedService(latestTime int64) (*FeedResponse, error) {
	newFeedResponse := FeedResponse{}
	var nextTime int64
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	videoList, err := repository.NewVideoDaoInstance().QueryVideos(latestTime)
	//未找到视频列表
	if err != nil {
		return nil, err
	}
	for i, v := range *videoList {
		fmt.Printf("%d:::%+v\n", i, v)
		var feedVideo FeedVideo
		//将返回视频中发布最早的时间作为下次请求的latestTime
		if i == len(*videoList)-1 {
			nextTime = v.CreateTime.Unix()
		}
		feedVideo.Id = v.Id
		feedVideo.PlayUrl = v.PlayUrl
		feedVideo.CoverUrl = v.CoverUrl
		feedVideo.FavoriteCount = v.FavoriteCount
		feedVideo.CommentCount = v.CommentCount
		feedVideo.IsFavorite = false
		feedVideo.Title = v.Title

		user, err := repository.NewUserDaoInstance().QueryUserById(v.UserId)
		fmt.Printf("\n%d\n", v.UserId)
		//未找到对应用户信息
		if err != nil {
			return nil, err
		}
		var feedUser FeedUser
		feedUser.Id = user.Id
		feedUser.Name = user.Name
		feedUser.FollowCount = user.FollowCount
		feedUser.FollowerCount = user.FollowerCount
		feedUser.IsFollow = false

		feedVideo.Author = feedUser

		newFeedResponse.VideoList = append(newFeedResponse.VideoList, feedVideo)
	}
	newFeedResponse.NextTime = nextTime
	fmt.Printf("FeedResponse::%+v\n", newFeedResponse)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	return &newFeedResponse, nil
}
