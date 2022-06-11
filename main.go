package main

import (
	"douyin-demo/controller"
	"douyin-demo/repository"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	if err := Init(); err != nil {
		os.Exit(-1)
	}
	r := gin.Default()
	r.Static("/static", "./")

	//douyinGroup := r.Group("/douyin")
	//{
	//	//feed
	//	douyinGroup.GET("/feed", controller.Feed)
	//	//user
	//	userGroup := r.Group("/user")
	//	{
	//
	//		userGroup.POST("/register", controller.UserRegister)
	//		userGroup.POST("/login", controller.UserLogin)
	//		userGroup.GET("/", controller.UserInfo)
	//
	//	}
	//	//publish
	//	publishGroup := r.Group("/publish")
	//	{
	//		publishGroup.POST("/action", controller.PublishAction)
	//		publishGroup.GET("/list", controller.PublishList)
	//	}
	//
	//	//favourite
	//	favouriteGroup := r.Group("/favourite")
	//	{
	//		favouriteGroup.POST("/action", controller.FavoriteAction)
	//		favouriteGroup.GET("/list", controller.FavoriteList)
	//
	//	}
	//	//comment
	//	commentGroup := r.Group("/comment")
	//	{
	//		commentGroup.POST("/action", controller.CommentAction)
	//		commentGroup.GET("/list", controller.CommentList)
	//	}
	//
	//	//relation
	//	relationGroup := r.Group("/relation")
	//	{
	//		relationGroup.POST("/action", controller.RelationAction)
	//		relationGroup.GET("/follow/list", controller.RelationFollowList)
	//		relationGroup.GET("/follower/list", controller.RelationFollowerList)
	//	}
	//}

	r.POST("/douyin/user/login/", controller.UserLogin)
	r.GET("/douyin/user/", controller.UserInfo)
	r.POST("/douyin/user/register/", controller.UserRegister)
	r.GET("/douyin/feed/", controller.Feed)
	err := r.Run(":8080")
	if err != nil {
		return
	}
}

func Init() error {
	if err := repository.Init(); err != nil {
		return err
	}
	return nil
}
