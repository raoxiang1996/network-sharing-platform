package model

import (
	"University-Information-Website/utils/errmsg"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"

	"gopkg.in/mgo.v2/bson"
)

type SingleComment struct {
	ID         string `bson:"_id"`
	CommentID  string `bson:"_id"`
	UserId     string `bson:"user_id"`
	Username   string `bson:"username"`
	Content    string `bson:"content"`
	Createtime string `bson:"createtime" json:"createtime"`
}

type Comments struct {
	ID          string          `bson:"_id"`
	CoursesID   string          `bson:"courses_id"`
	LessonID    string          `bson:"lesson_id"`
	allComments []SingleComment `bson:"all_comments"`
}

var commentCollection *mongo.Collection = nil

// 添加评论
func InsertComment(data *SingleComment, coursesId string, lessonId string) int {
	data.CommentID = bson.NewObjectId().Hex()
	filter := bson.D{{"courses_id", coursesId}, {"lesson_id", lessonId}}
	update := bson.D{
		{"$push", bson.D{
			{"all_comments", data},
		}},
	}
	insertResult, err := userCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println("insert a comment fail")
		log.Fatal("insert a comment fail,", err)
		return errmsg.ERROR
	}
	fmt.Println("Inserted a single comment: ", insertResult.UpsertedID.(string))
	return errmsg.SUCCESS
}

//删除评论
func DeleteComment(coursesId string, lessonId string, commentId string) int {
	filter := bson.D{{"courses_id", coursesId}, {"lesson_id", lessonId}}
	update := bson.D{
		{"$pull", bson.M{"all_comments": bson.D{{"comment_id", commentId}}}},
	}
	insertResult, err := userCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println("delete a comment fail")
		log.Fatal("delete a comment fail,", err)
		return errmsg.ERROR
	}
	fmt.Println("deleted a single comment: ", insertResult.UpsertedID.(string))
	return errmsg.SUCCESS
}

func CreateCourseComment(coursesId string, lessonId string) int {
	data := Comments{
		bson.NewObjectId().Hex(),
		coursesId,
		lessonId,
		nil,
	}
	insertResult, err := userCollection.InsertOne(context.TODO(), data)
	if err != nil {
		fmt.Println("create a comment fail")
		log.Fatal("create a comment fail,", err)
		return errmsg.ERROR
	}
	fmt.Println("create a single document: ", insertResult.InsertedID.(string))
	return errmsg.SUCCESS
}
func commentInit() {
	if db != nil {
		if commentCollection == nil {
			commentCollection = db.Collection("comment")
		} else {
			fmt.Println("comment collection has inited")
			log.Fatal("comment collection has inited")
		}
	} else {
		fmt.Println("mongodb comment collection init error, db has not inited")
		log.Fatal("mongodb comment collection init error, db has not inited")
	}
}
