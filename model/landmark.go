package model

type Landmark struct {
	Id        int     `gorm:"column:idx;type:int(11);primary_key;column:idx"`
	Name      string  `gorm:"column:name;type:varchar(255);NOT NULL"`
	CountryID int     `gorm:"column:country;type:int(11);NOT NULL"`
	Country   Country `gorm:"foreignKey:CountryID"`
	Detail    string  `gorm:"column:detail;type:varchar(2000);NOT NULL"`
	Url       string  `gorm:"column:url;type:varchar(1000);NOT NULL"`
}

func (m *Landmark) TableName() string {
	return "landmark"
}
