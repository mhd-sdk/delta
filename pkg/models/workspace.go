package models

type Workspace struct {
	Name   string `json:"name"`
	Layout []Tile `json:"layout"`
}
