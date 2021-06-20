package main

import (
	"University-Information-Website/model"

	"gopkg.in/mgo.v2/bson"
)

func testCourse() {
	userId := "571094e2976aeb1df982ad4e"
	course := model.Courses{
		"0",
		bson.NewObjectId().Hex(),
		"goland",
		"goland入门学习",
		nil,
		"",
		"2006-01-02 15:04:05",
	}

	model.CreateCourse(&course, userId)
}

func testDeleteCourse() {
	id := "60cf447db1f7f2279800edda"
	model.DeleteCourses(id)
}

func main() {
	model.InitDb()
	model.InitModel()
	//testCourse()
	testDeleteCourse()
	//routes.InitRouter()
}
