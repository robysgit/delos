package model

type ApiHistorySummary struct {
	Method     string
	Url        string
	Total      int
	UniqueUser int
}
