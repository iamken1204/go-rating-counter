package rating_counter

import (
	// "fmt"
	// "time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Targets struct {
	ID      int    `json:"id"`
	Keyword string `json:"keyword"`
	Url     string `json:"url"`
	Status  string `json:"status"`
	Log     []Logs `json:"log"`
}

type Logs struct {
	ID         int    `json:"id"`
	MTargetID  int    `json:"target_id"`
	Rating     int    `json:"rating"`
	Page       int    `json:"page"`
	RecordedAt string `json:"recorded_at"`
}

func (l Logs) TableName() string {
	return "targets_logs"
}

type Target struct {
	Key    string
	Url    string
	Rating int
	Page   int
	Start  int
}

func InitTarget(keyword, url string, startCount, startRating int) Target {
	t := Target{}
	t.Key = keyword
	t.Url = url
	t.Rating = startRating
	t.Page = 0
	t.Start = startCount
	return t
}

func GetTargets() []Targets {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang_test")
	checkError(err)
	defer db.Close()
	targets := []Targets{}
	db.Order("id desc").Find(&targets)
	return targets
}

func GetTargetsByUrl(url string) []Logs {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang_test")
	checkError(err)
	defer db.Close()
	targets := []Targets{}
	db.Where("url = ?", url).Find(&targets)
	id := targets[0].ID
	logs := []Logs{}
	db.Where("target_id = ?", id).Find(&logs)
	return logs
}

func Create(data func(string) string) bool {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang_test")
	checkError(err)
	defer db.Close()
	var log []Logs
	// timestamp := time.Now().Local().Format("2006-01-02 15:04:05")
	target := Targets{0, data("url"), data("keyword"), "1", log}
	db.NewRecord(target)
	db.Create(&target)
	return true
}

func NewEmptyLogs() []Logs {
	ts := []Logs{}
	return ts
}
