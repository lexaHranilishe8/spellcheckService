package spellcheck_service

import "time"

type Stats struct {
	TotalProducts   int
	TotalErrors     int
	LastCheckedTime time.Time
}

var errors []ProductError
var stats Stats

func SaveErrors(newErrors []ProductError) {
	errors = append(errors, newErrors...)
	stats.TotalErrors += len(newErrors)
}

func UpdateStats(total int) {
	stats.TotalProducts += total
	stats.LastCheckedTime = time.Now()
}

func GetStatistics() Stats {
	return stats
}

func GetErrors() []ProductError {
	return errors
}
