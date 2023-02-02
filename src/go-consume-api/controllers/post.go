// jsonplaceholder: https://jsonplaceholder.typicode.com/guide/
package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	//"fmt"
	"html/template"
	"log"
	"net/http"
)

// Please refer to guide of jsonplaceholder
// Automatically get values matching keys by requesting HTTP method of Get
type Post struct {
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int64  `json:"userId"`
}

var BASE_URL = "https://jsonplaceholder.typicode.com"

func Index(w http.ResponseWriter, r *http.Request) {

	var posts []Post

	// Get Rest API of https://jsonplaceholder.typicode.com
	response, err := http.Get(BASE_URL + "/posts")
	if err != nil {
		log.Print(err)
	}

	defer response.Body.Close()

	// Error handling by decoding data format of JSON Requested Get method
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&posts); err != nil {
		log.Print(err)
	}

	data := map[string]interface{}{
		"posts": posts,
	}

	temp, _ := template.ParseFiles("views/index.html")
	temp.Execute(w, data)

}

func Create(w http.ResponseWriter, r *http.Request) {

	var post Post
	var data map[string]interface{}

	id := r.URL.Query().Get("id")
	if id != "" {
		// Error handling by requesting request HTTP method of Get
		res, err := http.Get(BASE_URL + "/posts/" + id)
		if err != nil {
			log.Print(err)
		}
		defer res.Body.Close()

		// Error handling by decoding data format of JSON by getting response of Get method
		decoder := json.NewDecoder(res.Body)
		if err := decoder.Decode(&post); err != nil {
			log.Print(err)
		}

		data = map[string]interface{}{
			"post": post,
		}
	}

	temp, _ := template.ParseFiles("views/create.html")
	temp.Execute(w, data)

}

func Store(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	id := r.Form.Get("post_id")

	// Get id and parse as type of int
	idInt, _ := strconv.ParseInt(id, 10, 64)
	newPost := Post{
		Id:     idInt,
		Title:  r.Form.Get("post_title"),
		Body:   r.Form.Get("post_body"),
		UserId: 1,
	}

	// mashalling
	jsonValue, _ := json.Marshal(newPost)
	buff := bytes.NewBuffer(jsonValue) // convert mashalled data to byte

	var req *http.Request
	var err error

	if id != "" {
		// update request
		fmt.Println("Process update")
		req, err = http.NewRequest(http.MethodPut, BASE_URL+"/posts/"+id, buff)
	} else {
		// create request
		fmt.Println("Process create")
		req, err = http.NewRequest(http.MethodPost, BASE_URL+"/posts", buff)
	}

	if err != nil {
		log.Print(err)
	}

	// Set API Common Header
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	// Create HTTP client
	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		log.Print(err)
	}
	defer res.Body.Close()

	var postsResponse Post

	// Error handling by decoding data format of JSON Requested any method by the client
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&postsResponse); err != nil {
		log.Print(err)
	}

	fmt.Println(res.StatusCode)
	fmt.Println(res.Status)
	fmt.Println(postsResponse)

	// Redirect web page if request is success(code 201)
	if res.StatusCode == 201 || res.StatusCode == 200 {
		http.Redirect(w, r, "/posts", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	// Get http request of Delete by use id
	req, err := http.NewRequest(http.MethodDelete, BASE_URL+"/posts/"+id, nil)
	if err != nil {
		log.Print(err)
	}

	// Create HTTP client
	httpClient := &http.Client{}
	// Request Delete to Rest API of https://jsonplaceholder.typicode.com
	res, err := httpClient.Do(req)
	if err != nil {
		log.Print(err)
	}

	defer res.Body.Close()

	// Confirm HTTP status code by request
	fmt.Println(res.StatusCode)
	fmt.Println(res.Status)

	if res.StatusCode == 200 {
		http.Redirect(w, r, "/posts", http.StatusSeeOther)
	}
}
