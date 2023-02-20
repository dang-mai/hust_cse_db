package models

import (
	"gorm.io/gorm"
	"hust/cse/db/mysql"
)

type Student struct {
	gorm.Model
	Sno          string `gorm:"type:varchar(20);unique;" json:"sno" form:"sno" uri:"sno"`
	Sname        string `gorm:"type:varchar(20);" json:"sname" form:"sname" uri:"sname"`
	Ssex         string `gorm:"type:varchar(10);" json:"ssex" form:"ssex" uri:"ssex"`
	Sage         uint16 `gorm:"type:smallint;" json:"sage" form:"sage" uri:"sage"`
	Sdept        string `gorm:"type:varchar(20);" json:"sdept" form:"sdept" uri:"sdept"`
	Sscholarship string `gorm:"type:varchar(10);" json:"sscholarship" form:"sscholarship" uri:"sscholarship"`
}

type StudentQuery struct {
	PageNum  int    `form:"pageNum" uri:"pageNum"`
	PageSize int    `form:"pageSize" uri:"pageSize"`
	Search   string `form:"search" uri:"search"`
}

func StudentQueryDB(stuQuery *StudentQuery, students *[]Student, rows *int64) (error, int64) {
	mysql.DB.Model(&Student{}).Count(rows)
	if stuQuery.PageNum > 0 && stuQuery.PageSize > 0 {
		tx := mysql.DB.Limit(stuQuery.PageSize).Offset((stuQuery.PageNum - 1) * stuQuery.PageSize).Order("sno").Find(&students)
		return tx.Error, tx.RowsAffected
	} else {
		return nil, 0
	}
}

func StudentInsertDB(newStu *Student) (error, int64) {
	tx := mysql.DB.Create(&newStu)
	return tx.Error, tx.RowsAffected
}

func StudentUpdateDB(newStu *Student) (error, int64) {
	tx := mysql.DB.Model(&Student{}).Where("id = ?", newStu.ID).Save(newStu)
	return tx.Error, tx.RowsAffected
}

func StudentDeleteDB(id int) (error, int64) {
	tx := mysql.DB.Unscoped().Delete(&Student{}, id)
	return tx.Error, tx.RowsAffected
}
