package models

type Users struct {
	Id       int
	Username string
	Email    string
	HashedPassword string
}