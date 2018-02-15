package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//SearchResults structure with query and all Result data
type SearchResults struct {
	ready   bool
	Query   string
	Results []Result
}

//Result structure which contains parsed results
type Result struct {
	Name, Description, URL string
}

//UnmarshalJSON set data to strcuture
func (sr *SearchResults) UnmarshalJSON(bs []byte) error {
	array := []interface{}{}
	if err := json.Unmarshal(bs, &array); err != nil {
		return err
	}
	sr.Query = array[0].(string)
	for i := range array[1].([]interface{}) {
		sr.Results = append(sr.Results, Result{
			array[1].([]interface{})[i].(string),
			array[2].([]interface{})[i].(string),
			array[3].([]interface{})[i].(string),
		})
	}
	return nil
}

func wikipediaAPI(request string) (answer []string) {

	s := make([]string, 3)

	//Sending request
	if response, err := http.Get(request); err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()

		//Reading answer
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		//Unmarshal answer and set it to SearchResults struct
		sr := &SearchResults{}
		if err = json.Unmarshal([]byte(contents), sr); err != nil {
			s[0] = "Something going wrong, try to change your question"
		}

		//Check if struct is not empty
		if !sr.ready {
			s[0] = "Something going wrong, try to change your question"
		}

		//Loop through Results struct and assigning data to s slice
		for i := range sr.Results {
			s[i] = sr.Results[i].URL
		}
	}
	return s
}