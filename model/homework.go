package model

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"

	"University-Information-Website/utils/errmsg"
)

type Question struct {
	Id      string   `bson:"_id"`
	Content string   `bson:"content"`
	Option  []string `bson:"option"`
}

type Homework struct {
	Id        string     `bson:"_id"`
	CourseId  string     `bson:"course_id"`
	Questions []Question `bson:"questions"`
}

type Answer struct {
	Id         string   `bson:"_id"`
	UserId     string   `bson:"user_id"`
	CourseId   string   `bson:"course_id"`
	HomeworkId string   `bson:"homework_id"`
	Answers    []string `bson:"answers"`
}

var homeworkCollection *mongo.Collection = nil
var answerCollection *mongo.Collection = nil

func GetHomework(courseId string) (Homework, int) {
	filter := bson.M{"course_id": courseId}
	var homework Homework
	err := homeworkCollection.FindOne(context.TODO(), filter).Decode(&homework)
	if err != nil {
		fmt.Println("get a homework fail")
		log.Fatal("get a homework fail,", err)
		return Homework{}, errmsg.ERROR
	}
	return homework, errmsg.SUCCESS
}

func GetAnswer(userId string, homeworkId string) (Answer, int) {
	filter := bson.M{"user_id": userId, "homework_id": homeworkId}
	var answer Answer
	err := answerCollection.FindOne(context.TODO(), filter).Decode(&answer)
	if err != nil {
		fmt.Println("get a answer fail")
		log.Fatal("get a answer fail,", err)
		return Answer{}, errmsg.ERROR
	}
	return answer, errmsg.SUCCESS
}

//添加一份作业
func InsertHomework(data *Homework) int {
	data.Id = bson.NewObjectId().Hex()
	for i := 0; i < len(data.Questions); i++ {
		data.Questions[i].Id = bson.NewObjectId().Hex()
	}
	insertResult, err := homeworkCollection.InsertOne(context.TODO(), data)
	if err != nil {
		fmt.Println("create a homework fail")
		log.Fatal("create a homework fail,", err)
		return errmsg.ERROR
	}
	fmt.Println("create a homework document: ", insertResult.InsertedID.(string))
	return errmsg.SUCCESS
}

// 添加正确答案
func InsertHomeworkAnswer(data *Answer) int {
	data.Id = bson.NewObjectId().Hex()
	homework, msg := GetHomework(data.CourseId)
	if msg != errmsg.SUCCESS {
		fmt.Println("insert homework answer:get a homework fail")
		log.Fatal("insert homework answer:get a homework  fail")
		return errmsg.ERROR
	}
	data.HomeworkId = homework.Id
	data.UserId = "0"
	insertResult, err := answerCollection.InsertOne(context.TODO(), data)
	if err != nil {
		fmt.Println("insert a homework answer fail")
		log.Fatal("insert a homework answer  fail,", err)
		return errmsg.ERROR
	}
	fmt.Println("create a homework answer document: ", insertResult.InsertedID.(string))
	return errmsg.SUCCESS
}

//删除课程下所有答案
func DeleteAllAnswers(courseId string) int {
	filter := bson.M{"course_id": courseId}
	deleteResult, err := answerCollection.DeleteMany(context.TODO(), filter)
	if err != nil {
		fmt.Println("delete all answers fail")
		log.Fatal("delete all answers fail,", err)
		return errmsg.ERROR
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
	return errmsg.SUCCESS
}

//删除作业和所有答案
func DeleteHomeWorksAndAllAnswer(courseId string) int {
	msg := DeleteAllAnswers(courseId)
	if msg != errmsg.SUCCESS {
		fmt.Println("delete homework and all answer:delete all answers fail")
		log.Fatal("delete homework and all answer:delete all answers fail")
		return errmsg.ERROR
	}
	filter := bson.M{"course_id": courseId}
	deleteResult, err := homeworkCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Println("delete all homeworks fail")
		log.Fatal("delete all homeworks fail,", err)
		return errmsg.ERROR
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
	return errmsg.SUCCESS
}

//删除一份答案
func DeleteAnswer(uerId string, homeworkId string) int {
	filter := bson.M{"user_id": uerId, "homework_id": homeworkId}
	deleteResult, err := answerCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Println("delete a answer fail")
		log.Fatal("delete a answer fail,", err)
		return errmsg.ERROR
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
	return errmsg.SUCCESS
}

//删除所有用户答案
func DeleteAllUserAnswers(courseId string) int {
	filter := bson.M{"course_id": courseId}
	deleteResult, err := answerCollection.DeleteMany(context.TODO(), filter)
	if err != nil {
		fmt.Println("delete a answer fail")
		log.Fatal("delete a answer fail,", err)
		return errmsg.ERROR
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
	return errmsg.SUCCESS
}

//更新作业
func UpdateHomework(data *Homework) int {
	filter := bson.M{"course_id": data.CourseId}
	updateResult, err := homeworkCollection.UpdateOne(context.TODO(), filter, data)
	if err != nil {
		fmt.Println("update a homework fail")
		log.Fatal("update a homework fail,", err)
		return errmsg.ERROR
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return errmsg.SUCCESS
}

//更新答案
func UpdateAnswer(data *Answer) int {
	filter := bson.M{"user_id": data.UserId, "homework_id": data.HomeworkId}
	updateResult, err := answerCollection.UpdateOne(context.TODO(), filter, data)
	if err != nil {
		fmt.Println("update a answer fail")
		log.Fatal("update a answer fail,", err)
		return errmsg.ERROR
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return errmsg.SUCCESS
}

func homeworkInit() {
	if db != nil {
		if homeworkCollection == nil {
			homeworkCollection = db.Collection("homework")
		} else {
			fmt.Println("homework collection has inited")
			log.Fatal("homework collection has inited")
		}
	} else {
		fmt.Println("mongodb homework collection init error, db has not inited")
		log.Fatal("mongodb homework collection init error, db has not inited")
	}
}

func answerInit() {
	if db != nil {
		if answerCollection == nil {
			answerCollection = db.Collection("answer")
		} else {
			fmt.Println("answer collection has inited")
			log.Fatal("answer collection has inited")
		}
	} else {
		fmt.Println("mongodb answer collection init error, db has not inited")
		log.Fatal("mongodb answer collection init error, db has not inited")
	}
}
