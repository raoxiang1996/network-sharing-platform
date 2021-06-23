package model

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"

	"University-Information-Website/utils/errmsg"
)

var indexCollection *mongo.Collection = nil

type CourseInfo struct {
	ID          string `bson:"_id"`
	CourseId    string `bson:"course_id"`
	CourseName  string `bson:"course_name"`
	CourseImage string `bson:"course_image"`
	Createtime  string `bson:"createtime"  json:"createtime"`
}

func Timing() int {
	findoptions := options.Find()
	findoptions.SetLimit(30)
	findoptions.SetSort(bson.M{"createtime": -1})
	cursor, err := coursesCollection.Find(context.TODO(), bson.M{}, findoptions)
	if err != nil {
		fmt.Println("get courses fail")
		log.Fatal("get courses fail,", err)
		return errmsg.ERROR
	}

	if err := cursor.Err(); err != nil {
		log.Fatal("cursor is error", err)
		return errmsg.ERROR
	}
	defer cursor.Close(context.Background())
	var courseInfos []interface{} = make([]interface{}, 0)
	for cursor.Next(context.Background()) {
		var tmpCourse Courses
		var tmpCourseInfos CourseInfo
		if err = cursor.Decode(&tmpCourse); err != nil {
			log.Fatal("decode course fail,", err)
			return errmsg.ERROR
		}
		tmpCourseInfos.ID = bson.NewObjectId().Hex()
		tmpCourseInfos.CourseId = tmpCourse.ID
		tmpCourseInfos.CourseName = tmpCourse.CourseName
		tmpCourseInfos.CourseImage = tmpCourse.Images
		tmpCourseInfos.Createtime = tmpCourse.Createtime
		courseInfos = append(courseInfos, tmpCourseInfos)
	}
	err = indexCollection.Drop(context.TODO())
	if err != nil {
		log.Fatal("drop index fail,", err)
		return errmsg.ERROR
	}
	insertResult, err := indexCollection.InsertMany(context.TODO(), courseInfos)
	if err != nil {
		fmt.Println("insert a index fail")
		log.Fatal("insert a index fail,", err)
		return errmsg.ERROR
	}
	fmt.Println("create a single document: ", insertResult.InsertedIDs)
	return errmsg.SUCCESS
}

func Retrieve() ([]CourseInfo, int) {
	var courseInfos []CourseInfo = make([]CourseInfo, 0)
	cursor, err := indexCollection.Find(context.TODO(), bson.M{})
	defer cursor.Close(context.TODO())
	if err != nil {
		fmt.Println("create a course fail")
		log.Fatal("create a course fail,", err)
		return nil, errmsg.ERROR
	}
	if err := cursor.Err(); err != nil {
		log.Fatal("cursor is error", err)
	}
	for cursor.Next(context.TODO()) {
		var tmpCurseInfo CourseInfo
		err := cursor.Decode(&tmpCurseInfo)
		if err != nil {
			log.Fatal(err)
		}
		courseInfos = append(courseInfos, tmpCurseInfo)
	}
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", len(courseInfos))
	return courseInfos, errmsg.SUCCESS
}

func IndexInit() {
	if db != nil {
		if indexCollection == nil {
			indexCollection = db.Collection("index")
		} else {
			fmt.Println("course collection has inited")
			log.Fatal("course collection has inited")
		}
	} else {
		fmt.Println("mongodb course collection init error, db has not inited")
		log.Fatal("mongodb course collection init error, db has not inited")
	}
}
