package main

func fetchTasks(baseURL, availability string) []Issue {
	var availabilityParam string
	switch availability {
	case "Low":
		availabilityParam = "1"
	case "Medium":
		availabilityParam = "3"
	case "High":
		availabilityParam = "5"
	default:
		availabilityParam = ""
	}
	fullURL := baseURL + "?sort=estimate&limit=" + availabilityParam
	return getIssues(fullURL)
}
