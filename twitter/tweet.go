package twitter

type User struct {
	Id         int64
	ScreenName string `json:"screen_name"`
	Protected  bool
}

type Tweet struct {
	Id              int64
	User            *User
	Source          string
	Text            string
	CreatedAt       string `json:"created_at"`
	Retweeted       bool
	RetweetedStatus *Tweet `json:"retweeted_status"`
}

type List struct {
	Slug        string `json:"slug"`
	FullName    string `json:"full_name"`
	Name        string `json:"name"`
	IdStr       string `json:"id_str"`
	MemberCount int    `json:"member_count"`
	Description string `json:"description"`
	User        *User
}

type SearchResult struct {
	Statuses []Tweet
}
