package twitter

import (
	"fmt"
	"time"
)

type User struct {
	Id         int64
	ScreenName string `json:"screen_name"`
	Name       string
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
	RetweetCount    int    `json:"retweet_count"`
	Favorited       bool
	FavoriteCount   int `json:"favorite_count"`
}

func (tweet *Tweet) PastTimeFromNow() string {
	createdAtTime, _ := time.Parse(time.RubyDate, tweet.CreatedAt)
	now := time.Now()
	sub := now.Sub(createdAtTime)

	var strTime string
	if sub <= time.Second*30 {
		strTime = "now"
	} else if sub <= time.Minute*5 {
		strTime = "A few minutes ago"
	} else if sub <= time.Hour*2 {
		m := sub / time.Minute
		strTime = fmt.Sprintf("%d minutes ago", m)
	} else if sub <= time.Hour*36 {
		h := sub / time.Hour
		strTime = fmt.Sprintf("%d hours ago", h)
	} else if sub <= time.Hour*24*14 {
		d := (sub / (time.Hour * 24))
		di := (sub % (time.Hour * 24))
		// Round up if time of now is later than six o'clock
		if now.Hour() >= 6 && di > 0 {
			d++
		}
		strTime = fmt.Sprintf("%d/%d/%d %02d:%02d (%d days ago)",
			createdAtTime.Year(), createdAtTime.Month(),
			createdAtTime.Day(),
			createdAtTime.Hour(), createdAtTime.Minute(), d)
	} else {
		strTime = fmt.Sprintf("%d/%d/%d %02d:%02d",
			createdAtTime.Year(), createdAtTime.Month(),
			createdAtTime.Day(),
			createdAtTime.Hour(), createdAtTime.Minute())
	}

	return strTime
}

type Tweets []*Tweet

func (t Tweets) Len() int {
	return len(t)
}

func (t Tweets) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t Tweets) Less(i, j int) bool {
	return t[i].Id < t[j].Id
}
