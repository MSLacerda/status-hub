package services

type Service struct {
	StatusService string
	Name          string
	Url           string
}

var allServices = builtServices()

func create(statusService, name, url string) Service {
	var newService = Service{StatusService: statusService, Name: name, Url: url}

	return newService
}

func builtServices() *[]Service {
	var arr []Service

	service_github := create("StatusPage", "GitHub", "https://www.githubstatus.com/api/v2/status.json")
	service_netlify := create("StatusPage", "Netlify", "https://www.netlifystatus.com/api/v2/status.json")
	service_discord := create("StatusPage", "Discord", "https://status.discordapp.com/api/v2/status.json")
	service_heroku := create("Heroku", "Heroku", "https://status.heroku.com/api/v3/current-status")

	arr = append(arr, service_github)
	arr = append(arr, service_heroku)
	arr = append(arr, service_netlify)
	arr = append(arr, service_discord)

	return &arr
}

func GetServices() *[]Service {
	if len(*allServices) > 0 {
		return allServices
	} else {
		allServices = builtServices()

		return allServices
	}
}
