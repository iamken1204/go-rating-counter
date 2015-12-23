package rating_counter

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
)

/**
 * @param engine  string (google|yahoo)
 *                       The target search engine
 * @param keyword string Keyword of searching
 * @param url     string Url that want to assert
 */
func Crawl(targetID int, engine, keyword, targetUrl string) {
	// checkParam()
	var queryFormat string
	var querySelector string
	var startCount int
	var startRating int
	if engine == "google" {
		queryFormat = "http://www.google.com.tw/search?q=%s&start=%d"
		querySelector = "h3.r a"
		startCount = -10
		startRating = 1
	} else {
		queryFormat = "https://tw.search.yahoo.com/search?p=%s&b=%d"
		querySelector = "div.aTitle p a"
		startCount = -9
		startRating = 1
	}
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang_test")
	checkError(err)
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	target := InitTarget(keyword, targetUrl, startCount, startRating)
	doQuery := true
	doFind := true
	var realLink string

	for doQuery {
		target.Page++
		target.Start += 10
		queryUrl := fmt.Sprintf(queryFormat, target.Key, target.Start)
		fmt.Println(queryUrl)

		response, err := http.Get(queryUrl)
		checkError(err)
		defer response.Body.Close()

		doc, err := goquery.NewDocumentFromReader(io.Reader(response.Body))
		checkError(err)

		doc.Find(querySelector).Each(func(i int, s *goquery.Selection) {
			str, exists := s.Attr("href")
			if exists {
				target.Rating++
				u, err := url.Parse(str)
				checkError(err)
				if engine == "google" {
					m, _ := url.ParseQuery(u.RawQuery)
					realLink = m["q"][0]
				} else {
					realLink = u.String()
				}
				if realLink == target.Url {
					fmt.Printf("key: %s, url: %s, rating: %d, page: %d\n",
						target.Key, target.Url, target.Rating-1, target.Page)
					dbInsert, err := db.Prepare("insert into targets_logs (target_id, search_engine, rating, page, recorded_at) values (?, ?, ?, ?, ?)")
					checkError(err)
					defer dbInsert.Close()
					timestamp := time.Now().Local().Format("2006-01-02 15:04:05")
					_, err = dbInsert.Exec(targetID, engine, target.Rating-1, target.Page, timestamp)
					checkError(err)
					doFind = false
				}
			}
		})

		if target.Page >= 3 || !doFind {
			doQuery = false
		}
	}

	fmt.Println("No match query.")
}

func checkParam() {
	access := true
	if len(os.Args) < 4 {
		access = false
	}
	if !access {
		fmt.Println("Missing params")
		os.Exit(0)
	}
}
