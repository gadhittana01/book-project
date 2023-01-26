# Book Project
This repository is the API where you can browse and reserve a schedule to borrow a book.

# Pre-requisite to contribute
To run the service on your local machine, you must have these components installed on your local machine : 
- Golang compiler

# Getting started
- To run this project you can use this command below in the root of the repository:
```sh
// install the dependency
$ go get 

// start the http server
$ make run-http-server-local 
```

# Example Request
```sh
// Get all Book by Subject
$ curl --location --request GET 'http://localhost:8000/get-books?subject=love'

// Reserve a book pickup schedule
$ curl --location --request POST 'http://localhost:8000/borrow-book' \
--header 'Content-Type: application/json' \
--data-raw '{
    "key" : "/works/OL98501W",
    "pickup_date" : "2022-02-26",
    "subject" : "love",
    "user_id" : 2
}'

// Get All Book Reservation
$ curl --location --request GET 'http://localhost:8000/get-book-reservation'
```