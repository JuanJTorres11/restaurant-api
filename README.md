# restaurant-api

API built in Go and GraphQL to upload and consult clients and transactions data from a restaurant.

# How to deploy

1. Download the Dgraph database with `curl -sSf https://get.dgraph.io | bash`
2. To start the database open two terminals, in one write `dgraph alpha`, in the other write `dgraph zero`
3. Create the schema of the database with `curl -X POST localhost:8080/admin/schema --data-binary '@schema.graphql'`
4. Use `go run main.go` to start the backend server and start receiving petitions to the API in localhost:8000

# API Endpoints

The endpoints of the API are:

## load/*{date}*

It makes a petition to a data source to retrieve information about buyers, products and transactions. It can receive a date in UNIX timestamp format optionally to retrieve the data of that date, if a date is not given, it will default to the current date.

## buyers

This retrieves the information about the all the buyers.

## buyers/{id}

Given an ID, this will retrieve the information about the specified buyer, all the transaction that the person has made and some product recommendations based on what other people have bought.
