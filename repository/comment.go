package repository

type Comment struct {
	Id             int64  `gorm:"column:id"`
	VideoId        int64  `gorm:"column:videoId"`
	UserId         int64  `gorm:"column:userId"`
	Content        string `gorm:"column:content"`
	FavouriteCount int64  `gorm:"column:favouriteCount"`
}
