package main

import (
	"University-Information-Website/model"

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
	id := "60cffca2b1f7f220d4fe0cf6"
	model.UpdateCourse(id, "golang", "golang入门")
}

func testDeleteCourse() {
	id := "60cf447db1f7f2279800edda"
	model.DeleteCourses(id)
}

func testInsertLesson() {
	id := "60d03c21b1f7f20d38ed3a16"
	lesson := model.Lesson{
		bson.NewObjectId().Hex(),
		id,
		"root/lesson",
		"2006-01-02 15:04:05",
	}
	model.InsertLesson(&lesson, id)
}

func testDeleteLesson() {
	courseId := "60d03c21b1f7f20d38ed3a16"
	lessonId := "60d03f22b1f7f21cb8ba2216"
	model.DeleteLesson(courseId, lessonId)
}

func main() {
	model.InitDb()
	model.InitModel()
	testDeleteLesson()
	//testInsertLesson()
	//testUpdateCourse()
	//testInsertCourse()
	//testDeleteCourse()
	//routes.InitRouter()
}
