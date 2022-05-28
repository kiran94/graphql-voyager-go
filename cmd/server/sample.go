package main

import "fmt"

// Repository provides dummy data for sample app
type Repository interface {
	QueryTutorial() []Tutorial
    QueryAuthors() []Author
}

type dataRepository struct{}

func (r dataRepository) QueryTutorial() []Tutorial {

	author := Author{Name: "JK Rowling", Tutorials: []int{1}}
	tutorial := Tutorial{ID: 1, Title: "GraphQL Tutorial", Author: author}

	for i := 0; i < 5; i++ {
		newComment := Comment{Body: fmt.Sprintf("My %d Comment!", i+1)}
		tutorial.Comments = append(tutorial.Comments, newComment)
	}

	return []Tutorial{tutorial}
}

func (r dataRepository) QueryAuthors() []Author {
	author := Author{Name: "JK Rowling", Tutorials: []int{1}}
	author2 := Author{Name: "Stephen King", Tutorials: []int{}}
	return []Author{author, author2}
}
