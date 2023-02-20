package models

import (
	"gorm.io/gorm"
	"hust/cse/db/mysql"
)

type SC struct {
	gorm.Model
	Sno   string `gorm:"type:varchar(20);" json:"sno" form:"sno" uri:"sno"`
	Cno   string `gorm:"type:varchar(20);" json:"cno" form:"cno" uri:"cno"`
	Grade uint16 `gorm:"type:smallint;" json:"grade" form:"grade" uri:"grade"`
}

type SCQuery struct {
	PageNum  int    `form:"pageNum" uri:"pageNum"`
	PageSize int    `form:"pageSize" uri:"pageSize"`
	Search   string `form:"search" uri:"search"`
}

func SCQueryDB(scQuery *SCQuery, scs *[]SC, rows *int64) (error, int64) {
	mysql.DB.Model(&SC{}).Count(rows)
	if scQuery.PageNum > 0 && scQuery.PageSize > 0 {
		tx := mysql.DB.Limit(scQuery.PageSize).Offset((scQuery.PageNum - 1) * scQuery.PageSize).Order("sno").Find(&scs)
		return tx.Error, tx.RowsAffected
	} else {
		return nil, 0
	}
}

func SCInsertDB(newSC *SC) (error, int64) {
	tx := mysql.DB.Create(&newSC)
	return tx.Error, tx.RowsAffected
}

func SCUpdateDB(newSC *SC) (error, int64) {
	tx := mysql.DB.Model(&SC{}).Where("id = ?", newSC.ID).Save(newSC)
	return tx.Error, tx.RowsAffected
}

func SCDeleteDB(id int) (error, int64) {
	tx := mysql.DB.Unscoped().Delete(&SC{}, id)
	return tx.Error, tx.RowsAffected
}
