package repository

import (
	"fmt"
	"gorm.io/gorm"
	"sync"
)

type User struct {
	Id             int64  `gorm:"column:id"`
	Name           string `gorm:"column:name"`
	Password       string `gorm:"column:password"`
	Token          string `gorm:"column:token"`
	FollowCount    int64  `gorm:"column:followCount"`
	FollowerCount  int64  `gorm:"column:followerCount"`
	VideoCount     int64  `gorm:"column:videoCount"`
	FavouriteCount int64  `gorm:"column:favouriteCount"`
}

func (User) TableName() string {
	return "user"
}

type UserDao struct {
}

var userDao *UserDao
var userOnce sync.Once

func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}

func (*UserDao) QueryUserById(id int64) (*User, error) {
	var user User
	err := db.Table("user").Where("id = ?", id).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		fmt.Println("find user by id err:" + err.Error())
		return nil, err
	}
	return &user, nil
}

func (*UserDao) MQueryUserById(ids []int64) (map[int64]*User, error) {
	var users []*User
	err := db.Table("user").Where("id in (?)", ids).Find(&users).Error
	if err != nil {
		fmt.Println("batch find user by id err:" + err.Error())
		return nil, err
	}
	userMap := make(map[int64]*User)
	for _, user := range users {
		userMap[user.Id] = user
	}
	return userMap, nil
}

func (*UserDao) QueryPasswordByUserName(userName string) (*User, error) {
	var user *User
	err := db.Where("name = ?", userName).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return user, nil
	}
	if err != nil {
		fmt.Println("find user by userName err:" + err.Error())
		return user, err
	}
	return user, nil
}
func (*UserDao) CreateUser(userName string, password string) error {
	var user User
	user.Name = userName
	user.Password = password
	//user表中新建
	err := db.Create(&user).Error
	//创建失败报错
	if err != nil {
		fmt.Println("create user err:" + err.Error())
		return err
	}
	return nil
}
