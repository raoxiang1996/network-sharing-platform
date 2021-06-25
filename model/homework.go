package model

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
	HomeworkId string   `bson:"homework_id"`
	Answers    []string `bson:"answers"`
}
