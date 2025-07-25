package posts

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"forum/backend/controllers"
	"forum/backend/models"
	"forum/utils"
	"net/http"
)

type post struct {
	Title   string `json:"title"`
	Content string `json:"contentInput"`
	Categorys []string `json:"categorys"`
}


func FetchCreatPosts(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		controllers.Response("Method Not Allowed...", 405, w)
		return
	}
	
	userID := r.Context().Value("userId").(int)
	var postt post
	err := json.NewDecoder(r.Body).Decode(&postt)
	if err != nil {
		controllers.Response("Invalid JSON...", 40., w)
		return
	}

	fmt.Println("---->" ,  postt)

	return 


	
	if utils.ContainsEmpty(postt.Title, postt.Content) {
		controllers.Response("Please fill in all fields before continuing...", 405, w)
		return
	}

	postid , err := models.InsertPost(postt.Title, postt.Content,userID, db)
	if err != nil {
		controllers.Response("error in the server...", 500, w)
		fmt.Println("error to insert post..<<--", err)
		return
	}

	if postid == 0 {
		
			controllers.Response("error in the server...", 500, w)
			fmt.Println("error to insert post..<<--", err)
			return
		
	}

	err = models.InsertCategory( "coding",postid, db)
	if err != nil {
		controllers.Response("error in the server...", 500, w)
		fmt.Println("error to insert post..<<--", err)
		return
	}

	controllers.Response("your post has created", http.StatusCreated, w)

}
