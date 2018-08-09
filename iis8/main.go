package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

// User is User
type User struct {
	Name  string
	Email string
}

// printMongo is printMongo
func printMongo(context *gin.Context) {
	mongoConn := &mgo.DialInfo{
		Addrs:    []string{"192.168.99.100:27017"},
		Username: "mongo_user",
		Password: "mongo_user_pass",
	}

	sess, err := mgo.DialWithInfo(mongoConn)
	if err != nil {
		panic(err)
	}
	defer sess.Close()
	c := sess.DB("testing").C("user")
	users := make([]User, 2)
	c.Find(nil).All(&users)
	var message bytes.Buffer
	for _, user := range users {
		message.WriteString(fmt.Sprintf("Name: %s, Email: %s", user.Name, user.Email))
	}
	context.JSON(200, gin.H{
		"message": message.String(),
	})
}

// printSimple is printSimple
func printSimple(context *gin.Context)  {
	context.JSON(200, gin.H{
		"message": "hello simple",
	})	
}

func main() {
	r := gin.Default()
	r.GET("/user", printMongo)
	r.GET("/simple", printSimple)
	r.Run(":"+os.Getenv("HTTP_PLATFORM_PORT"))
}
