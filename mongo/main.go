package main

import (
	mgo "gopkg.in/mgo.v2"
)

// User is User
type User struct {
	Name string
}

func main() {

	sess, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{"localhost:27017"},
		Username: "mongo_user",
		Password: "mongo_user_pass",
	})
	if err != nil {
		panic(err)
	}
	defer sess.Close()

	c := sess.DB("testing").C("user")
	usr := User{"testing_name_1"}
	if err := c.Insert(usr); err != nil {
		panic(err)
	}
}
