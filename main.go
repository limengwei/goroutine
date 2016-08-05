package main

import (
	"github.com/limengwei/gtc/gopher"
	"github.com/limengwei/gtc/spider"
)

func main() {

	spider.Daily()

	go gopher.RssRefresh()
	gopher.StartServer()
}
