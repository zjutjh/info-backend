package handler

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"info/data"
	"info/model"
)

func ReaInfo(path string, passwd string) {
	InitDB()
	var f *excelize.File
	var err error
	if passwd != "" {
		f, err = excelize.OpenFile(path, excelize.Options{Password: passwd})
	} else {
		f, err = excelize.OpenFile(path)
	}
	if err != nil {
		panic(err)
	}
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		panic(err)
	}
	for i, row := range rows {
		if i > 0 {
			stu := model.Student{
				Name:    row[0],
				ID:      row[2],
				Campus:  row[3],
				Faculty: row[4],
				Class:   row[6],
				UID:     row[7],
			}
			result := data.DB.Select("Name", "ID", "Campus", "Faculty", "Class", "UID").Create(&stu)
			if result.Error != nil {
				fmt.Println(result.Error.Error())
			}
		}
	}
}