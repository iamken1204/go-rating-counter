package rating_counter

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
)

/**
 * @param os.Args[1] string (google|yahoo)
 *                          The target search engine
 * @param os.Args[2] string Keyword of searching
 * @param os.Args[3] string Url that want to assert
 */
func Crawl() {
	checkParam()

	var queryFormat string
	var querySelector string
	var startCount int
	var startRating int
	if os.Args[1] == "google" {
		queryFormat = "https://www.google.com.tw/search?q=%s&start=%d"
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
	target := InitTarget(os.Args[2], os.Args[3], startCount, startRating)
	doQuery := true

	var realLink string
	for doQuery {
		target.Page++
		target.Start += 10

		response, err := http.Get(fmt.Sprintf(queryFormat, target.Key, target.Start))
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
				if os.Args[1] == "google" {
					m, _ := url.ParseQuery(u.RawQuery)
					realLink = m["q"][0]
				} else {
					realLink = u.String()
				}
				if realLink == target.Url {
					fmt.Printf("key: %s, url: %s, rating: %d, page: %d\n",
						target.Key, target.Url, target.Rating-1, target.Page)
					dbInsert, err := db.Prepare("insert into test_search_query (keyword, url, rating, page) values (?, ?, ?, ?)")
					checkError(err)
					defer dbInsert.Close()
					_, err = dbInsert.Exec(target.Key, target.Url, target.Rating-1, target.Page)
					checkError(err)
					os.Exit(1)
				}
			}
		})

		if target.Page >= 3 {
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
