package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	mgo "gopkg.in/mgo.v2"
)

// SeriesDatum is SeriesDatum
type SeriesDatum struct {
	Name  string `json:"name"`
	Color string `json:"color"`
	Data  []int  `json:"data"`
}

// PieDatum is PieDatum
type PieDatum struct {
	Label  string `json:"label"`
	Series []SeriesDatum
}

func main() {
	// mongodb://<dbuser>:<dbpassword>@ds225840.mlab.com:25840/jadipergi
	mongoConn := &mgo.DialInfo{
		Addrs:    []string{"jadipergi:jadi123!@ds225840.mlab.com:25840"},
		Username: "jadipergi",
		Password: "jadi123!",
	}
	sess, err := mgo.DialWithInfo(mongoConn)
	if err != nil {
		panic(err)
	}
	defer sess.Close()
	c := sess.DB("jadipergi").C("piechart")

	content, err := ioutil.ReadFile("pie.json")
	if err != nil {
		panic(err)
	}
	var pies []PieDatum
	if err := json.Unmarshal(content, &pies); err != nil {
		panic(err)
	}

	for _, pie := range pies {
		if err := c.Insert(pie); err != nil {
			panic(err)
		} else {
			fmt.Println("label: ", pie.Label)
		}
	}
}
