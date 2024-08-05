package model

type Department struct {
	ID             int    `json:"id" db:"id" gorm:"column:id;primarykey;autoIncrement"`
	DepartmentId   string `json:"departmentId" db:"department_id" gorm:"column:department_id"`
	DepartmentName string `json:"departmentName" db:"department_name" gorm:"column:department_name"`
}
