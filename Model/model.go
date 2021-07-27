package Model

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/machinebox/graphql"
)

type QueryBuyer struct {
	Buyer SimpleBuyer `json:"getBuyer,omitempty"`
}

type QueryTransaction struct {
	Buyers []TransactionBuyer `json:"queryTransaction,omitempty"`
}

type queryProduct struct {
	Product []ProductTransactions `json:"queryProduct,omitempty"`
}

const BUYERS_ENDPOINT = "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/buyers?date="
const TRANSACTIONS_ENDPOINT = "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/transactions?date="
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

func GetBuyer(id string) (QueryBuyer, []string, []string, error) {
	ctx := context.Background()

	q1 := graphql.NewRequest(`
	query ($id: String!) {
		getBuyer(id: $id) {
		  name
		  age
		  transactions {
			ip
			products {
			  id
			  name
			  price
			}
		  }
		}
	}
	`)

	q1.Var("id", id)
	var resp QueryBuyer
	err := client.Run(ctx, q1, &resp)
	if err != nil {
		log.Panicln(err)
	}

	ips, ids := obtainIpsIds(resp)

	buyerNames := queryTransaction(ips)

	productNames := queryProducts(ids)

	return resp, buyerNames, productNames, err
}

func obtainIpsIds(result QueryBuyer) ([]string, []string) {
	var ips []string
	var ids []string
	mapId := make(map[string]bool)
	mapIp := make(map[string]bool)

	for _, st := range result.Buyer.Transactions {
		if _, value := mapIp[st.IP]; !value {
			mapIp[st.IP] = true
			ips = append(ips, st.IP)
		}
		for _, pi := range st.Products {
			if _, val := mapId[pi.ID]; !val {
				mapId[pi.ID] = true
				ids = append(ids, pi.ID)
			}
		}

	}
	return ips, ids
}

func queryTransaction(ips []string) []string {
	ctx := context.Background()

	q1 := graphql.NewRequest(`
	query ($ip: String!) {
		queryTransaction (filter: {ip: {anyofterms: $ip}}) {
		  buyer {
			name
		  }
		}
	}
	`)

	q1.Var("ip", strings.Join(ips, " "))
	var resp QueryTransaction
	err := client.Run(ctx, q1, &resp)
	if err != nil {
		log.Panicln(err)
	}

	var names []string
	mapNames := make(map[string]bool)

	for _, v := range resp.Buyers {
		if _, val := mapNames[v.Buyer.Name]; !val {
			mapNames[v.Buyer.Name] = true
			names = append(names, v.Buyer.Name)
		}
	}

	return names
}

func queryProducts(ids []string) []string {
	ctx := context.Background()

	q1 := graphql.NewRequest(`
	query ($id: [String!]!) {
		queryProduct (filter: {id: {in: $id}}) {
		  transactions {
			products (filter: { not: {id: {in: $id}}}){
			  name
			}
		  }
		}
	}
	`)

	q1.Var("id", ids)
	var resp queryProduct
	err := client.Run(ctx, q1, &resp)
	if err != nil {
		log.Panicln(err)
	}

	var names []string
	mapNames := make(map[string]bool)

	for _, v := range resp.Product {
		for _, pn := range v.Transactions {
			for _, pn2 := range pn.Products {
				if _, val := mapNames[pn2.Name]; !val {
					mapNames[pn2.Name] = true
					names = append(names, pn2.Name)
					if len(names) > 9 {
						break
					}
				}
			}
			if len(names) > 9 {
				break
			}
		}
		if len(names) > 9 {
			break
		}
	}

	return names
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
		log.Panicln(err)
	}

	return resp, err
}

func putBuyers(newBuyers []Buyer) (interface{}, error) {
	ctx := context.Background()

	q := graphql.NewRequest(`
	mutation ($data: [AddBuyerInput!]!){
		addBuyer(input: $data , upsert: true){
			numUids
	  }
	}
	`)

	q.Var("data", newBuyers)
	var resp interface{}
	err := client.Run(ctx, q, &resp)
	if err != nil {
		log.Panicln(err)
	}

	return resp, err
}

func GetProducts(date string) (interface{}, error) {
	response, err := http.Get(PRODUCTS_ENDPOINT + date)
	if err != nil {
		log.Panicln("There was an error trying to connect to the products endpoint")
	}
	defer response.Body.Close()

	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Panicln("There was an error trying to read the CSV from the products endpoint")
	}

	formatedProducts := formatProducts(res)

	return putProducts(formatedProducts)
}

func putProducts(newProducts []Product) (interface{}, error) {
	ctx := context.Background()

	q := graphql.NewRequest(`
	mutation ($data: [AddProductInput!]!){
		addProduct(input: $data , upsert: true){
		numUids
	  }
	}
	`)

	q.Var("data", newProducts)
	var resp interface{}
	err := client.Run(ctx, q, &resp)
	if err != nil {
		log.Panicln(err)
	}

	return resp, err
}

func GetTransactions(date string) (interface{}, error) {
	response, err := http.Get(TRANSACTIONS_ENDPOINT + date)
	if err != nil {
		log.Panicln("There was an error trying to connect to the transactions endpoint")
	}
	defer response.Body.Close()

	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Panicln("There was an error trying to read the document from the transactions endpoint")
	}

	formatedTransactions := formatTransactions(res)

	return putTransactions(formatedTransactions)
}

func putTransactions(newTransactions []Transaction) (interface{}, error) {
	ctx := context.Background()

	q := graphql.NewRequest(`
	mutation ($data: [AddTransactionInput!]!){
		addTransaction(input: $data , upsert: true){
		numUids
	  }
	}
	`)

	q.Var("data", newTransactions)
	var resp interface{}
	err := client.Run(ctx, q, &resp)
	if err != nil {
		log.Panicln(err)
	}

	return resp, err
}
