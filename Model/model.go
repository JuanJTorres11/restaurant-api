package Model

import (
	"context"
	"log"

	"github.com/machinebox/graphql"
)

var client = graphql.NewClient("http://localhost:8080/graphql")

func GetBuyers() (interface{}, error) {
	ctx := context.Background()

	q := graphql.NewRequest(`
	query {
		queryBuyer {
			id
			name
			age
		  }
	}
	`)
	var resp interface{}
	err := client.Run(ctx, q, &resp)
	if err != nil {
		log.Println(err)
		log.Println(resp)
	}

	return resp, err

}
