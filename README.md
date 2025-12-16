# overview

This project is a backend-only RESTful API for a bus ticket booking system.
It provides endpoints for managing users, routes, buses, seat reservations,
and ticket bookings.  
The service is designed to be consumed by any frontend client (web or mobile).

## requirements

go.mod</br>
go.sum

## go version

go 1.25.0

## data model
Main entities:
- user
- bus
- seat
- ticket
- trip
- payment
- company

## Tech Stack

- Golang
- Gin
- MySQL

## ERD

<img width="1421" height="870" alt="Untitled (2)" src="https://github.com/user-attachments/assets/138d9e8d-4a5f-4019-9622-e468396da3e9" />

for more detail watch this [link](https://dbdiagram.io/d/68a331fcec93249d1e1717ec)


## Design Decisions

- The project is implemented as an API-only service with no UI layer.
- REST architecture is used for simplicity and interoperability.
- Database access is abstracted via repository layer.
- each module separated to individual package(service,controller,model,config,db).

## Features

- User registration and login
- Bus and route management APIs
- Seat availability tracking
- Ticket booking and cancellation
- Trip management APIs
- Company management APIs
- Payment management for buy ticket

## The service follows a layered architecture:

- controller layer: HTTP request/response handling
- Service layer: business logic
- db layer: database access
- Model layer: domain entities
- config: environment variable and other configs for ptoject

## Project Structure

BUSTICKETWITHGOLANG/</br>
├── config/</br>
├── controller/</br>
├── db/</br>
├── model/</br>
├── service/</br>
├── .env</br>
├── go.mod</br>
├── go.sum</br>
├── main.go</br>
└── README.md</br>

## Postman Collection

You can explore and test the API endpoints using the Postman collection below:

🔗 https://manizm1382-8786911.postman.co/workspace/mani-z's-Workspace~946e1bcc-81fb-4edf-8d1d-f632d52d1069/collection/50450074-32c5d9ff-5d12-4ebf-8498-4cf250e452c0?action=share&creator=50450074


## Configuration

The service is configured using environment variables:

config these parameters of .env file base on individual setting:

- userName
- passWord
- host
- port
- dbName
- dsn = "[userName]:[passWord]@tcp([host]:[port)/[dbName]?parseTime=true"
- refTime = "2006-01-02 15:04:05"
- connectionPort

## run project

execute this command:
``` bash
go run main.go
```











