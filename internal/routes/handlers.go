package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/MSLacerda/status-hub/internal/services"
)

type Page struct {
	Id         string
	Name       string
	Url        string
	Time_zone  string
	Updated_at time.Time
}

type Status struct {
	Indicator   string
	Description string
}

type StatusProdDev struct {
	Production  string
	Development string
}

type StatusPage struct {
	Page   Page
	Status Status
}

type HerokuStatus struct {
	Status StatusProdDev
	Issues []string
}

type NormalizedStatus struct {
	ServiceName string
	Status      string
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func decodeHerokuStatus(data HerokuStatus, name string, err bool) NormalizedStatus {
	var normalized NormalizedStatus
	normalized.ServiceName = name

	if data.Status.Development != "green" && data.Status.Production != "green" || err {
		normalized.Status = "red"
	} else {
		normalized.Status = "green"
	}

	return normalized

}

func decodeStatusPage(data StatusPage, name string, err bool) NormalizedStatus {
	var normalized NormalizedStatus

	fmt.Println(data)

	normalized.ServiceName = name
	if data.Status.Description != "All Systems Operational" || err {
		normalized.Status = "red"
	} else {
		normalized.Status = "green"
	}

	return normalized
}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func HandlerRequest(service services.Service) NormalizedStatus {
	switch service.StatusService {
	case "StatusPage":
		status := StatusPage{}
		err := getJson(service.Url, &status)

		return decodeStatusPage(status, service.Name, err != nil)
	case "Heroku":
		status := HerokuStatus{}
		getJson(service.Url, &status)

		fmt.Println(status.Status.Development)

		return decodeHerokuStatus(status, service.Name, err != nil)
	default:
		panic("Unexpected service name.")
	}
}
