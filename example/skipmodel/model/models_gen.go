// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type ProfileInput struct {
	Name string `json:"name"`
}

type UserInput struct {
	ID      string        `json:"id"`
	Profile *ProfileInput `json:"profile"`
	Friends []*UserInput  `json:"friends"`
}