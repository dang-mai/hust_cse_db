package router

import (
	"github.com/gin-gonic/gin"
	"hust/cse/db/controller"
	"hust/cse/db/middlewares"
	"hust/cse/db/setting"
)

func SetupRouter() *gin.Engine {
	// 是否为Release模式
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	// 使用默认配置
	r := gin.Default()
	// 使用中间件
	r.Use(middlewares.Header())
	// 加载静态文件
	r.Static("/js", "./static/js")
	r.Static("/css", "./static/css")
	r.Static("/img", "./static/img")
	// 加载页面文件
	r.LoadHTMLGlob("templates/*")
	// 根路由 前端页面
	r.GET("/", controller.IndexHandler)
	// API接口
	studentGroup := r.Group("api/student")
	{ // Student表接口
		studentGroup.GET("/query", controller.StudentQuery)
		studentGroup.POST("/insert", controller.StudentInsert)
		studentGroup.PUT("/update", controller.StudentUpdate)
		studentGroup.DELETE("/delete", controller.StudentDelete)
	}
	courseGroup := r.Group("api/course")
	{ // Course表接口
		courseGroup.GET("/query", controller.CourseQuery)
		courseGroup.POST("/insert", controller.CourseInsert)
		courseGroup.PUT("/update", controller.CourseUpdate)
		courseGroup.DELETE("/delete", controller.CourseDelete)
	}
	scGroup := r.Group("api/sc")
	{ // SC表接口
		scGroup.GET("/query", controller.SCQuery)
		scGroup.POST("/insert", controller.SCInsert)
		scGroup.PUT("/update", controller.SCUpdate)
		scGroup.DELETE("/delete", controller.SCDelete)
	}
	appGroup := r.Group("api/app")
	{ // 功能接口
		appGroup.GET("/infograde", controller.InfoGradeQuery)
		appGroup.GET("/deptinfo", controller.DeptInfoQuery)
		appGroup.GET("/deptrank", controller.DeptRankQuery)
	}
	return r
}
