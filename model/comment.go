package model

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"

	"University-Information-Website/utils/errmsg"
)

type SingleComment struct {
	ID         string `bson:"_id"`
	UserId     string `bson:"user_id"`
	Username   string `bson:"username"`
	Content    string `bson:"content"`
	Createtime string `bson:"createtime" json:"createtime"`
}

type Comments struct {
	ID          string          `bson:"_id"`
	CoursesID   string          `bson:"courses_id"`
	LessonID    string          `bson:"lesson_id"`
	AllComments []SingleComment `bson:"all_comments"`
}

var commentCollection *mongo.Collection = nil

//添加对应每个小节的Comments表
func CreateComments(coursesId string, lessonId string) int {
	data := Comments{
		bson.NewObjectId().Hex(),
		coursesId,
		lessonId,
		make([]SingleComment, 0),
	}
	insertResult, err := commentCollection.InsertOne(context.TODO(), data)
	if err != nil {
		fmt.Println("create a comments fail")
		log.Fatal("create a comments fail,", err)
		return errmsg.ERROR
	}
	fmt.Println("create a comments document: ", insertResult.InsertedID.(string))
	return errmsg.SUCCESS
}

// 添加评论
func InsertComment(data *SingleComment, coursesId string, lessonId string) int {
	data.ID = bson.NewObjectId().Hex()
	filter := bson.M{"courses_id": coursesId, "lesson_id": lessonId}
	update := bson.M{"$push": bson.M{"all_comments": data}}
	insertResult, err := commentCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println("insert a comment fail")
		log.Fatal("insert a comment fail,", err)
		return errmsg.ERROR
	}
	fmt.Printf("Matched %v documents and insert %v documents.\n", insertResult.MatchedCount, insertResult.ModifiedCount)
	return errmsg.SUCCESS
}

//删除一条评论
func DeleteComment(coursesId string, lessonId string, commentId string) int {
	filter := bson.M{"courses_id": coursesId, "lesson_id": lessonId}
	update := bson.M{"$pull": bson.M{"all_comments": bson.M{"_id": commentId}}}
	insertResult, err := commentCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println("delete a comment fail")
		log.Fatal("delete a comment fail,", err)
		return errmsg.ERROR
	}
	fmt.Printf("Matched %v documents and insert %v documents.\n", insertResult.MatchedCount, insertResult.ModifiedCount)
	return errmsg.SUCCESS
}

//删除某节课的所有评论
func DeleteLessonAllComment(coursesId string, lessonId string) int {
	filter := bson.M{"courses_id": coursesId, "lesson_id": lessonId}
	deleteResult, err := commentCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Println("delete lesson comments fail")
		log.Fatal("delete lesson comments fail,", err)
		return errmsg.ERROR
	}
	fmt.Printf("Matched and delete %v documents.\n", deleteResult.DeletedCount)
	return errmsg.SUCCESS
}

func DeleteCourseAllComment(coursesId string) int {
	filter := bson.M{"courses_id": coursesId}
	deleteResult, err := commentCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Println("delete courses comments fail")
		log.Fatal("delete courses comments fail,", err)
		return errmsg.ERROR
	}
	fmt.Printf("Matched and delete %v documents.\n", deleteResult.DeletedCount)
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
