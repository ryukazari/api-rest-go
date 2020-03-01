package models

import (
	"fmt"
	"errors"
	"api-rest-v1/database"
	_ "database/sql"
)

/* Estructura de la tabla Post */
type Post struct {
	ID int `json:"id_post,omitempty"`
	Categoria int `json:"categoria,omitempty"`
	Nombre string `json:"nombre,omitempty"`
	Descripcion string `json:"descripcion"`
	Url string `json:"url,omitempty"`
	Image string `json:"image"`
}

// func [nombre de la funcion]([parametros]) [Tipos de dato que retornara la funcion ] {...}
func CreatePost(post Post) error {
	query := `INSERT INTO 
				api_rest_v1.post (categoria, nombre, descripcion, enlace, imagen)
				VALUES ($1, $2, $3, $4, $5)`
	connection := database.GetConnectionDB()
	defer connection.Close() //Cierra la conexión
	stmt, err := connection.Prepare(query) //Preparando la sentencia: reemplazando los valores $
	if err != nil {
		return err
	}
	defer stmt.Close()

	// el EXEC se utiliza para INSERT, DELET y UPDATE, devuelve la cantidad de registros afectados
	r, err := stmt.Exec(post.Categoria, post.Nombre, post.Descripcion, post.Url, post.Image)
	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New(`ERROR: Se esperaba una fila afectada`)
	}

	return nil //retorna nil porque no existe errores
}

func DeletePost(id int) error {
	query := `DELETE FROM
				api_rest_v1.post
				WHERE post.id_post = $1`
	connection := database.GetConnectionDB()
	defer connection.Close() //Cierra la conexión
	stmt, err := connection.Prepare(query) //Preparando la sentencia: reemplazando los valores $
	if err != nil {
		return err
	}
	defer stmt.Close()

	// el EXEC se utiliza para INSERT, DELET y UPDATE, devuelve la cantidad de registros afectados
	r, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New(`ERROR: Se esperaba una fila afectada`)
	}

	return nil //retorna nil porque no existe errores
}

func SelectOnePost(id int) (error, Post) {
	var post Post
	query := `SELECT id_post, categoria, nombre, descripcion, enlace, imagen
				FROM api_rest_v1.post
				WHERE id_post = $1`
	connection := database.GetConnectionDB()
	defer connection.Close() //Cierra la conexión
	stmt, err := connection.Prepare(query) //Preparando la sentencia: reemplazando los valores $
	if err != nil {
		return err, post
	}
	defer stmt.Close()
	rows, err := stmt.Query(id) 
	if err != nil {
		return err, post
	}
	defer rows.Close()
	for rows.Next() {
		p := Post{}
		err = rows.Scan(&p.ID, &p.Categoria, &p.Nombre, &p.Descripcion, &p.Url, &p.Image)
		if err != nil {
			fmt.Println("Error Scan...")
			return err, post
		} // retorna posts que sería un arreglo vacío y err que sería un error
		post = p
	}
	return nil, post
}

// se nombran los parámetros y al poner return devolverán los valores que tienen en ese momento
func SelectAllPosts() (posts []Post,err error) {
	query := `SELECT id_post, categoria, nombre, descripcion, enlace, imagen
				FROM api_rest_v1.post`
	connection := database.GetConnectionDB()
	defer connection.Close()

	rows, err := connection.Query(query)
	if err != nil { 
		fmt.Println("Error Query...")
		return 
	}
	defer rows.Close()
	for rows.Next() {
		p := Post{}
		err = rows.Scan(&p.ID, &p.Categoria, &p.Nombre, &p.Descripcion, &p.Url, &p.Image)
		if err != nil {
			fmt.Println("Error Scan...")
			return 
		} // retorna posts que sería un arreglo vacío y err que sería un error
		posts = append(posts, p)
	}
	return posts, nil
}

func UpdatePost(post Post, id int) error {
	query := `UPDATE api_rest_v1.post 
				SET categoria = $1, nombre = $2, descripcion = $3, enlace = $4, imagen =$5
				WHERE id_post = $6`
	connection := database.GetConnectionDB()
	defer connection.Close() //Cierra la conexión
	stmt, err := connection.Prepare(query) //Preparando la sentencia: reemplazando los valores $
	if err != nil {
		return err
	}
	defer stmt.Close()

	// el EXEC se utiliza para INSERT, DELET y UPDATE, devuelve la cantidad de registros afectados
	r, err := stmt.Exec(post.Categoria, post.Nombre, post.Descripcion, post.Url, post.Image, id)
	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New(`ERROR: Se esperaba una fila afectada`)
	}

	return nil //retorna nil porque no existe errores
}