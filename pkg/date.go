package pkg

import "time"

func GetRangeDateByDateQuery(dateQuery string) (*time.Time, *time.Time) {
	currentTime := time.Now()
	switch dateQuery {
	case "week":
		startOfWeek := currentTime.AddDate(0, 0, -int(currentTime.Weekday())).Truncate(24 * time.Hour)
		endOfWeek := startOfWeek.AddDate(0, 0, 7).Add(-time.Second)
		return &startOfWeek, &endOfWeek
	case "month":
		startOfMonth := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, time.UTC)
		endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Second)
		return &startOfMonth, &endOfMonth
	}

	return &currentTime, &currentTime
}
