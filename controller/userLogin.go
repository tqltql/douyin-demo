package controller

import (
	"douyin-demo/repository"
	"douyin-demo/service"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserLoginResponse struct {
	repository.Response
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

func UserLogin(c *gin.Context) {
	//查询参数
	username := c.Query("username")
	password := c.Query("password")

	//登录服务
	userLoginResponse, err := UserLoginService(username, password)

	//用户名或密码不正确，报错返回
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: repository.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	//用户名密码匹配成功，返回UserId和Token
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: repository.Response{StatusCode: 0},
		UserId:   userLoginResponse.UserId,
		Token:    userLoginResponse.Token,
	})
}
func UserLoginService(userName string, password string) (UserLoginResponse, error) {
	newUserLoginResponse := UserLoginResponse{}

	//校验用户名和密码的合法性
	error := service.CheckUser(userName, password)
	//不合法报错
	if error != nil {
		return newUserLoginResponse, error
	}
	//查找用户名对应的的账号信息
	var user *repository.User
	user, error = repository.NewUserDaoInstance().QueryPasswordByUserName(userName)
	//未找到用户名
	if error != nil {
		return newUserLoginResponse, error
	}
	//判断密码是否正确
	if user.Password != password {
		return newUserLoginResponse, errors.New("密码错误")
	}
	newUserLoginResponse.UserId = user.Id
	newUserLoginResponse.Token = user.Token
	return newUserLoginResponse, nil

}
