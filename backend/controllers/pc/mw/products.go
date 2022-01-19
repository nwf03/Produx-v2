package mw

func ValidFields(field string) bool {
	switch field {
	case
		"Suggestions",
		"Bugs",
		"Changelogs",
		"Announcements":
		return true
	}
	return false

}

func GetPageCount(count int64) int64 {
	var numPages int64
	if count%10 == 0 {
		numPages = count / 10
	} else {
		numPages = int64(count/10) + 1
	}
	return numPages
}
