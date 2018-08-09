package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	waybackURL      = "http://archive.org/wayback/available?url=beachorganicsskincare.com/"
	productsKeyword = []string{
		"Natural Bodywash",
		"Organic Soap",
		"Body Butter",
		"Body Powder",
		"Body Spray",
		"Conditioner",
		"Dog Shampoo",
		"Fur Fresh",
		"Lip Balm",
		"Body Lotion",
		"Makeup Remover",
		"Scrub",
		"Shampoo",
	}
	productsReplacement = []string{
		"Bodywash",
		"Soap",
		"BodyButter",
		"BodyPowder",
		"BodySpray",
		"Conditioner",
		"DogShampoo",
		"FurFresh",
		"LipBalm",
		"BodyLotion",
		"MakeupRemover",
		"Scrub",
		"Shampoo",
	}
	wbLogFile, _    = os.Create("./logs/log.log")
	wbOutputFile, _ = os.Create("./result.log")
	wbl             = log.New(wbLogFile, "", 0)
	wbo             = log.New(wbOutputFile, "", 0)
)

// WBResponse is WBResponse
type WBResponse struct {
	URL     string
	ArchivedSnapshots Closest struct {
		Status    string
		Available bool
		URL       string
		Timestamp string
	}
}

// WaybackPull is WaybackPull
func WaybackPull(wbReq string, chWBResp chan WBResponse, chDone chan bool) {
	r, err := http.Get(wbReq)
	defer func() {
		chDone <- true
	}()
	if err != nil {
		wbl.Fatal(err)
	}

	var wbr WBResponse
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&wbr); err != nil {
		wbl.Println("json error", err)
		if serr, ok := err.(*json.SyntaxError); ok {
			wbl.Println("Occurred at offset:", serr.Offset)
		}
	}
	chWBResp <- wbr
}

// FormWaybackRequestURI is FormWaybackRequestURI
func FormWaybackRequestURI(b *bytes.Buffer, uris ...string) {
	for _, s := range uris {
		b.WriteString(s)
	}
	b.WriteString("~")
}

func main() {
	// read csv
	bList, err := ioutil.ReadFile("product.csv")
	if err != nil {
		wbl.Fatal(err)
	}
	r := csv.NewReader(strings.NewReader(string(bList)))
	products, err := r.ReadAll()
	if err != nil {
		wbl.Fatal(err)
	}
	// clean csv
	checks := map[string]bool{}
	var pNames []string
	for _, p := range products {
		if _, found := checks[p[1]]; !found {
			checks[p[1]] = true
			pNames = append(pNames, p[1])
		}
	}
	// generate variation
	var vS1, vS2, vS3, vS4 string
	var wbReqBuff bytes.Buffer
	for _, hay := range pNames {
		if strings.Contains(hay, "by PSW") {
			vS1 = strings.Title(strings.ToLower(strings.Replace(hay, " ", "-", -1)))
			FormWaybackRequestURI(&wbReqBuff, waybackURL, vS1, ".html")
			break
		}
		found := false
		for i := 0; i < len(productsKeyword); i++ {
			if strings.Contains(hay, productsKeyword[i]) {
				found = true
				vS1 = strings.Join([]string{
					strings.Title(strings.ToLower(strings.TrimSpace(strings.Replace(strings.Replace(hay, productsKeyword[i], "", 1), " ", "", -1)))),
					strings.Replace(productsReplacement[i], " ", "", -1),
				}, "-")
				FormWaybackRequestURI(&wbReqBuff, waybackURL, vS1, ".html")
				vS2 = strings.Join([]string{
					strings.Title(strings.Replace(strings.TrimSpace(strings.Replace(hay, productsKeyword[i], "", 1)), " ", "-", -1)),
					strings.ToLower(strings.Replace(productsReplacement[i], " ", "", -1)),
				}, "-")
				FormWaybackRequestURI(&wbReqBuff, waybackURL, vS2, ".html")
				vS3 = strings.Join([]string{
					strings.Title(strings.Replace(strings.TrimSpace(strings.Replace(hay, productsKeyword[i], "", 1)), " ", "-", -1)),
					strings.Replace(productsReplacement[i], " ", "", -1),
				}, "-")
				FormWaybackRequestURI(&wbReqBuff, waybackURL, vS3, ".html")
				vS4 = strings.Join([]string{
					strings.Title(strings.ToLower(strings.TrimSpace(strings.Replace(strings.Replace(hay, productsKeyword[i], "", 1), " ", "", -1)))),
					strings.ToLower(strings.Replace(productsReplacement[i], " ", "", -1)),
				}, "-")
				FormWaybackRequestURI(&wbReqBuff, waybackURL, vS4, ".html")
				break
			}
			// last
			if i == len(productsKeyword)-1 && !found {
				vS1 = strings.Title(strings.ToLower(strings.Replace(hay, " ", "-", -1)))
				FormWaybackRequestURI(&wbReqBuff, waybackURL, vS1, ".html")
				break
			}
		}
	}
	// concurrent start:
	var WBResponses []WBResponse
	chWBResp := make(chan WBResponse)
	chDone := make(chan bool)
	// check each product variation way
	uris := strings.Split(wbReqBuff.String(), "~")
	for _, uri := range uris {
		if uri != "" {
			go WaybackPull(uri, chWBResp, chDone)
		}
	}
	for c := 0; c < len(uris)-1; {
		select {
		case wbResp := <-chWBResp:
			wbo.Println(wbResp.Closest)
			WBResponses = append(WBResponses, wbResp)
		case <-chDone:
			c++
		}
	}
	// if exist, pull ingredients
	// if exist, save to json
	// concurrent end:
	close(chWBResp)

}
