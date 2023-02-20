package models

import (
	"gorm.io/gorm"
	"hust/cse/db/mysql"
)

type Course struct {
	gorm.Model
	Cno     string `gorm:"type:varchar(20);unique;" json:"cno" form:"cno" uri:"cno"`
	Cname   string `gorm:"type:varchar(40);" json:"cname" form:"cname" uri:"cname"`
	Cpno    string `gorm:"type:varchar(10);" json:"cpno" form:"cpno" uri:"cpno"`
	Ccredit uint16 `gorm:"type:smallint;" json:"ccredit" form:"ccredit" uri:"ccredit"`
}

type CourseQuery struct {
	PageNum  int    `form:"pageNum" uri:"pageNum"`
	PageSize int    `form:"pageSize" uri:"pageSize"`
	Filter   string `form:"filter" uri:"filter"`
}

func CourseQueryDB(couQuery *CourseQuery, courses *[]Course, rows *int64) (error, int64) {
	if couQuery.Filter == "true" {
		tx := mysql.DB.Model(&Course{}).Raw("SELECT * FROM gorm_course WHERE NOT EXISTS (SELECT * FROM gorm_sc WHERE gorm_course.cno = gorm_sc.cno);").Scan(&courses)
		*rows = tx.RowsAffected
		return tx.Error, tx.RowsAffected
	} else {
		mysql.DB.Model(&Course{}).Count(rows)
		if couQuery.PageNum > 0 && couQuery.PageSize > 0 {
			tx := mysql.DB.Limit(couQuery.PageSize).Offset((couQuery.PageNum - 1) * couQuery.PageSize).Order("cno").Find(&courses)
			return tx.Error, tx.RowsAffected
		} else {
			return nil, 0
		}
	}

}

func CourseInsertDB(newCou *Course) (error, int64) {
	tx := mysql.DB.Create(&newCou)
	return tx.Error, tx.RowsAffected
}

func CourseUpdateDB(newCou *Course) (error, int64) {
	tx := mysql.DB.Model(&Course{}).Where("id = ?", newCou.ID).Save(newCou)
	return tx.Error, tx.RowsAffected
}

func CourseDeleteDB(id int) (error, int64) {
	tx := mysql.DB.Unscoped().Delete(&Course{}, id)
	return tx.Error, tx.RowsAffected
}
