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

var homepageCollection *mongo.Collection = nil

const GET_COURSE_NUMBER = 30

type CourseInfo struct {
	ID          string `bson:"_id"`
	CourseId    string `bson:"course_id"`
	CourseName  string `bson:"course_name"`
	CourseImage string `bson:"course_image"`
	Createtime  string `bson:"createtime"  json:"createtime"`
}

func Timing() int {
	findoptions := options.Find()
	findoptions.SetLimit(GET_COURSE_NUMBER)
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
	err = homepageCollection.Drop(context.TODO())
	if err != nil {
		log.Fatal("drop CourseInfo fail,", err)
		return errmsg.ERROR
	}
	insertResult, err := homepageCollection.InsertMany(context.TODO(), courseInfos)
	if err != nil {
		fmt.Println("insert a CourseInfo fail")
		log.Fatal("insert a CourseInfo fail,", err)
		return errmsg.ERROR
	}
	fmt.Println("create a single document: ", insertResult.InsertedIDs)
	return errmsg.SUCCESS
}

func Retrieve() ([]CourseInfo, int) {
	var courseInfos []CourseInfo = make([]CourseInfo, 0)
	cursor, err := homepageCollection.Find(context.TODO(), bson.M{})
	defer cursor.Close(context.TODO())
	if err != nil {
		fmt.Println("get a CourseInfo fail")
		log.Fatal("get a CourseInfo fail,", err)
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

func HomepageInit() {
	if db != nil {
		if homepageCollection == nil {
			homepageCollection = db.Collection("homepage")
		} else {
			fmt.Println("homepage collection has inited")
			log.Fatal("homepage collection has inited")
		}
	} else {
		fmt.Println("mongodb homepage collection init error, db has not inited")
		log.Fatal("mongodb homepage collection init error, db has not inited")
	}
}
