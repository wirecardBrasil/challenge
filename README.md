# Moip Technical Challenge

At Moip we cherish our startup culture and our engineers are active participants in the innovation process. On this challenge, ee’ll be looking at how you organize your code, your knowledge of design patterns, data persistence, unit testing, TDD, etc.

You’ll need to build a simple application to provide an API for payment to your customers and a checkout for who don't want to integrate with your API. In other words, you'll recreate a small part of the Moip :)

## Level 1: The API

We need to create a new API to process payments from our clients. 
That been said, we need:

#### 1. An endpoint to create payments
- The API needs to accept two payment methods: Credit Card and Boleto. 
- When the payment method is boleto, we only need to return the boleto's number in our response.
- When the method is card, we need to return if it was successful or not *(please don't worry about processing the payment, just mock the answers)*.
- The API must receive the information of the buyer, of our client and of the payment. The information needed is the following:
```
Client:
 - ID

Buyer:
 - Name
 - Email
 - CPF

Payment:
 - Amount
 - Type
 - Card (when the payment type is credit card)
 
Card:
 - Card holder name
 - Card number
 - Card expiration date
 - Card CVV (Number behind the card)

```

#### 2. An endpoint to check the payment status
- The API needs to return all the information about the payment, as well as the status of that payment.


## Level 2: The Checkout

For those clients that won't integrate with our API, we need to create a simple checkout page (or a simple order completion page).

In this checkout, we need:
- To insert the buyer information and choose the payment method.
- To validate if the card is valid and who is the card issuer.
- To simulate a form of identification of the client that will be sent to the API.
- To show on the screen if the transaction was successful or not.
- To persist and consume the data effectively for this test.


# Additional informations
This challenge can be done in any language and any database.

Please take a time to write a READ-ME file explaining how to run your project, the architecture and the design adopted to solve the challenges.

Good luck!

*PS: Using docker is a plus and we want unit tests! :D*
*PS2: Look to our public repository (https://github.com/moip). Maybe something can help you in this test :D*
