package Model

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/machinebox/graphql"
)

const BUYERS_ENDPOINT = "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/buyers?date="
const TRANSACTIONS_ENPOINT = "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/transactions?date="
const PRODUCTS_ENDPOINT = "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/products?date="

var client = graphql.NewClient("http://localhost:8080/graphql")

func GetBuyers(date string) (interface{}, error) {
	response, err := http.Get(BUYERS_ENDPOINT + date)
	if err != nil {
		log.Panicln("There was an error trying to connect to the buyers endpoint")
	}
	defer response.Body.Close()

	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Panicln("There was an error trying to read the JSON from the buyers endpoint")
	}

	formatedReponse := formatBuyers(res)

	return putBuyers(formatedReponse)
}

func ListBuyers() (interface{}, error) {
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

func putBuyers(newBuyers string) (interface{}, error) {
	ctx := context.Background()

	q := graphql.NewRequest(`
	mutation {
		addBuyer(input: ` + newBuyers + `, upsert: true) {
			numUids
	  }
	}
	`)

	log.Println(q)
	var resp interface{}
	err := client.Run(ctx, q, &resp)
	if err != nil {
		log.Println(err)
	}

	return resp, err
}
