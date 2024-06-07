package controllers

import (
	"net/http"
	"strconv"

	"github.com/lmhuong711/go-go-be/entities"
	"github.com/lmhuong711/go-go-be/services"
	"github.com/lmhuong711/go-go-be/utils"

	"github.com/gofiber/fiber/v2"
)

func standardizeStudent(st *entities.Students) *entities.Students {
	_st := entities.Students{}
	if st.ID > 0 {
		_st.ID = st.ID
	}
	if st.FirstName != "" {
		_st.FirstName = st.FirstName
	}
	if st.LastName != "" {
		_st.LastName = st.LastName
	}
	if st.Phone != "" && utils.StandardizePhone(st.Phone) != "" {
		_st.Phone = st.Phone
	}
	if st.Email != "" && utils.StandardizeEmail(st.Email) != "" {
		_st.Email = st.Email
	}
	if !st.Birthday.IsZero() {
		_st.Birthday = st.Birthday
	}
	if st.Address != "" {
		_st.Address = st.Address
	}
	if st.Gender == 0 || st.Gender == 1 {
		_st.Gender = st.Gender
	}
	if st.University != "" {
		_st.University = st.University
	}
	return &_st
}
func CreateStudent(ctx *fiber.Ctx) error {
	resp := fiber.Map{
		"success": false,
		"message": "Student not created",
		"data":    nil,
	}

	var newStudent *entities.Students
	if err := ctx.BodyParser(&newStudent); err != nil {
		resp["message"] = err.Error()
		return ctx.Status(http.StatusBadGateway).JSON(resp)
	}
	newStudent.ID = 0

	newStudent.Phone = utils.StandardizePhone(newStudent.Phone)
	newStudent.Email = utils.StandardizeEmail(newStudent.Email)
	if newStudent.Phone == "" || newStudent.Email == "" {
		resp["message"] = "Phone or Email is invalid or duplicated"
		return ctx.Status(http.StatusBadGateway).JSON(resp)
	}

	student, err := services.CreateStudent(newStudent)
	if err != nil {
		resp["message"] = err.Error()
		return ctx.Status(http.StatusBadGateway).JSON(resp)
	}

	return ctx.JSON(fiber.Map{
		"success": true,
		"message": "Student created successfully",
		"data":    student,
	})
}

func GetStudent(ctx *fiber.Ctx) error {
	resp := fiber.Map{
		"success": false,
		"message": "Student not found",
		"data":    nil,
	}
	st, err := services.GetStudent(ctx.Params("id"))
	if err != nil {
		resp["message"] = err.Error()
		return ctx.Status(http.StatusNotFound).JSON(resp)
	}
	return ctx.JSON(fiber.Map{
		"success": true,
		"message": "Student found",
		"data":    st,
	})
}

func GetStudents(ctx *fiber.Ctx) error {
	resp := fiber.Map{
		"success": false,
		"message": "Students not found",
		"data":    nil,
	}
	limit, offset := 10, 0
	if _limit, e := strconv.Atoi(ctx.Query("limit")); e == nil && _limit > 0 {
		limit = _limit
	}
	if _offset, e := strconv.Atoi(ctx.Query("offset")); e == nil && _offset > 0 {
		offset = _offset
	}
	st, count, err := services.GetStudents(limit, offset)
	if err != nil || count == 0 {
		return ctx.Status(http.StatusNotFound).JSON(resp)
	}
	return ctx.JSON(fiber.Map{
		"success": true,
		"message": "Students found",
		"data":    st,
		"count":   count,
	})
}

func UpdateStudent(ctx *fiber.Ctx) error {
	resp := fiber.Map{
		"success": false,
		"message": "Student not updated",
		"data":    nil,
	}
	var st *entities.Students
	if err := ctx.BodyParser(&st); err != nil || strconv.FormatUint(uint64(st.ID), 10) != ctx.Params("id") {
		resp["message"] = "ID not match"
		return ctx.Status(http.StatusBadGateway).JSON(resp)
	}
	_st := standardizeStudent(st)
	st, err := services.UpdateStudent(_st, ctx.Params("id"))
	if err != nil {
		resp["message"] = err.Error()
		return ctx.Status(http.StatusBadGateway).JSON(resp)
	}

	return ctx.JSON(fiber.Map{
		"success": true,
		"message": "Student updated successfully",
		"data":    st,
	})
}

func DeleteStudent(ctx *fiber.Ctx) error {
	resp := fiber.Map{
		"success": false,
		"message": "Student not deleted",
		"data":    nil,
	}
	_, err := services.DeleteStudent(ctx.Params("id"))
	if err != nil {
		resp["message"] = err.Error()
		return ctx.Status(http.StatusNotFound).JSON(resp)
	}
	return ctx.JSON(fiber.Map{
		"success": true,
		"message": "Student deleted successfully",
		"data":    nil,
	})
}
