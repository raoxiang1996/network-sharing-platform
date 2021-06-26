package main

import (
	"University-Information-Website/model"
	"University-Information-Website/utils/errmsg"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

func testInsertCourse() {
	userId := "571094e2976aeb1df982ad4e"
	course := model.Courses{
		"0",
		bson.NewObjectId().Hex(),
		"goland",
		"goland入门学习",
		make([]model.Lesson, 0),
		"",
		"2006-01-02 15:04:05",
	}

	model.CreateCourse(&course, userId)
}

func testUpdateCourse() {
	//id := "60cffca2b1f7f220d4fe0cf6"
	//model.UpdateCourse(id, "golang", "golang入门")
}

func testDeleteCourse() {
	id := "60cf447db1f7f2279800edda"
	model.DeleteCourses(id)
}

func testInsertLesson() {
	id := "60d1e8642eb5a44d8cd87efa"
	lesson := model.Lesson{
		bson.NewObjectId().Hex(),
		id,
		"root/lesson2",
		"2006-01-02 15:04:05",
	}
	model.InsertLesson(&lesson, id)
}

func testEditLesson() {
	id := "60d1e8642eb5a44d8cd87efa"
	lesson := model.Lesson{
		"60d4366e2eb5a40a78f0771e",
		id,
		"root/lesson2/edit",
		"",
	}
	model.UpdateLesson(&lesson, lesson.ID)
}

func testDeleteLesson() {
	courseId := "60d03c21b1f7f20d38ed3a16"
	lessonId := "60d03f22b1f7f21cb8ba2216"
	model.DeleteLesson(courseId, lessonId)
}

func testInsertComment() {
	courseId := "60d03c21b1f7f20d38ed3a16"
	lessonId := "60d03f22b1f7f21cb8ba2216"
	sc := model.SingleComment{
		"",
		"60d04db7b1f7f22e38216b6a",
		"raoxiang",
		"测试评论",
		"2006-01-02 15:04:05",
	}
	model.InsertComment(&sc, courseId, lessonId)
}

func testCreateComments() {
	courseId := "60d03c21b1f7f20d38ed3a16"
	lessonId := "60d03f22b1f7f21cb8ba2216"
	model.CreateComments(courseId, lessonId)
}

func testDeleteComment() {
	courseId := "60d03c21b1f7f20d38ed3a16"
	lessonId := "60d03f22b1f7f21cb8ba2216"
	commentId := "60d0587b2eb5a45f78af294d"
	model.DeleteComment(courseId, lessonId, commentId)
}

func testInsertUser() {
	data := model.User{
		bson.NewObjectId().Hex(),
		"raoxiang",
		"raoxiang",
		1,
	}
	model.InsertUser(&data)
}

func print() {
	fmt.Println("hello world!")
}

func testTiming() {
	model.Timing()
	res, _ := model.Retrieve()
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i].CourseName, " ", res[i].CourseId)
	}
}

func testGetAllCourse() {
	courses, _ := model.GetAllCourse(0, 4, true, 1, "createtime")
	for i := 0; i < len(courses); i++ {
		fmt.Println(courses[i].ID + "  " + courses[i].CourseName + "  " + courses[i].Introduction + "  " + courses[i].Createtime + "  " + courses[i].UserId)
	}
}

func testSearchCourse() {
	courses, err := model.SearchCourse(0, 30, "深入理解计算机系统")
	if err != errmsg.SUCCESS {
		fmt.Println("error")
	}
	for i := 0; i < len(courses); i++ {
		fmt.Println(courses[i].CourseId + " " + courses[i].CourseName + " " + courses[i].CourseImage + " " + courses[i].CourseName)
	}
}

func main() {

	//testCreateComments()
	//testInsertComment()
	//testDeleteComment()
	//testDeleteLesson()
	model.InitDb()
	model.InitModel()
	testSearchCourse()
	//testTiming()
	//TestGetAllCourse()
	//now, _ := common.GetNowTime()
	//fmt.Println(now)
	//testInsertUser()
	//testDeleteLesson()
	//testInsertLesson()
	testEditLesson()
	//testUpdateCourse()
	//testInsertCourse()
	//testDeleteCourse()
	//routes.InitRouter()
}
