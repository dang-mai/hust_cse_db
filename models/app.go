package models

import (
	"hust/cse/db/mysql"
)

type InfoGradeQuery struct {
	Sno      string `form:"sno" uri:"sno"`
	PageNum  int    `form:"pageNum" uri:"pageNum"`
	PageSize int    `form:"pageSize" uri:"pageSize"`
}

type DeptRank struct {
	Sno   string `form:"sno" uri:"sno"`
	Sname string `form:"sname" uri:"sname"`
	Cno   string `form:"cno" uri:"cno"`
	Cname string `form:"cname" uri:"cname"`
	Grade int    `form:"grade" uri:"grade"`
}

type DeptInfo struct {
	Dept      string  `form:"dept" uri:"dept"`
	AvgGrade  float64 `form:"avgGrade" uri:"avgGrade"`
	MaxGrade  float64 `form:"maxGrade" uri:"maxGrade"`
	MinGrade  float64 `form:"minGrade" uri:"minGrade"`
	GreatRate float64 `form:"greatRate" uri:"greatRate"`
	FailNum   int     `form:"failNum" uri:"failNum"`
}

func InfoGradeQueryDB(infoGradeQuery *InfoGradeQuery, student *Student, scs *[]SC, scRows *int64) (error, int64) {
	tx := mysql.DB.Model(&Student{}).Where("sno = ?", infoGradeQuery.Sno).Take(&student)
	if tx.Error == nil {
		if infoGradeQuery.PageNum > 0 && infoGradeQuery.PageSize > 0 {
			tx = mysql.DB.Model(&SC{}).Where("sno = ?", infoGradeQuery.Sno).Count(scRows).Limit(infoGradeQuery.PageSize).Offset((infoGradeQuery.PageNum - 1) * infoGradeQuery.PageSize).Order("cno").Find(&scs)
		}
	}
	return tx.Error, tx.RowsAffected
}

func DeptRankQueryDB(dept string, deptRank *[]DeptRank, rows *int64) (error, int64) {
	tx := mysql.DB.Model(&Course{}).Raw("SELECT gorm_sc.sno, sname, gorm_sc.cno, cname, grade FROM gorm_student, gorm_course, gorm_sc WHERE gorm_student.sno=gorm_sc.sno AND gorm_sc.cno=gorm_course.cno AND gorm_student.sdept = ? ORDER BY grade DESC;", dept).Scan(&deptRank)
	*rows = tx.RowsAffected
	return tx.Error, tx.RowsAffected
}

func DeptInfoQueryDB(deptInfo *[]DeptInfo, rows *int64) (error, int64) {
	tx := mysql.DB.Model(&Course{}).Raw("SELECT Sdept Dept, AVG(Grade) AvgGrade, MAX(Grade) MaxGrade, MIN(Grade) MinGrade, " +
		"sum(IF(Grade > 90, 1, 0)) * 100.0 / count(*) * 1.0 GreatRate, sum(IF(Grade < 60, 1, 0)) FailNum " +
		"FROM gorm_sc, gorm_student WHERE gorm_student.sno = gorm_sc.sno GROUP BY gorm_student.sdept").Scan(&deptInfo)
	*rows = tx.RowsAffected
	return tx.Error, tx.RowsAffected
}
