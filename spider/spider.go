package spider

import (
	"fmt"
	"net/http"
	"time"
)

func Daily() {
	Toutiaoio()

	//时间格式化格式 "2006-01-02 15:04:05.999999999 -0700 MST"

	//计时器
	ticker := time.NewTicker(time.Hour * 1)
	func() {
		for _ = range ticker.C {

			if time.Now().Format("15") == "10" { //上午十点执行

			}
		}
	}()
}

//http://toutiao.io/
func Toutiaoio() {
	fmt.Println(getdate(-1, "2006-01-02"))
	http.Get("")
}

//action
func getdate(action time.Duration, style string) string {

	baseTime := time.Date(1980, 1, 6, 0, 0, 0, 0, time.UTC)
	date := baseTime.Add(1722*7*24*time.Hour + action*24*time.Hour + 66355*time.Second)
	return date.Format(style)
}
