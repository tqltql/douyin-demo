package controller

import (
	"douyin-demo/repository"
	"douyin-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserRegisterResponse struct {
	repository.Response
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

func UserRegister(c *gin.Context) {
	//查询参数
	username := c.Query("username")
	password := c.Query("password")

	//注册服务
	userRegisterResponse, err := UserRegisterService(username, password)

	//报错返回
	if err != nil {
		c.JSON(http.StatusOK, UserRegisterResponse{
			Response: repository.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	//注册成功
	c.JSON(http.StatusOK, UserRegisterResponse{
		Response: repository.Response{StatusCode: 0},
		UserId:   userRegisterResponse.UserId,
		Token:    userRegisterResponse.Token,
	})
}
func UserRegisterService(userName string, password string) (UserRegisterResponse, error) {
	newUserRegisterResponse := UserRegisterResponse{}

	//校验用户名和密码的合法性
	error := service.CheckUser(userName, password)
	//不合法报错
	if error != nil {
		return newUserRegisterResponse, error
	}
	//创建用户
	error = repository.NewUserDaoInstance().CreateUser(userName, password)
	if error != nil {
		return newUserRegisterResponse, error
	}

	//查找用户名对应的的账号信息
	var user *repository.User

	user, error = repository.NewUserDaoInstance().QueryPasswordByUserName(userName)
	//未找到用户名
	if error != nil {
		return newUserRegisterResponse, error
	}
	newUserRegisterResponse.UserId = user.Id
	newUserRegisterResponse.Token = user.Token
	return newUserRegisterResponse, nil
}
