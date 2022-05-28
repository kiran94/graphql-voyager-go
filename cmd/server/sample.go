package main

import "fmt"

type Repository interface {
	QueryTutorial() []Tutorial
}

type DataRepository struct{}

func (r DataRepository) QueryTutorial() []Tutorial {

	author := Author{Name: "JK Rowling", Tutorials: []int{1}}
	tutorial := Tutorial{ID: 1, Title: "GraphQL Tutorial", Author: author}

	for i := 0; i < 5; i++ {
		newComment := Comment{Body: fmt.Sprintf("My %d Comment!", i+1)}
		tutorial.Comments = append(tutorial.Comments, newComment)
	}

	return []Tutorial{tutorial}
}

func (r DataRepository) QueryAuthors() []Author {
	author := Author{Name: "JK Rowling", Tutorials: []int{1}}
	author2 := Author{Name: "Stephen King", Tutorials: []int{}}
	return []Author{author, author2}
}
