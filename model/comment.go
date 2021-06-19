package model

import (
	"University-Information-Website/utils/common"
	"University-Information-Website/utils/errmsg"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"

	"gopkg.in/mgo.v2/bson"
)

type SingleComment struct {
	ComentID   bson.ObjectId   `bson:"coment_id"`
	Userid     bson.ObjectId   `bson:"userid"`
	Username   string          `bson:"username"`
	Content    string          `bson:"content"`
	Createtime common.JsonTime `bson:"createtime" json:"createtime"`
}

type Comments struct {
	CoursesID   bson.ObjectId   `bson:"courses_id"`
	LessonID    bson.ObjectId   `bson:"lesson_id"`
	allComments []SingleComment `bson:"all_comments"`
}

var commentCollection *mongo.Collection = nil

// 添加评论
func InsertComment(data *SingleComment, coursesId string, courseId string) int {
	return errmsg.SUCCESS
}

//删除评论
func DeleteComment(coursesId string, courseId string, commentId string) int {
	return errmsg.SUCCESS
}

func commentInit() {
	if db != nil {
		if commentCollection == nil {
			userCollection = db.Collection("comment")
		} else {
			fmt.Println("comment collection has inited")
			log.Fatal("comment collection has inited")
		}
	} else {
		fmt.Println("mongodb comment collection init error, db has not inited")
		log.Fatal("mongodb comment collection init error, db has not inited")
	}
}
