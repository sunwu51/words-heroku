package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
	"github.com/gin-gonic/gin"
)

var TOKEN string = os.Getenv("TOKEN")
var SECRET string = os.Getenv("SECRET")
var jsonServer string = "http://localhost:5556/words/"
var zone, _ = time.LoadLocation("Asia/Shanghai")
var ch = make(chan string)

type Item struct {
	Id   string   `json:"id"`
	List []string `json:"list"`
}

func gitPull() error {
	cmd := exec.Command("/bin/sh", "-c", "cd /root/code/go-tutorial/words-heroku/words-db && git pull origin master && git config --global user.email \"sunwu51@126.com\" && git config --global user.name \"frank\"")
	err := cmd.Run()
	if err != nil {
		log.Println("git pull error", err)
		return err
	}
	log.Println("git pull success")
	return nil
}

func gitPush(words []string) error {
	if len(words) > 0 {
		wordStr := strings.Join(words, ",")
		cmd := exec.Command("/bin/sh", "-c",
			fmt.Sprintf("cd./words-db &&  git add words.json && git commit -m \"add word %s\" && git push https://%s@github.com/sunwu51/words-db.git", wordStr, TOKEN),
		)

		err := cmd.Run()

		if err != nil {
			log.Println("git push error", err)
			return err
		}
	}
	return nil
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func getAllWords(c *gin.Context) {
	res, _ := http.Get(jsonServer)
	bytes, _ := ioutil.ReadAll(res.Body)
	m := make([]Item, 100)
	json.Unmarshal(bytes, &m)
	c.JSON(http.StatusOK, m)
}

func getWordsById(c *gin.Context) {
	id := c.Param("id")
	log.Println("id", id)
	res, _ := http.Get(fmt.Sprintf("%s%s", jsonServer, id))
	bytes, _ := ioutil.ReadAll(res.Body)
	m := Item{}
	json.Unmarshal(bytes, &m)
	c.JSON(http.StatusOK, m)
}

func addWord(c *gin.Context) {
	word := c.PostForm("word")
	secret := c.PostForm("secret")

	log.Println(secret, SECRET, secret != SECRET)
	if secret != SECRET {
		c.JSON(http.StatusOK, gin.H{
			"status": "error",
		})
		return
	}

	id := getMonday()
	res, _ := http.Get(fmt.Sprintf("%s%s", jsonServer, id))
	bytes, _ := ioutil.ReadAll(res.Body)
	m := Item{}
	json.Unmarshal(bytes, &m)

	// 这一周还没有单词，那么需要post创建
	if m.Id == "" && len(m.List) == 0 {
		m.Id = id
		m.List = []string{word}
		b, _ := json.Marshal(m)
		http.Post(jsonServer, "application/json", strings.NewReader(string(b)))
		c.JSON(http.StatusOK, m)
	} else { // 反之 如果已经有了，那么需要put更新
		m.List = append(m.List, word)
		b, _ := json.Marshal(m)
		req, _ := http.NewRequest("PUT", jsonServer+id, strings.NewReader(string(b)))
		req.Header.Set("Content-Type", "application/json")
		http.DefaultClient.Do(req)
		c.JSON(http.StatusOK, m)
	}
	ch <- word
}

func getMonday() string {
	t := time.Now().In(zone)
	offset := int(time.Monday - t.Weekday())
	if offset > 0 {
		offset = -6
	}

	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, zone).
		AddDate(0, 0, offset).Format("2006-01-02")
}

func cronJob() {
	ticker := time.NewTicker(1 * time.Minute)
	list := []string{}
	for {
		select {
		case x := <-ch:
			list = append(list, x)
		case <-ticker.C:
			gitPush(list)
			list = []string{}
		}
	}
}

func main() {
	// gitPull()

	go cronJob()

	r := gin.Default()
	r.GET("/health", health)
	r.GET("/", getAllWords)
	r.GET("/:id", getWordsById)
	r.POST("/add", addWord)

	r.Run()
}
