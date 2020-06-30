package routers

import (
	"QUZHIYOU/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	// 初始化默认路由
	router := gin.Default()
	router.Static("/static", "./static")

	login := router.Group("/login")
	{
		login.GET("/getcode", api.WxLogin)                //获取openid
		login.POST("/getwxuserinfo", api.WxLoginUserInfo) //获取用户详情
	}

	v1 := router.Group("v1")
	{
		//首页活动列表
		v1.GET("/wechat/activity/selectActivityList", api.ActivityList)
		//活动详情信息
		v1.GET("/wechat/activity/selectActivicyInfoById", api.ActivityInfo)
		//获取code码
		v1.GET("/getqrcode", api.Getqrcode)
	}

	//动态首页

	home := router.Group("home")
	{

		home.POST("/homediarys", api.HomeList)
		home.GET("/classify", api.Classify)
		home.GET("/ad", api.GetAd)
		home.POST("/adddiary", api.PostAddDiary)
		home.POST("/upload", api.PostDiaryPic)
		home.GET("/getcommunity", api.Getcommunity)
		home.GET("/getsubtopic", api.Getsubtopic)
		home.POST("/likediary", api.PostDiaryLike)
		home.POST("/diaryinfo", api.GetDiaryInfo) //获取一条日记信息
	}

	//评论组
	comment := router.Group("comment")
	{
		comment.POST("/creatediarycomment", api.CreateComment)
		comment.POST("/getdiarycomment", api.GetComment)
		comment.POST("/likediarycomment", api.LikeDiaryComment)
	}

	return router

}
