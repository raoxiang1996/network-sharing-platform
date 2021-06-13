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

type Comment struct {
	ID         bson.ObjectId   `bson:"_id"`
	Userid     bson.ObjectId   `bson:"userid"`
	Username   string          `bson:"username"`
	Content    string          `bson:"content"`
	Createtime common.JsonTime `bson:"createtime" json:"createtime"`
}

var commentCollection *mongo.Collection = nil

// 添加评论
func InsertComment(data *Comment, courseId string, userId string) int {
	data.ID = bson.NewObjectId()
	newData := data
	insertResult, err := commentCollection.InsertOne(context.TODO(), newData)
	if err != nil {
		fmt.Println("insert a comment fail")
		log.Fatal("insert a comment fail,", err)
		return errmsg.ERROR
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID.(string))
	return errmsg.SUCCESS
}

//删除评论
func DeleteComment(courseId string, userId string, commentId string) int {
	deleteResult, err := userCollection.DeleteOne(context.TODO(), bson.M{"_id": bson.ObjectIdHex(courseId)})
	if err != nil {
		log.Fatal("delete a comment fail", err)
		return errmsg.ERROR
	}

	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
	return errmsg.SUCCESS
}

func commentInit() {
	if db != nil {
		if commentCollection == nil {
			userCollection = db.Collection("course")
		} else {
			fmt.Println("comment collection has inited")
			log.Fatal("comment collection has inited")
		}
	} else {
		fmt.Println("mongodb comment collection init error, db has not inited")
		log.Fatal("mongodb comment collection init error, db has not inited")
	}
}
