package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	// "database/sql"
	"fmt"
	"strconv" //para convertir Integer -> String 
	_ "github.com/lib/pq"
)


/* 
	?Modelos, estructuras en C++ 
	Todos los nombres deben empezar con mayúsculas
*/
type Post struct {
	ID int `json:"id_post,omitempty"`
	Categoria int `json:"categoria,omitempty"`
	Nombre string `json:"nombre,omitempty"`
	Descripcion string `json:"descripcion"`
	Url string `json:"url,omitempty"`
	Image string `json:"image"`
}

type ErrorMessage struct {
	Codigo int `json:"codigo,omitempty"`
	Nombre string `json:"nombre,omitempty"`
	Descripcion string `json:"descripcion,omitempty"`
}

/* Base de datos ficticia */
var posts []Post

func GetPostsEndpoint(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func GetOnePostEndpoint(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	w.Header().Set("Content-Type", "application/json")
	for _, item := range posts {
		if strconv.Itoa(item.ID) == params["id"] { // Devuelve un json añadiendo a la cabecera de la respuesta Content-Type: JSON
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// Creando el mensaje de error a devolver en caso no exista el id
	var errorDevolver ErrorMessage
	errorDevolver.Codigo = 500
	errorDevolver.Nombre = "Error GetOnePost"
	errorDevolver.Descripcion = fmt.Sprintf("No existe el post con id %s", params["id"])
	
	json.NewEncoder(w).Encode(&errorDevolver)
}

func CreatePostEndpoint(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req) //Obtener parametros en variable params
	var post Post
	_ = json.NewDecoder(req.Body).Decode(&post)
	i, err  := strconv.Atoi(params["id"]) //Transformar String a Integer
	if err == nil {
		post.ID = i
		posts = append(posts, post)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(post)
	}
}

func EditPostEndpoint(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	w.Header().Set("Content-Type", "application/json")
	var errorDevolver ErrorMessage
	for index, item := range posts {
		if strconv.Itoa(item.ID) == params["id"] {
			// Creamos un 'objeto' post y lo llenamos con el body de la consulta
			var post Post
			_ = json.NewDecoder(req.Body).Decode(&post)
			i, err  := strconv.Atoi(params["id"]) //Transformar String a Integer
			if err != nil {
				errorDevolver.Codigo = 500
				errorDevolver.Nombre = "Error EditPostEndpoint"
				errorDevolver.Descripcion = fmt.Sprintf("Error al pasar el ID tipo string a entero")
				json.NewEncoder(w).Encode(&errorDevolver)
			}
			post.ID = i
			item.Categoria = post.Categoria
			item.Descripcion = post.Descripcion
			item.Nombre = post.Nombre
			item.Url = post.Url
			item.Image = post.Image
			posts[index] = item
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	errorDevolver.Codigo = 204
	errorDevolver.Nombre = "Error EditPostEndpoint"
	errorDevolver.Descripcion = fmt.Sprintf("No existe el post con id %s para editar", params["id"])
	json.NewEncoder(w).Encode(&errorDevolver)
}

func DeletePostEndpoint(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	w.Header().Set("Content-Type", "application/json") //devolver JSON
	for index, item := range posts {
		if strconv.Itoa(item.ID) == params["id"] {
			posts = append(posts[:index], posts[index + 1:]...) // remover el item encontrado
			// return
			json.NewEncoder(w).Encode(&posts)
			return
		}
	}
	// Creando el mensaje de error a devolver en caso no exista el id
	var errorDevolver ErrorMessage
	errorDevolver.Codigo = 500
	errorDevolver.Nombre = "Error GetOnePost"
	errorDevolver.Descripcion = fmt.Sprintf("No existe el post con id %s", params["id"])

	json.NewEncoder(w).Encode(&errorDevolver)
}

func main(){
	router := mux.NewRouter()

	/* Llenando datos manualmente */
	posts = append(posts, Post{ID: 1, Categoria: 1, Nombre: "Primer Post", Descripcion: "Post de prueba", Url: "www.google.com", Image: "www.google.com/images"})
	posts = append(posts, Post{ID: 2, Categoria: 1, Nombre: "Segundo Post", Url: "www.youtube.com"})



	/* Endpoints */
	router.HandleFunc("/posts", GetPostsEndpoint).Methods("GET")
	router.HandleFunc("/post/{id}", GetOnePostEndpoint).Methods("GET")
	router.HandleFunc("/post/{id}", CreatePostEndpoint).Methods("POST")
	router.HandleFunc("/post/{id}", EditPostEndpoint).Methods("PUT")
	router.HandleFunc("/post/{id}", DeletePostEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}