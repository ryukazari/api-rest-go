package models

type Categoria struct {
	ID int `json:"id_post,omitempty"`
	Nombre string `json:"nombre,omitempty"`
	Descripcion string `json:"descripcion"`
}