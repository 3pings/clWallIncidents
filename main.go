package main

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/3pings/clWallIncidents/config"
	"github.com/3pings/clWallIncidents/incidents"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"log/syslog"
)

type incidentData struct {
	Description string
	Severity    int
	IncidentID  int64
	End         string
	Coordinates string
}

func main() {

	//loc := "41.346847,2.143373,41.45781,2.175545"
	loc := "32.553,-116.936,32.982,-117.254"
	authKey := "AuqW9LUO744wxmr6ik0PiJYAGbtKqmelkReiUT5_jhP4Dgh4cwDH-m9vE5wLY2z3"

	logwriter, e := syslog.New(syslog.LOG_NOTICE, "incident")
	if e == nil {
		log.SetOutput(logwriter)
	}

	q := incidents.GetIncident(loc, authKey)
	//Parse fields for specific information
	id := incidentData{}

	id.Description = q.ResourceSets[0].Resources[0].Description
	id.Severity = q.ResourceSets[0].Resources[0].Severity
	id.IncidentID = q.ResourceSets[0].Resources[0].IncidentID
	id.Coordinates = fmt.Sprintf("%f", q.ResourceSets[0].Resources[0].Point.Coordinates[0]) + "," + fmt.Sprintf("%f",q.ResourceSets[0].Resources[0].Point.Coordinates[1])
	//id.End = strings.SplitAfter(q.ResourceSets[0].Resources[0].End,")")[:1]
	id.End = strings.Split(strings.Split(q.ResourceSets[0].Resources[0].End, "(")[1], ")")[0][0:10]

	d := insertData(config.DB, id)
	if d != nil {
		fmt.Println(d)

	}

	fmt.Println(id.Coordinates)
	fmt.Println(id.IncidentID)
	fmt.Println(id.Severity)
	fmt.Println(id.End)
}

//noinspection ALL
//noinspection ALL
func insertData(s *sql.DB, i incidentData) error {

	//Insert Data into Database

	_, err := s.Exec("INSERT incidents(Severity, Incident_ID, Coordinates, Description, End) VALUES(?,?,?,?,?)", i.Severity, i.IncidentID, i.Coordinates, i.Description, i.End)
	log.Print("Successfully created DB record for incident info")

	return err

}
