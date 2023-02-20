package controller

import (
	"github.com/gin-gonic/gin"
	"hust/cse/db/models"
	"net/http"
)

func InfoGradeQuery(ctx *gin.Context) {
	// 接收数据
	var infoGradeQuery models.InfoGradeQuery
	err := ctx.ShouldBindQuery(&infoGradeQuery)
	if err != nil {
		return
	}
	// 查询
	var totalRows int64
	var student models.Student
	scs := make([]models.SC, 0)
	err, rows := models.InfoGradeQueryDB(&infoGradeQuery, &student, &scs, &totalRows)
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
			"infoData":     student,
			"gradeData":    scs,
		})
	}
}

func DeptInfoQuery(ctx *gin.Context) {
	// 查询
	var totalRows int64
	deptInfo := make([]models.DeptInfo, 0)
	err, rows := models.DeptInfoQueryDB(&deptInfo, &totalRows)
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
			"data":         deptInfo,
		})
	}
}

func DeptRankQuery(ctx *gin.Context) {
	// 绑定数据
	dept := ctx.Query("dept")
	// 查询
	var totalRows int64
	deptRank := make([]models.DeptRank, 0)
	err, rows := models.DeptRankQueryDB(dept, &deptRank, &totalRows)
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
			"data":         deptRank,
		})
	}
}
