package main

import "github.com/graphql-go/graphql"

type Tutorial struct {
	ID       int
	Title    string
	Author   Author
	Comments []Comment
}

type Author struct {
	Name      string
	Tutorials []int
}

type Comment struct {
	Body string
}

func GenerateSchema() (*graphql.Schema, error) {
	repo := DataRepository{}

	commentType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Comment",
			Fields: graphql.Fields{
				"body": &graphql.Field{Type: graphql.String},
			},
		},
	)

	authorType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Author",
			Fields: graphql.Fields{
				"name":      &graphql.Field{Type: graphql.String},
				"tutorials": &graphql.Field{Type: graphql.NewList(graphql.Int)},
			},
		},
	)

	tutorialType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Tutorial",
			Fields: graphql.Fields{
				"id":       &graphql.Field{Type: graphql.Int},
				"title":    &graphql.Field{Type: graphql.String},
				"author":   &graphql.Field{Type: authorType},
				"comments": &graphql.Field{Type: graphql.NewList(commentType)},
			},
		},
	)

	// Root Query
	fields := graphql.Fields{
		"tutorial": &graphql.Field{
			Type:        tutorialType,
			Description: "Represents a Tuorial",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.Int},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				tutorials := repo.QueryTutorial()

				id, ok := p.Args["id"].(int)
				if ok {
					for _, tutorial := range tutorials {
						if tutorial.ID == id {
							return tutorial, nil
						}
					}
				}

				return tutorials, nil
			},
		},
		"list": &graphql.Field{
			Type:        graphql.NewList(tutorialType),
			Description: "Get Full Tutorial List",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				tutorials := repo.QueryTutorial()
				return tutorials, nil
			},
		},
		"authors": &graphql.Field{
			Type: authorType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return repo.QueryAuthors(), nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		return nil, err
	}

	return &schema, nil
}
