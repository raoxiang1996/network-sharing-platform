package model

import (
	"University-Information-Website/utils/common"
	"University-Information-Website/utils/errmsg"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"

	"gopkg.in/mgo.v2/bson"
)

type Lesson struct {
	LessonID   bson.ObjectId   `bson:"lesson_id"`
	CoursesId  bson.ObjectId   `bson:"courses_id"`
	path       string          `bson:"path"`
	uploadTime common.JsonTime `bson:"upload_time"  json:"createtime"`
}

type Courses struct {
	CoursesID    bson.ObjectId   `bson:"courses_id"`
	Userid       bson.ObjectId   `bson:"userid"`
	CourseName   string          `bson:"course_name"`
	Introduction string          `bson:"introduction"`
	Subject      []Lesson        `bson:"subject"`
	Createtime   common.JsonTime `bson:"createtime"  json:"createtime"`
}

var coursesCollection *mongo.Collection = nil

// 添加课程
func InsertCourse(data *Lesson, courseId string, userId string) int {

	return errmsg.SUCCESS
}

//删除课程
func DeleteCourse(courseId string) int {

	return errmsg.SUCCESS
}

//删除课程
func DeleteCourses(coursesId string) int {
	deleteResult, err := coursesCollection.DeleteOne(context.TODO(), bson.M{"_id": bson.ObjectIdHex(coursesId)})
	if err != nil {
		log.Fatal("delete a user fail", err)
		return errmsg.ERROR
	}

	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
	return errmsg.SUCCESS
}

func courseInit() {
	if db != nil {
		if coursesCollection == nil {
			userCollection = db.Collection("course")
		} else {
			fmt.Println("course collection has inited")
			log.Fatal("course collection has inited")
		}
	} else {
		fmt.Println("mongodb course collection init error, db has not inited")
		log.Fatal("mongodb course collection init error, db has not inited")
	}
}
