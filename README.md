# Wirecard Technical Challenge

## Requirements
- Npm package
- MongoDb
- NodeJS

To run this application, you'll need Git, Node.js (which comes with Npm package) and MongoDB installed on your computer. 

## Getting Started

- Fork this repo
- Clone this repo

You can download [MongoDB](https://www.mongodb.com/)

## Running

Use the following commands to run this project:

```
# Go to challenge-1 folder
$ cd challenge-1
# Install dependencies
$ npm install
# Run back-end server
$ npm start

```
## Payments
### Credit Card
- Visa
- MasterCard
### Boleto

## Postman 
Use the following examples to test on postman:

```
- Post 
http://http://localhost:3000/payment
{
	"client": "83740y64o79073097685i6",
    "amount": 200,
    "type": "boleto",
    "buyerName": "maria",
    "buyerEmail": "maria@jose.com",
    "buyerCpf": "87658984322"
}

{
	"client": "83740y64o7907309775864",
    "amount": 200,
    "type": "card",
    "buyerName": "jose",
    "buyerEmail": "jose@jose.com",
    "buyerCpf": "12675678976",
    "cardHolderName":"Elaine",
    "cardNumber":"536789089053",
    "cardExpirationDate":"2020",
    "cardCvv":"2004"
}

```

## Build With:
Node.JS, Express.JS and MongoDB.

### To do
Unit tests. 
