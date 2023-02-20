package controller

import (
	"github.com/gin-gonic/gin"
	"hust/cse/db/models"
	"net/http"
	"strconv"
)

func SCInsert(ctx *gin.Context) {
	// 绑定数据
	var newSC models.SC
	err := ctx.ShouldBind(&newSC)
	if err != nil {
		return
	}
	// 插入
	err, rows := models.SCInsertDB(&newSC)
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

func SCQuery(ctx *gin.Context) {
	// 绑定数据
	var scQuery models.SCQuery
	err := ctx.ShouldBindQuery(&scQuery)
	if err != nil {
		return
	}
	// 查询
	var totalRows int64
	scs := make([]models.SC, 0)
	err, rows := models.SCQueryDB(&scQuery, &scs, &totalRows)
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
			"data":         scs,
		})
	}
}

func SCUpdate(ctx *gin.Context) {
	// 绑定数据
	var newSC models.SC
	err := ctx.ShouldBind(&newSC)
	if err != nil {
		return
	}
	// 更新
	err, rows := models.SCUpdateDB(&newSC)
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

func SCDelete(ctx *gin.Context) {
	// 绑定数据
	var id int
	id, _ = strconv.Atoi(ctx.Query("id"))
	// 删除
	err, rows := models.SCDeleteDB(id)
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
