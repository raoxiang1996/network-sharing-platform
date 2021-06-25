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

type Lesson struct {
	ID         string `bson:"_id"`
	CoursesId  string `bson:"courses_id"`
	Path       string `bson:"path"`
	UploadTime string `bson:"upload_time"  json:"createtime"`
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

// 搜索一门课程
func GetCourse(CourseId string) (Courses, int) {
	filter := bson.M{"_id": CourseId}
	var course Courses
	err := coursesCollection.FindOne(context.TODO(), filter).Decode(&course)
	if err != nil {
		fmt.Println("get a course fail")
		log.Fatal("get a course fail,", err)
		return Courses{}, errmsg.ERROR
	}
	return course, errmsg.SUCCESS
}

// 按照规则查找课程 五个参数，index，返回的数量，是否排序，顺序还是逆序，排序的字段
func GetAllCourse(index int, limit int, sortOption ...interface{}) ([]Courses, int) {
	findoptions := options.Find()
	if len(sortOption) == 3 {
		isSort := sortOption[0].(bool)
		if isSort == true {
			option := sortOption[2].(string)
			if sortOption[1].(int) == -1 {
				findoptions.SetSort(bson.M{option: -1})
			} else if sortOption[1].(int) == 1 {
				findoptions.SetSort(bson.M{option: 1})
			} else {
				fmt.Println("parameter error：second option should be 1 or -1")
				log.Fatal("parameter error：second option should be 1 or -1")
				return nil, errmsg.ERROR
			}
		}
	} else if len(sortOption) > 0 && len(sortOption) != 3 {
		fmt.Println("parameter error：number of option should be 3")
		log.Fatal("parameter error：number of option should be 3")
		return nil, errmsg.ERROR
	}

	if limit > 0 {
		findoptions.SetLimit(int64(limit))
		findoptions.SetSkip(int64(limit) * int64(index))
	} else if limit == 0 {
		return nil, errmsg.SUCCESS
	}

	cursor, err := coursesCollection.Find(context.TODO(), bson.M{}, findoptions)
	if err != nil {
		fmt.Println("get courses fail")
		log.Fatal("get courses fail,", err)
		return nil, errmsg.ERROR
	}

	if err := cursor.Err(); err != nil {
		log.Fatal("cursor is error", err)
		return nil, errmsg.ERROR
	}
	defer cursor.Close(context.Background())
	var courses []Courses = make([]Courses, 0)
	for cursor.Next(context.Background()) {
		var tmpCourse Courses
		if err = cursor.Decode(&tmpCourse); err != nil {
			log.Fatal("decode course fail,", err)
			return nil, errmsg.ERROR
		}
		courses = append(courses, tmpCourse)
	}
	return courses, errmsg.SUCCESS
}

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
	filter := bson.M{"_id": coursesId}
	update := bson.M{"$push": bson.M{"subject": data}}
	insertResult, err := coursesCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println("insert a lesson fail")
		log.Fatal("insert a lesson fail,", err)
		return errmsg.ERROR
	}
	fmt.Printf("Matched %v documents and insert %v documents.\n", insertResult.MatchedCount, insertResult.ModifiedCount)
	return errmsg.SUCCESS
}

//更新课程信息
func UpdateCourse(data *Courses, coursesId string) int {
	filter := bson.M{"_id": coursesId}
	upDateData := bson.M{
		"course_name":  data.CourseName,
		"introduction": data.Introduction,
		"images":       data.Images}
	update := bson.M{"$set": upDateData}
	updateResult, err := coursesCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println("update a course fail")
		log.Fatal("update a course fail,", err)
		return errmsg.ERROR
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return errmsg.SUCCESS
}

//更新小节信息
func UpdateLesson(data *Lesson, lessonId string) int {
	filter := bson.M{"_id": data.CoursesId}
	upDateData := bson.M{
		"subject.$[elem].path": data.Path}
	update := bson.M{"$set": upDateData}
	opaf := []interface{}{bson.M{"elem._id": lessonId}}
	af := options.FindOneAndUpdateOptions{ArrayFilters: &options.ArrayFilters{nil, opaf}}
	updateResult := coursesCollection.FindOneAndUpdate(context.TODO(), filter, update, &af)
	if updateResult.Err() != nil {
		log.Println("Find error: ", updateResult.Err())
	}
	return errmsg.SUCCESS
}

//删除一节课程
func DeleteLesson(coursesId string, lessonId string) int {
	//删除这节的所有评论
	code := DeleteLessonAllComment(coursesId, lessonId)
	if code != errmsg.SUCCESS {
		fmt.Println("delete a course fail")
		log.Fatal("delete a course fail,", errmsg.GetErrMsg(code))
		return errmsg.ERROR
	}

	filter := bson.M{"_id": coursesId}
	update := bson.M{"$pull": bson.M{"subject": bson.M{"_id": lessonId, "courses_id": coursesId}}}
	updateResult, err := coursesCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		fmt.Println("delete a course fail")
		log.Fatal("delete a course fail,", err)
		return errmsg.ERROR
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return errmsg.SUCCESS
}

//删除课程
func DeleteCourses(coursesId string) int {
	//删除课程下所有评论
	code := DeleteCourseAllComment(coursesId)
	if code != errmsg.SUCCESS {
		log.Fatal("delete courses fail", errmsg.GetErrMsg(code))
		return errmsg.ERROR
	}

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
