package repository

import (
	"fmt"
	"gorm.io/gorm"
	"sync"
	"time"
)

const (
	LIMIT = 5
)

type Video struct {
	Id            int64     `gorm:"column:id"`
	UserId        int64     `gorm:"column:usersId"`
	PlayUrl       string    `gorm:"column:playUrl"`
	CoverUrl      string    `gorm:"column:coverUrl"`
	FavoriteCount int64     `gorm:"column:favouriteCount"`
	CommentCount  int64     `gorm:"column:commentCount"`
	Title         string    `gorm:"column:title"`
	CreateTime    time.Time `gorm:"column:createTime"`
}

func (Video) TableName() string {
	return "video"
}

type VideoDao struct {
}

var videoDao *VideoDao
var VideoOnce sync.Once

func NewVideoDaoInstance() *VideoDao {
	userOnce.Do(
		func() {
			videoDao = &VideoDao{}
		})
	return videoDao
}
func (*VideoDao) QueryVideos(latestTime int64) (*[]Video, error) {
	var videos []Video
	videos = make([]Video, LIMIT)
	//时间戳转换成time格式
	latest := time.Unix(latestTime, 0).Format("2021-06-10 15:04:05")
	//查询视频列表
	err := db.Table("video").Where("createTime > ?", latest).Order("createTime desc").Find(&videos).Limit(LIMIT).Error
	//fmt.Println(videos[0].UserId)
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}
	if err != nil {
		fmt.Println("find videos err:" + err.Error())
		return nil, err
	}
	return &videos, nil
}
