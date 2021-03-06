package model

type Student struct {
	ID      string `gorm:"primaryKey;comment:证件号;size:20"`
	UID     string `gorm:"not null;uniqueIndex;comment:学号;size:20"`
	Name    string `gorm:"not null;Index;comment:姓名;size:15"`
	Faculty string `gorm:"comment:学院;size:30"`
	Class   string `gorm:"comment:班级;size:50"`
	Campus  string `gorm:"comment:校区;size:10"`
	House   string `gorm:"comment:寝室楼;size:15"`
	Room    string `gorm:"comment:寝室号;size:15"`
	Bed     int8   `gorm:"comment:床号"`
}
