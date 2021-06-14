package processors

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"practice-news-board-web/database"
	"practice-news-board-web/messages"
	"practice-news-board-web/models"
	"strconv"
)

func GetBoardList() []messages.Board {
	db := database.GetDB()

	var board []messages.Board
	db.Table(models.Board{}.Table()).Find(&board)

	return board
}

func GetBoardDetail(id string) messages.Board {
	db := database.GetDB()

	board := messages.Board{}
	db.Table(models.Board{}.Table()).Find(&board, id)

	return board
}

func GetLatestNews() []messages.News {
	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/newstories.json?print=pretty")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var datas []uint
	json.Unmarshal([]byte(data), &datas)

	var news []messages.News

	for i := 0; i < 10; i++ {
		var n messages.News
		var url = "https://hacker-news.firebaseio.com/v0/item/" + strconv.FormatUint(uint64(datas[i]), 10) + ".json?print=pretty"
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		r, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		json.Unmarshal([]byte(r), &n)

		news = append(news, n)
	}

	return news
}
