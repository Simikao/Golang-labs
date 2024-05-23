package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"simikao/server-in-go/internal/datatype"
	"strconv"
	"syscall"

	"github.com/gin-gonic/gin"
)

// var (
// 	posts   = make(map[int]datatype.Post)
// 	nextID  = 1
// 	postsMu sync.Mutex
// )

// const portNum = ":3000"

var (
	posts = 100
	db    = []datatype.SharkAttack{}
)

//	func Home(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprint(w, "Homepage")
//	}
func MainPage(c *gin.Context, posts *[]datatype.SharkAttack) {
	mainSize := 20
	var mainPage []datatype.SharkAttack
	for range mainSize {
		mainPage = append(mainPage, (*posts)[rand.Intn(len((*posts)))])
	}

	c.JSON(http.StatusOK, mainPage)
}

func SinglePost(c *gin.Context, posts *[]datatype.SharkAttack) {
	vars := c.Param("id")
	id, err := strconv.Atoi(vars)
	if err != nil {
		fmt.Println("error: ", err)
		c.JSON(404, struct{ problem string }{problem: "not found"})
		return
	}

	if i, found := findPostIndexByID(id); found {
		c.JSON(http.StatusOK, (*posts)[i])
		return
	} else {
		c.JSON(http.StatusNotFound, datatype.Response{
			Success: false,
			Data:    "Post not found",
		})
		return
	}
}

func bodyDecoder(c *gin.Context, body interface{}) error {
	err := json.NewDecoder(c.Request.Body).Decode(body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, datatype.Response{
			Success: false,
			Data:    "Badly formated JSON",
		})
		return err
	}
	return nil
}

func AddPost(c *gin.Context, posts *[]datatype.SharkAttack, iter *datatype.Iterator) {
	var newPost datatype.SharkAttack
	err := bodyDecoder(c, &newPost)
	if err != nil {
		return
	}

	iter.Next()
	newPost.ID = iter.Current()

	(*posts) = append((*posts), newPost)

	c.JSON(http.StatusCreated, datatype.Response{
		Success: true,
		Data:    "Post added",
	})
}

func remove(posts *[]datatype.SharkAttack, i int) {
	(*posts)[i] = (*posts)[len((*posts))-1]
	(*posts) = (*posts)[:len((*posts))-1]
}

func findPostIndexByID(id int) (int, bool) {
	for index, post := range db {
		if post.ID == id {
			return index, true
		}
	}
	return -1, false
}

func RemovePost(c *gin.Context, posts *[]datatype.SharkAttack) {
	vars := c.Param("id")
	id, err := strconv.Atoi(vars)
	if err != nil {
		fmt.Println("error: ", err)
		c.JSON(404, struct{ problem string }{problem: "not found"})
		return
	}

	if i, found := findPostIndexByID(id); found {
		remove(posts, i)
		c.JSON(http.StatusOK, datatype.Response{
			Success: true,
			Data:    "Post removed",
		})
		return
	} else {
		c.JSON(http.StatusNotFound, datatype.Response{
			Success: false,
			Data:    "Post not found",
		})
		return
	}
}

func DownloadPost(c *gin.Context, posts *[]datatype.SharkAttack) {
	vars := c.Param("id")
	id, err := strconv.Atoi(vars)
	if err != nil {
		fmt.Println("error: ", err)
		c.JSON(404, struct{ problem string }{problem: "not found"})
		return
	}

	postJson, err := json.Marshal((*posts)[id])
	if err != nil {
		c.JSON(http.StatusTeapot, datatype.Response{
			Success: false,
			Data:    "Somehow something went wrong",
		})
	}
	string := "./tmp/post" + strconv.Itoa(id) + ".json"
	err = os.WriteFile(string, postJson, 0644)
	if err != nil {
		c.JSON(http.StatusTeapot, datatype.Response{
			Success: false,
			Data:    err.Error(),
		})
	}
}

func main() {
	bitties, err := os.ReadFile("./global-shark-attack.json")
	if err != nil {
		panic(err)
	}

	var result []datatype.SharkAttack
	err = json.Unmarshal(bitties, &result)
	if err != nil {
		panic(err)
	}

	iter := datatype.NewIterator()
	for ; iter.Current() < posts; iter.Next() {
		entry := result[rand.Intn(len(result))]
		entry.ID = iter.Current()

		db = append(db, entry)
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) { MainPage(c, &db) })
	r.GET("/post/:id", func(c *gin.Context) { SinglePost(c, &db) })
	r.POST("/post/add", func(c *gin.Context) { AddPost(c, &db, &iter) })
	r.DELETE("/post/remove/:id", func(c *gin.Context) { RemovePost(c, &db) })
	r.PUT("/post/:id", func(c *gin.Context) { DownloadPost(c, &db) })

	go r.Run()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
