package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"strconv" //para convertir Integer -> String 
	_ "github.com/lib/pq"
	"api-rest-v1/models"
)


/* 
	?Modelos, estructuras en C++ 
	Todos los nombres deben empezar con mayúsculas
*/


type ErrorMessage struct {
	Codigo int `json:"codigo,omitempty"`
	Nombre string `json:"nombre,omitempty"`
	Descripcion string `json:"descripcion,omitempty"`
}

/* Base de datos ficticia */
var posts []models.Post

func GetPostsEndpoint(w http.ResponseWriter, req *http.Request){
	var errorDevolver ErrorMessage
	w.Header().Set("Content-Type", "application/json")
	posts, err := models.SelectAllPosts()
	if err == nil {
		json.NewEncoder(w).Encode(posts)
		return 
	}
	log.Fatal(fmt.Sprintf("Error: %s", err))
	errorDevolver.Codigo = 500
	errorDevolver.Nombre = "Error GetPostsEndpoint"
	errorDevolver.Descripcion = fmt.Sprintf("Error al obtener todos los posts %s", err)
	json.NewEncoder(w).Encode(&errorDevolver)
	return
}

func GetOnePostEndpoint(w http.ResponseWriter, req *http.Request){
	var errorDevolver ErrorMessage
	params := mux.Vars(req)
	w.Header().Set("Content-Type", "application/json")
	i, err := strconv.Atoi(params["id"])
	if err != nil {
		errorDevolver.Codigo = 500
		errorDevolver.Nombre = "Error al convertir"
		errorDevolver.Descripcion = fmt.Sprintf("Error al convertir el parametro id %s a entero", params["id"])
		return
	}
	err, post := models.SelectOnePost(i)
	if err != nil {
		// Creando el mensaje de error a devolver en caso no exista el id
		errorDevolver.Codigo = 500
		errorDevolver.Nombre = "Error GetOnePost"
		errorDevolver.Descripcion = fmt.Sprintf("No existe el post con id %s... %s", params["id"], err)
		json.NewEncoder(w).Encode(&errorDevolver)
		return
	}
	json.NewEncoder(w).Encode(&post)
}

func CreatePostEndpoint(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req) //Obtener parametros en variable params
	var post models.Post
	var errorDevolver ErrorMessage
	_ = json.NewDecoder(req.Body).Decode(&post)
	i, err  := strconv.Atoi(params["id"]) //Transformar String a Integer
	if err != nil {
		errorDevolver.Codigo = 500
		errorDevolver.Nombre = "Error EditPostEndpoint"
		errorDevolver.Descripcion = fmt.Sprintf("Error al pasar el ID tipo string a entero")
		json.NewEncoder(w).Encode(&errorDevolver)
	}
	post.ID = i
	err = models.CreatePost(post)
	if err != nil {
		errorDevolver.Codigo = 500
		errorDevolver.Nombre = "Error Insertando en la tabla Post"
		errorDevolver.Descripcion = fmt.Sprintf("Error al insertar el Post en la tabla de post: %s", err)
		json.NewEncoder(w).Encode(&errorDevolver)
	}
	fmt.Println("Post creado en la base de datos...")
	posts = append(posts, post)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func EditPostEndpoint(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	w.Header().Set("Content-Type", "application/json")
	var errorDevolver ErrorMessage
	var post models.Post
	_ = json.NewDecoder(req.Body).Decode(&post)
	
	id, errorConvertir := strconv.Atoi(params["id"])
	if errorConvertir != nil {
		errorDevolver.Codigo = 500
		errorDevolver.Nombre = "Error GetOnePost"
		errorDevolver.Descripcion = fmt.Sprintf("Error al convertir el id: %s a un dato de tipo Integer", params["id"])
		json.NewEncoder(w).Encode(&errorDevolver)
		return
	}

	err := models.UpdatePost(post, id)
	if err != nil {
		errorDevolver.Codigo = 204
		errorDevolver.Nombre = "Error EditPostEndpoint"
		errorDevolver.Descripcion = fmt.Sprintf("No existe el post con id %s para editar. :(", params["id"])
		json.NewEncoder(w).Encode(&errorDevolver)
		return
	}
	json.NewEncoder(w).Encode(&post)
	return
}

func DeletePostEndpoint(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	var errorDevolver ErrorMessage

	w.Header().Set("Content-Type", "application/json") //devolver JSON
	id, errorConvertir := strconv.Atoi(params["id"])
	if errorConvertir != nil {
		errorDevolver.Codigo = 500
		errorDevolver.Nombre = "Error GetOnePost"
		errorDevolver.Descripcion = fmt.Sprintf("Error al convertir el id: %s a un dato de tipo Integer", params["id"])
		json.NewEncoder(w).Encode(&errorDevolver)
		return
	}
	err := models.DeletePost(id)
	// Creando el mensaje de error a devolver en caso no exista el id
	if err == nil {
		json.NewEncoder(w).Encode("Se eliminó el post")
		return
	}

	errorDevolver.Codigo = 500
	errorDevolver.Nombre = "Error GetOnePost"
	errorDevolver.Descripcion = fmt.Sprintf("No existe el post con id %s . :(", params["id"])
	json.NewEncoder(w).Encode(&errorDevolver)
}

func main(){
	router := mux.NewRouter()
	/* Endpoints */
	router.HandleFunc("/posts", GetPostsEndpoint).Methods("GET")
	router.HandleFunc("/post/{id}", GetOnePostEndpoint).Methods("GET")
	router.HandleFunc("/post/{id}", CreatePostEndpoint).Methods("POST")
	router.HandleFunc("/post/{id}", EditPostEndpoint).Methods("PUT")
	router.HandleFunc("/post/{id}", DeletePostEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}