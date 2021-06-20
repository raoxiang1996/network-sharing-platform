package model

import (
	"University-Information-Website/utils/errmsg"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"

	"gopkg.in/mgo.v2/bson"
)

type Lesson struct {
	ID         string `bson:"_id"`
	CoursesId  string `bson:"courses_id"`
	path       string `bson:"path"`
	uploadTime string `bson:"upload_time"  json:"createtime"`
}

type Courses struct {
	ID           string   `bson:"_id"`
	UserId       string   `bson:"user_id"`
	CourseName   string   `bson:"course_name"`
	Introduction string   `bson:"introduction"`
	Subject      []Lesson `bson:"subject"`
	Images       string   `bson:"images"`
	Createtime   string   `bson:"createtime"  json:"createtime"`
}

var coursesCollection *mongo.Collection = nil

// 添加课程
func CreateCourse(data *Courses, userId string) int {
	data.ID = bson.NewObjectId().Hex()
	data.UserId = userId
	newData := data
	//data.Subject[0].LessonID = bson.NewObjectId()
	//CreateCourseComment(data.CoursesID.String(), data.Subject[0].LessonID.String())
	insertResult, err := coursesCollection.InsertOne(context.TODO(), newData)
	if err != nil {
		fmt.Println("create a course fail")
		log.Fatal("create a course fail,", err)
		return errmsg.ERROR
	}
	fmt.Println("create a single document: ", insertResult.InsertedID.(string))
	return errmsg.SUCCESS
}

// 添加一节课课程
func InsertLesson(data *Lesson, coursesId string) int {
	data.ID = bson.NewObjectId().Hex()
	insertResult, err := coursesCollection.InsertOne(context.TODO(), data)
	if err != nil {
		fmt.Println("insert a lesson fail")
		log.Fatal("insert a lesson fail,", err)
		return errmsg.ERROR
	}
	fmt.Println("insert a single document: ", insertResult.InsertedID.(string))
	return errmsg.SUCCESS
}

//更新课程信息
func UpdateCourse(coursesId string, courseName string, introduction string) int {
	filter := bson.D{{"courses_id", coursesId}}
	update := bson.D{{"course_name", courseName}, {"introduction", introduction}}
	updateResult, err := coursesCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println("update a course fail")
		log.Fatal("update a course fail,", err)
		return errmsg.ERROR
	}
	fmt.Println("deleted a single comment: ", updateResult.UpsertedID.(string))
	return errmsg.SUCCESS
}

//删除一节课程
func DeleteLesson(coursesId string, lessonId string) int {
	filter := bson.D{{"courses_id", coursesId}}
	update := bson.D{
		{"$pull", bson.M{"subject": bson.D{{"comment_id", lessonId}}}},
	}
	insertResult, err := coursesCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println("delete a course fail")
		log.Fatal("delete a course fail,", err)
		return errmsg.ERROR
	}
	fmt.Println("deleted a single comment: ", insertResult.UpsertedID.(string))
	return errmsg.SUCCESS
}

//删除课程
func DeleteCourses(coursesId string) int {
	deleteResult, err := coursesCollection.DeleteOne(context.TODO(), bson.M{"_id": coursesId})
	if err != nil {
		log.Fatal("delete courses fail", err)
		return errmsg.ERROR
	}

	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
	return errmsg.SUCCESS
}

func courseInit() {
	if db != nil {
		if coursesCollection == nil {
			coursesCollection = db.Collection("course")
		} else {
			fmt.Println("course collection has inited")
			log.Fatal("course collection has inited")
		}
	} else {
		fmt.Println("mongodb course collection init error, db has not inited")
		log.Fatal("mongodb course collection init error, db has not inited")
	}
}
