package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// makeRange is makeRange
func makeRange(min, max int) []int {
	r := make([]int, max-min+1)
	a := min
	for i := range r {
		r[i] = a
		a++
	}
	return r
}

// West is West
type West struct {
	OriginalProjectName       string
	FinalProjectName          string
	Contract                  string
	AWE                       string
	OwnerLocation             string
	GeneralContractorLocation string
	ProjectManager            string
}

// Westeros is Westeros
type Westeros struct {
	URL                       string
	OriginalProjectName       string
	FinalProjectName          string
	Contract                  string
	AWE                       string
	OwnerLocation             string
	GeneralContractorLocation string
	ProjectManager            string
}

func main() {

	// mongo conn
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
	c := sess.DB("fake").C("west")

	// read xls
	// xls, err := excelize.OpenFile("info.xlsx")
	// if err != nil {
	// 	panic(err)
	// }
	// cellHeader := []string{"A", "B", "C", "D", "E", "K", "L"}
	// cellRowNum := makeRange(2, 54)
	// for _, num := range cellRowNum {
	// 	westDatum := West{}
	// 	for _, head := range cellHeader {
	// 		cell := fmt.Sprintf("%s%s", head, strconv.Itoa(num))
	// 		cellValue := xls.GetCellValue("Sheet2", cell)
	// 		switch head {
	// 		case "A":
	// 			westDatum.OriginalProjectName = strings.TrimSpace(cellValue)
	// 		case "B":
	// 			westDatum.FinalProjectName = strings.TrimSpace(cellValue)
	// 		case "C":
	// 			westDatum.Contract = strings.TrimSpace(cellValue)
	// 		case "D":
	// 			westDatum.AWE = strings.TrimSpace(cellValue)
	// 		case "E":
	// 			westDatum.OwnerLocation = strings.TrimSpace(cellValue)
	// 		case "K":
	// 			westDatum.GeneralContractorLocation = strings.TrimSpace(cellValue)
	// 		case "L":
	// 			westDatum.ProjectManager = strings.TrimSpace(cellValue)
	// 		}
	// 	}

	// 	// insert to mongo
	// 	if err := c.Insert(westDatum); err != nil {
	// 		panic(err)
	// 	} else {
	// 		fmt.Println("Inserted Name: ", westDatum.OriginalProjectName)
	// 	}
	// }

	// read txt
	proj, err := ioutil.ReadFile("projects.txt")
	if err != nil {
		panic(err)
	}
	projects := strings.Split(string(proj), "\n")
	i := sess.DB("fake").C("westeros")
	for _, project := range projects {
		// split id and name
		tmps := strings.Split(string(project), "|")
		id := strings.TrimSpace(tmps[0])
		name := strings.TrimSpace(tmps[1])
		// find in mongo
		query := bson.M{"$or": []bson.M{bson.M{"originalprojectname": name}, bson.M{"finalprojectname": name}}}
		result := Westeros{}
		fmt.Println("Searching:", name)
		c.Find(query).One(&result)
		if result.OriginalProjectName != "" {
			// add url
			result.URL = fmt.Sprintf("%s%s%s", "https://awestelectric.wpengine.com/x/#/content/", id, "/inspector")
			// insert to mongo
			if err := i.Insert(result); err != nil {
				panic(err)
			} else {
				fmt.Println("Inserted Name: ", result.OwnerLocation)
			}
		} else {
			fmt.Println("Not Found:", strings.TrimSpace(project))
		}
	}
}
