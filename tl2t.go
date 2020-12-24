package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/mmcdole/gofeed"
)

func main() {
	logsDir := "./logs/"
	defaultCategoriesRegexp := ("Foo|Bar")
	defaultFeedURL := "https://localhost/invalid"
	flagFeedURL := flag.String("feedurl", defaultFeedURL, "Default feed URL")
	flagCategoriesRegexp := flag.String("categoriesregexp", defaultCategoriesRegexp, "Default categories regexp")

	flag.Parse()
	feedURL := *flagFeedURL
	stringCategoriesRegexp := "^(" + *flagCategoriesRegexp + ")$"
	categoriesRegexp := regexp.MustCompile(stringCategoriesRegexp)

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(feedURL)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for _, item := range feed.Items {
		if categoriesRegexp.MatchString(item.Categories[0]) {
			outputFile := logsDir + strings.ReplaceAll(strings.ToLower(item.Categories[0]), " ", "_")
			f, err := os.OpenFile(outputFile,
				os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}
			if _, err := f.WriteString(item.Title + "\n"); err != nil {
				log.Println(err)
			}
			f.Close()
		}
	}
}
