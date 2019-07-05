package main

import (
	"sort"

  "github.com/waka/twg/twitter"
)

type TweetsStore struct {
	timelineTweets twitter.Tweets
	mentionsTweets twitter.Tweets
	listTweets     map[string]twitter.Tweets
}

var tweetStore *TweetsStore = newTweetsStore()

func newTweetsStore() *TweetsStore {
	return &TweetsStore{listTweets: map[string]twitter.Tweets{}}
}

func GetTweetsStore() *TweetsStore {
	return tweetStore
}

func (store *TweetsStore) GetTimelineTweets() twitter.Tweets {
	return store.timelineTweets
}

func (store *TweetsStore) SetTimelineTweets(tweets twitter.Tweets) {
	results := filterTweets(
		uniqTweets(sortTweets(mergeTweets(store.timelineTweets, tweets))))
	store.timelineTweets = results
}

func (store *TweetsStore) GetMentionsTweets() twitter.Tweets {
	return store.mentionsTweets
}

func (store *TweetsStore) SetMentionsTweets(tweets twitter.Tweets) {
	results := filterTweets(
		uniqTweets(sortTweets(mergeTweets(store.mentionsTweets, tweets))))
	store.mentionsTweets = results
}

func (store *TweetsStore) GetListTweets(listName string) twitter.Tweets {
	return store.listTweets[listName]
}

func (store *TweetsStore) SetListTweets(listName string, tweets twitter.Tweets) {
	list := store.listTweets[listName]
	results := filterTweets(
		uniqTweets(sortTweets(mergeTweets(list, tweets))))
	store.listTweets[listName] = results
}

// merge with excluding same id
func mergeTweets(tweets twitter.Tweets, newer twitter.Tweets) twitter.Tweets {
	return append(tweets, newer...)
}

// find by tweet id
func uniqTweets(tweets twitter.Tweets) twitter.Tweets {
	results := make([]*twitter.Tweet, 0, len(tweets))
	encountered := map[int64]bool{}
	for _, tweet := range tweets {
		if !encountered[tweet.Id] {
			encountered[tweet.Id] = true
			results = append(results, tweet)
		}
	}
	return results
}

func sortTweets(tweets twitter.Tweets) twitter.Tweets {
	sort.Sort(sort.Reverse(tweets))
	return tweets
}

// filter by limitation number (=100)
func filterTweets(tweets twitter.Tweets) twitter.Tweets {
	if len(tweets) > 100 {
		return tweets[0:100]
	}
	return tweets
}
