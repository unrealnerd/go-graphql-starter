package gql

import (

	"github.com/graphql-go/graphql"

	"go-graph-starter/repo"

)

var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: graphql.NewList(userType),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					_, isOK := p.Args["id"].(string)
					if isOK {
						r := repo.Repo{}
						return r.GetData(), nil
					}
					return nil, nil
				},
			},
		},
	})
