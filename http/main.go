package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gopkg.in/mgo.v2"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

// User is User
//   {
//     "id": 1,
//     "name": "Leanne Graham",
//     "username": "Bret",
//     "email": "Sincere@april.biz",
//     "address": {
//       "street": "Kulas Light",
//       "suite": "Apt. 556",
//       "city": "Gwenborough",
//       "zipcode": "92998-3874",
//       "geo": {
//         "lat": "-37.3159",
//         "lng": "81.1496"
//       }
//     },
//     "phone": "1-770-736-8031 x56442",
//     "website": "hildegard.org",
//     "company": {
//       "name": "Romaguera-Crona",
//       "catchPhrase": "Multi-layered client-server neural-net",
//       "bs": "harness real-time e-markets"
//     }
//   },
type User struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
	Body     string  `json:"body"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
	Company  Company `json:"company"`
}

// Address is Address
type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     Geo    `json:"geo"`
}

// Geo is Geo
type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

// Company is Company
type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

func main() {
	mongoConn := &mgo.DialInfo{
		Addrs:    []string{"localhost:27017"},
		Username: "mongo_user",
		Password: "mongo_user_pass",
	}
	sess, err := mgo.DialWithInfo(mongoConn)
	if err != nil {
		panic(err)
	}
	defer sess.Close()

	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	users := []User{}

	json.NewDecoder(resp.Body).Decode(&users)
	// insert to mongo
	c := sess.DB("fake").C("users")
	for _, uInst := range users {
		if err := c.Insert(uInst); err != nil {
			panic(err)
		} else {
			fmt.Println("Inserted Name: ", uInst.Name)
		}
	}
}
