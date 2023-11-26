package model

type Country struct {
	ID   uint   `gorm:"primaryKey;column:idx"`
	Name string `gorm:"type:varchar(255)"` //;column:idx ใส่ได้
}

// TableName specifies the custom table name for the Country struct.
func (Country) TableName() string {
	return "country"
}
