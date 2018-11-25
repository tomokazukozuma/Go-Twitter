package main

import (
	"net/http"
	"fmt"
	"flag"
	"os"
	"strings"
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"github.com/olekukonko/tablewriter"
)

type Tweet struct {
	Name string
	ScreenName string
	Text string
	Time string
	FavoriteCount string
	RetweetCount string
	ReplyCount string
}

func main() {

	isLatest := flag.Bool("latest", false, "retreive latest tweet")
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		log.Fatal(errors.New("Search word is not setting!!"))
	}
	word := args[0]
	if word == "" {
		log.Fatal(errors.New("Search word is empty!!"))
	}
	url := fmt.Sprintf("https://twitter.com/search?lang=ja&q=%s" + "", word)
	if *isLatest {
		url += "&f=tweets"
	}

	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	doc, _ := goquery.NewDocumentFromReader(resp.Body)

	var tweets []Tweet
	doc.Find("li.js-stream-item").Each(func(_ int, s *goquery.Selection) {
		name := strings.TrimSpace(s.Find("span.FullNameGroup").Text())
		screenName := strings.TrimSpace(s.Find("span.username").Text())
		favoriteCount, exists := s.Find("span.ProfileTweet-action--favorite").Find("span.ProfileTweet-actionCount").Attr("data-tweet-stat-count")
		if !exists {
			favoriteCount = "0"
		}
		retweetCount, exists := s.Find("span.ProfileTweet-action--retweet").Find("span.ProfileTweet-actionCount").Attr("data-tweet-stat-count")
		if !exists {
			retweetCount= "0"
		}
		replyCount, exists := s.Find("span.ProfileTweet-action--reply").Find("span.ProfileTweet-actionCount").Attr("data-tweet-stat-count")
		if !exists {
			replyCount = "0"
		}

		tweets = append(tweets, Tweet{
			Name: name,
			ScreenName: screenName,
			Text: s.Find("p.tweet-text").Text(),
			Time: s.Find("span._timestamp").Text(),
			FavoriteCount: favoriteCount,
			RetweetCount: retweetCount,
			ReplyCount: replyCount,
		})
	})

	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)

	table.SetHeader([]string{"Time", "Screen Name", "Tweet", "Fav", "Ret", "Rep"})
	for _, t := range tweets {
		table.Append([]string {t.Time, strings.Trim(t.Name, "\n") + strings.Trim(t.ScreenName, "\n"), strings.TrimRight(t.Text, "\n"), t.FavoriteCount, t.RetweetCount, t.ReplyCount})
	}
	table.Render()
}