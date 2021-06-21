package common

import (
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

type JsonTime time.Time

const (
	timeFormart = "2006-01-02 15:04:05"
)

//实现json序列化，将时间转换成字符串byte数组
func (t JsonTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}

//实现json反序列化，从传递的字符串中解析成时间对象
func (t *JsonTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = JsonTime(now)
	return
}

//mongodb是存储bson格式，因此需要实现序列化bsonvalue(这里不能实现MarshalBSON，MarshalBSON是处理Document的)，将时间转换成mongodb能识别的primitive.DateTime
func (t *JsonTime) MarshalBSONValue() (bsontype.Type, []byte, error) {
	targetTime := primitive.NewDateTimeFromTime(time.Time(*t))
	return bson.MarshalValue(targetTime)
}

//实现bson反序列化，从mongodb中读取数据转换成time.Time格式，这里用到了bsoncore中的方法读取数据转换成datetime然后再转换成time.Time
func (t *JsonTime) UnmarshalBSONValue(t2 bsontype.Type, data []byte) error {
	v, _, valid := bsoncore.ReadValue(data, t2)
	if valid == false {
		return errors.New(fmt.Sprintf("%s, %s, %s", "读取数据失败:", t2, data))
	}
	*t = JsonTime(v.Time())
	return nil
}

func GetNowTime() (string, error) {
	t := JsonTime(time.Now())
	now, err := t.MarshalJSON()
	if err != nil {
		return "", err
	}
	return string(now), nil
}
