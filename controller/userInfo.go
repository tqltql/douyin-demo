package controller

import (
	"douyin-demo/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserInfoQuery struct {
	UserId        int64  `json:"user_id"`
	UserName      string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type UserInfoResponse struct {
	repository.Response
	UserInf UserInfoQuery `json:"user"`
}

func UserInfo(c *gin.Context) {
	//查询参数
	userId := c.Query("user_id")
	token := c.Query("token")

	//登录服务
	userInfoQuery, err := UserInfoService(userId, token)

	//用户信息不存在，报错返回
	if err != nil {
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: repository.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	//用户信息存在，返回用户信息
	c.JSON(http.StatusOK, UserInfoResponse{
		Response: repository.Response{StatusCode: 0},
		UserInf:  userInfoQuery,
	})
}
func UserInfoService(id string, token string) (UserInfoQuery, error) {
	newUserInfoQuery := UserInfoQuery{}
	//将userId从string类型转为int类型
	userId, err := strconv.ParseInt(id, 10, 64)
	//类型转换失败
	if err != nil {
		return newUserInfoQuery, err
	}
	//根据userId获取用户信息
	var user *repository.User
	user, err = repository.NewUserDaoInstance().QueryUserById(userId)
	//未找到userId
	if err != nil {
		return newUserInfoQuery, err
	}
	newUserInfoQuery.UserId = user.Id
	newUserInfoQuery.UserName = user.Name
	newUserInfoQuery.FollowCount = user.FollowerCount
	newUserInfoQuery.FollowerCount = user.FollowerCount
	newUserInfoQuery.IsFollow = false
	return newUserInfoQuery, nil

}
