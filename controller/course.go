package controller

import (
	"github.com/gin-gonic/gin"
	"hust/cse/db/models"
	"net/http"
	"strconv"
)

func CourseInsert(ctx *gin.Context) {
	// 绑定数据
	var newCou models.Course
	err := ctx.ShouldBind(&newCou)
	if err != nil {
		return
	}
	// 插入
	err, rows := models.CourseInsertDB(&newCou)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  1,
			"msg":   "fail",
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":         0,
			"msg":          "success",
			"rowsAffected": rows,
		})
	}
}

func CourseQuery(ctx *gin.Context) {
	// 绑定数据
	var couQuery models.CourseQuery
	err := ctx.ShouldBindQuery(&couQuery)
	if err != nil {
		return
	}
	// 查询
	var totalRows int64
	courses := make([]models.Course, 0)
	err, rows := models.CourseQueryDB(&couQuery, &courses, &totalRows)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  1,
			"msg":   "fail",
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":         0,
			"msg":          "success",
			"rows":         totalRows,
			"rowsAffected": rows,
			"data":         courses,
		})
	}
}

func CourseUpdate(ctx *gin.Context) {
	// 绑定数据
	var newCou models.Course
	err := ctx.ShouldBind(&newCou)
	if err != nil {
		return
	}
	// 更新
	err, rows := models.CourseUpdateDB(&newCou)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  1,
			"msg":   "fail",
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":         0,
			"msg":          "success",
			"rowsAffected": rows,
		})
	}
}

func CourseDelete(ctx *gin.Context) {
	// 绑定数据
	var id int
	id, _ = strconv.Atoi(ctx.Query("id"))
	// 删除
	err, rows := models.CourseDeleteDB(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  1,
			"msg":   "fail",
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":         0,
			"msg":          "success",
			"rowsAffected": rows,
		})
	}
}
