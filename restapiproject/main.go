package main

import(
	"fmt"
	"log"
	"net/http"
	 "encoding/json"
)

type article struct{
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}
type articles []article

func allArticles(w http.ResponseWriter, r *http.Request){
	articles := articles{
		article{Title:"Test title", Desc:"Test Description", Content: "Hello World"},
	}
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}


func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "home page of main")
}

func handleRequests(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8080",nil))
}
func main(){
	handleRequests()
}