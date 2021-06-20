package model

import (
	"University-Information-Website/middleware"
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/scrypt"
	"gopkg.in/mgo.v2/bson"

	"University-Information-Website/utils/errmsg"
)

var userCollection *mongo.Collection = nil

type User struct {
	ID       string `bson:"_id"`
	Username string `bson:"username"`
	Password string `bson:"password"`
	Role     int    `bson:"role"`
}

// 查询用户是否存在
func CheckUser(data *User) int {
	if data.Username == "" {
		return errmsg.ERROR_UERNAME_EMPTY
	}
	if data.Password == "" {
		return errmsg.ERROR_PASSWORD_EMPTY
	}
	var result User
	filter := bson.M{"username": data.Username}
	err := userCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err.Error() != "mongo: no documents in result" {
			fmt.Println("check user error")
			log.Fatal("check user error,", err)
			return errmsg.ERROR
		} else {
			return errmsg.SUCCESS
		}
	}
	if result.ID != "" {
		return errmsg.ERROR_UERNAME_USED
	}
	return errmsg.SUCCESS
}

// 添加用户
func InsertUser(data *User) int {
	data.ID = bson.NewObjectId().Hex()
	newData := data
	// 密码加密
	newData.Password = ScrypyPw(data.Password)
	insertResult, err := userCollection.InsertOne(context.TODO(), newData)
	if err != nil {
		fmt.Println("insert a user fail")
		log.Fatal("insert a user fail,", err)
		return errmsg.ERROR
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID.(string))
	return errmsg.SUCCESS
}

// 查询用户列表
func GetUsers(pageSize int, pageNum int) ([]User, int) {
	return nil, errmsg.SUCCESS
}

// 删除用户
func DeleteUser(id string) int {
	deleteResult, err := userCollection.DeleteOne(context.TODO(), bson.M{"_id": bson.ObjectIdHex(id)})
	if err != nil {
		log.Fatal("delete a user fail", err)
		return errmsg.ERROR
	}

	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
	return errmsg.SUCCESS
}

// 修改用户
func UpdateUser(id string, data *User) int {
	filter := bson.M{"_id": bson.ObjectIdHex(id)}
	updateResult, err := userCollection.UpdateOne(context.TODO(), filter, data)
	if err != nil {
		fmt.Println("update a user fail")
		log.Fatal("update  auser fail,", err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return errmsg.SUCCESS
}

// 密码加密
func ScrypyPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 66, 11, 222, 11}
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}

// 登录验证
func CheckLogin(data *User) int {
	password := data.Password
	var result User
	filter := bson.M{"username": data.Username}
	err := userCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println("check login error")
		log.Fatal("check login error,", err)
	}
	if result.ID == "" {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScrypyPw(password) != result.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if result.Role != 1 || result.Role != 0 {
		return errmsg.ERROR_USER_NOT_RIGHT
	}
	return errmsg.SUCCESS
}

func userInit() {
	if db != nil {
		if userCollection == nil {
			userCollection = db.Collection("user")
		} else {
			fmt.Println("user collection has inited")
			log.Fatal("user collection has inited")
		}
	} else {
		fmt.Println("mongodb user collection init error, db has not inited")
		log.Fatal("mongodb user collection init error, db has not inited")
	}
}

// tokenHeader := c.Request.Header.Get("Authorization")
func CheckAuthority(tokenHeader string) int {
	checkToken := strings.Split(tokenHeader, " ")
	key, _ := middleware.CheckToken(checkToken[1])
	return GetAuthority(key.Id)
}

func GetAuthority(userid string) int {
	var result User
	filter := bson.M{"_id": userid}
	err := userCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println("GetAuthority fail")
		log.Fatal("GetAuthority fail,", err)
	}
	return result.Role
}
