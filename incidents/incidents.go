package incidents

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type core struct {
	AuthenticationResultCode string       `json:"authenticationResultCode"`
	BrandLogoURI             string       `json:"brandLogoUri"`
	Copyright                string       `json:"copyright"`
	ResourceSets             resourceSets `json:"resourceSets"`
	StatusCode               int          `json:"statusCode"`
	StatusDescription        string       `json:"statusDescription"`
	TraceID                  string       `json:"traceId"`
}
type point struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}
type toPoint struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}
type resources []struct {
	Type         string  `json:"__type"`
	Point        point   `json:"point"`
	Description  string  `json:"description"`
	End          string  `json:"end"`
	IncidentID   int64   `json:"incidentId"`
	LastModified string  `json:"lastModified"`
	RoadClosed   bool    `json:"roadClosed"`
	Severity     int     `json:"severity"`
	Source       int     `json:"source"`
	Start        string  `json:"start"`
	ToPoint      toPoint `json:"toPoint"`
	IncidentType int     `json:"type"`
	Verified     bool    `json:"verified"`
}
type resourceSets []struct {
	EstimatedTotal int       `json:"estimatedTotal"`
	Resources      resources `json:"resources"`
}

func GetIncident(location, key string) (r core) {

	//Set Variables
	var i core
	baseUrl := "http://dev.virtualearth.net/REST/v1/Traffic/Incidents/"
	url := baseUrl + location + "?key=" + key
	//Build New Request
	req, err := http.NewRequest("Get", url, nil)
	if err != nil {
		log.Fatalln("error with GET Response", err)
	}
	//Get Response from Request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln("error with GET Response", err)
	}
	//fmt.Println(url)
	defer res.Body.Close()
	//Unmarshal Json into data Struct
	body, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(body, &i)
	if err != nil {
		log.Fatalln("error unmarshalling", err)
	}

	return i
}
