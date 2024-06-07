package services

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/lmhuong711/go-go-be/database"
	"github.com/lmhuong711/go-go-be/entities"
	"github.com/lmhuong711/go-go-be/utils"

	"gorm.io/gorm"
)

func CreateStudent(student *entities.Students) (*entities.Students, error) {
	var err error
	tx := database.Conn.Create(&student)
	if tx.Error != nil {
		err = tx.Error
		utils.WarningLog.Println(err.Error())
	}
	return student, err
}

func GetStudent(studentId string) (*entities.Students, error) {
	var err error
	var student *entities.Students
	tx := database.Conn.Find(&student, studentId)
	if tx.Error != nil {
		err = tx.Error
		utils.WarningLog.Println(err.Error())
	}

	return student, err
}

func GetStudents(limit int, offset int) (*[]entities.Students, int64, error) {
	var err error
	var student *[]entities.Students
	tx := database.Conn.Limit(limit).Offset(offset).Find(&student)
	if tx.Error != nil {
		err = tx.Error
		utils.WarningLog.Println(err.Error())
	} else if *student == nil {
		err = errors.New("students not found")
	}
	var count int64
	tx = database.Conn.Model(&entities.Students{}).Count(&count)
	return student, count, err
}

func UpdateStudent(new *entities.Students, id string) (*entities.Students, error) {
	student, err := GetStudent(id)
	var studentMap map[string]interface{}
	studentByte, _ := json.Marshal(new)
	json.Unmarshal(studentByte, &studentMap)

	delete(studentMap, "id")
	delete(studentMap, "created_at")
	delete(studentMap, "updated_at")
	if studentMap["birthday"] == "0001-01-01T00:00:00Z" {
		studentMap["birthday"] = gorm.Expr("NULL")
	}
	if studentMap["gender"] == 0 {
		studentMap["gender"] = gorm.Expr("NULL")
	}
	log.Printf("studentMap: %v", studentMap)

	if err == nil {
		tx := database.Conn.Model(entities.Students{}).Omit("created_at").Omit("updated_at").Omit("id").Where("id = ?", id).Updates(studentMap).Where("id = ?", id)
		if tx.Error != nil {
			err = tx.Error
			utils.WarningLog.Println(err.Error())
		}
	}
	return student, err
}

func DeleteStudent(id string) (*entities.Students, error) {
	st, err := GetStudent(id)
	if err == nil {
		tx := database.Conn.Delete(st)
		if tx.Error != nil {
			err = tx.Error
			utils.WarningLog.Println(err.Error())
		}
	}
	return st, err
}
