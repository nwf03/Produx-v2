package queries

type Post struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	User        User   `json:"user"`
}
type User struct {
	Name string `json:"name"`
	Pfp  string `json:"pfp"`
}

// func GetPosts(field string, productId string, lastId string) ([]db.Bug, int64, error) {
// 	var posts []db.IDGetter 
// 	var count int64
// 	field = strings.ToLower(field)
// 	var queryError error
// 	switch field {
// 	case "suggestions":
// 		suggestions, err := db.DB.GetBugs(productId, lastId, true)
// 		queryError = err	
// 		for _, v := range suggestions {
// 			posts = append(posts, &v)
// 		}
// 	case "bugs":
// 		bugs, err := db.DB.GetBugs(productId, lastId, true)
// 		return bugs, count, err
// 	case "changelogs":
// 		changelogs, err := db.DB.GetChangelogs(productId, lastId, true)
// 		queryError = err
// 		for _, v := range changelogs {
// 			posts = append(posts, &v)
// 		}
// 	case "announcements":
// 		announcements, err := db.DB.GetAnnouncements(productId, lastId, true)
// 		queryError = err
// 		for _, v := range announcements {
// 			posts = append(posts, &v)
// 		}
// 	}
// 	return make([]db.Bug, 12), count, queryError

// }
