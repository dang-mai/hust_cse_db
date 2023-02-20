package controller

import (
	"github.com/gin-gonic/gin"
	"hust/cse/db/models"
	"net/http"
	"strconv"
)

func StudentInsert(ctx *gin.Context) {
	var newStu models.Student
	// 绑定数据
	err := ctx.ShouldBind(&newStu)
	if err != nil {
		return
	}
	// 插入
	err, rows := models.StudentInsertDB(&newStu)
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

func StudentQuery(ctx *gin.Context) {
	// 绑定数据
	var stuQuery models.StudentQuery
	err := ctx.ShouldBindQuery(&stuQuery)
	if err != nil {
		return
	}
	// 查询
	var totalRows int64
	students := make([]models.Student, 0)
	err, rows := models.StudentQueryDB(&stuQuery, &students, &totalRows)
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
			"data":         students,
		})
	}
}

func StudentUpdate(ctx *gin.Context) {
	// 绑定数据
	var newStu models.Student
	err := ctx.ShouldBind(&newStu)
	if err != nil {
		return
	}
	// 更新
	err, rows := models.StudentUpdateDB(&newStu)
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

func StudentDelete(ctx *gin.Context) {
	// 绑定数据
	var id int
	id, _ = strconv.Atoi(ctx.Query("id"))
	// 删除
	err, rows := models.StudentDeleteDB(id)
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
