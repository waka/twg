package views

import (
	"fmt"
	"strings"

	runewidth "github.com/mattn/go-runewidth"
	termbox "github.com/nsf/termbox-go"
	"github.com/waka/twg/twitter"
)

type TweetsView struct {
	cursorY      int
	scroll       int
	scrollOffset int
}

func NewTweetsView() *TweetsView {
	return &TweetsView{cursorY: 0, scroll: 0, scrollOffset: 0}
}

func (view *TweetsView) Draw(tweets []*twitter.Tweet) {
	_, height := getWindowSize()
	y := (view.scroll - view.scrollOffset) * -1

	for i, tweet := range tweets {
		countLine := view.tweetLinesCount(tweet)
		if y > height {
			break
		} else if y+countLine < 0 {
			y += countLine
			continue
		}

		y = view.drawUserLine(tweet, i, y)
		y = view.drawContents(tweet, i, y)
		y = view.drawMetaInfo(tweet, i, y)
	}
}

func (view *TweetsView) drawUserLine(tweet *twitter.Tweet, index int, y int) int {
	cursorColor := view.cursorColor(index)

	// cursor
	x := 0
	drawText(" ", x, y, NewColors(ForeGround(cursorColor), BackGround(cursorColor)))

	// screen name
	x = 1
	nameColor := GetRandomColor(tweet.User.Id)
	drawText("@"+tweet.User.ScreenName, x, y, NewColors(ForeGround(nameColor), BackGround(ColorBackground)))

	// name
	//x += runewidth.StringWidth("@"+tweet.User.ScreenName) + 1
	//drawText(tweet.User.Name, x, y, NewColors(ForeGround(ColorLowlight), BackGround(ColorBackground)))

	// favorited and retweeted
	x += runewidth.StringWidth(tweet.User.Name) + 1
	if tweet.Favorited {
		drawText("★", x, y, NewColors(ForeGround(ColorYellow), BackGround(ColorBackground)))
		x += runewidth.StringWidth("★") + 1
	}
	if tweet.Retweeted {
		drawText("RT", x, y, NewColors(ForeGround(ColorGreen), BackGround(ColorBackground)))
		x += runewidth.StringWidth("RT") + 1
	}

	return y + 1
}

func (view *TweetsView) drawContents(tweet *twitter.Tweet, index int, y int) int {
	width, _ := getWindowSize()
	cursorColor := view.cursorColor(index)

	lines := strings.Split(runewidth.Wrap(tweet.Text, width-2), "\n")
	for i, text := range lines {
		// cursor
		drawText(" ", 0, y+i, NewColors(ForeGround(cursorColor), BackGround(cursorColor)))
		drawText(text, 1, y+i, NewColors(ForeGround(ColorWhite), BackGround(ColorBackground)))
	}
	y += len(lines)

	return y
}

func (view *TweetsView) drawMetaInfo(tweet *twitter.Tweet, index int, y int) int {
	cursorColor := view.cursorColor(index)
	strTime := tweet.PastTimeFromNow()

	// cursor
	x := 0
	drawText(" ", x, y, NewColors(ForeGround(cursorColor), BackGround(cursorColor)))

	x = 1
	drawText(strTime, x, y, NewColors(ForeGround(ColorGray1), BackGround(ColorBackground)))

	if tweet.FavoriteCount > 0 {
		strFav := fmt.Sprintf("Fav: %d", tweet.FavoriteCount)
		drawText(" "+strFav, 18, y, NewColors(ForeGround(ColorYellow), BackGround(ColorBackground)))
	}
	if tweet.RetweetCount > 0 {
		strRT := fmt.Sprintf("RT: %d", tweet.RetweetCount)
		drawText(" "+strRT, 30, y, NewColors(ForeGround(ColorGreen), BackGround(ColorBackground)))
	}

	return y + 1
}

func (view *TweetsView) TopCursor() {
	view.cursorY = 0
	view.scroll = 0
}

func (view *TweetsView) UpCursor(tweets []*twitter.Tweet, nextCursorY int) {
	if view.cursorY == 0 {
		return
	}
	view.cursorY = nextCursorY

	sum := 0
	tweets = tweets[:view.cursorY]
	for _, tweet := range tweets {
		sum += view.tweetLinesCount(tweet)
	}
	sub := sum - view.scroll
	if sub <= 0 {
		view.scroll += sub
		if view.scroll < 0 {
			view.scroll = 0
		}
	}
}

func (view *TweetsView) DownCursor(tweets []*twitter.Tweet, nextCursorY int) {
	if view.cursorY == len(tweets)-1 {
		return
	}
	view.cursorY = nextCursorY

	_, height := getWindowSize()
	sum := 0
	tweets = tweets[:view.cursorY+1]
	for _, tweet := range tweets {
		sum += view.tweetLinesCount(tweet)
	}
	sub := (sum - (view.scroll - view.scrollOffset)) - (height - 2)
	if sub >= 0 {
		view.scroll += sub
	}
}

func (view *TweetsView) GetCursorPosition() int {
	return view.cursorY
}

func (view *TweetsView) cursorColor(index int) termbox.Attribute {
	cursorColor := ColorBackground
	if index == view.cursorY {
		cursorColor = ColorHighlight
	}
	return cursorColor
}

func (view *TweetsView) tweetLinesCount(tweet *twitter.Tweet) int {
	width, _ := getWindowSize()
	if tweet.RetweetedStatus != nil {
		tweet = tweet.RetweetedStatus
	}
	text := tweet.Text
	lines := strings.Split(runewidth.Wrap(text, width-2), "\n")
	lineCount := 1 + len(lines) + 1

	return lineCount
}
