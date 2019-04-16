# Wirecard Backend Challenge

At Wirecard we cherish our startup culture and our engineers are active participants in the innovation process. On this challenge, we’ll be looking at how you organize your code, your knowledge of design patterns, data persistence, unit testing, TDD, etc.

You’ll need to build a simple application to provide an API for payment to our customers and a checkout for who don't want to integrate with our API. In other words, you'll recreate a small part of the Wirecard :)

## Level 1: The API

We need to create a new API to process payments for our customers.
That been said, we need:

#### 1. An endpoint to create payments
- The API needs to accept two payment methods: Credit Card and Boleto.
- When the payment method is boleto, we only need to return the boleto's number in our response.
- When the method is card, we need to return if it was successful or not *(please don't worry about processing the payment, just mock the answers)*.
- The API must receive the information of the buyer, customer and payment. The information needed is the following:
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

#### 3. Unit Tests
Unit tests are required

## Level 2: The Checkout

We need a way to use the API: It could be a simple checkout page (or a simple order completion page, you don't need to care about design or layout), some `curl`'s examples or `Postman`'s documentation.

This step is focused on seeing how you integrate APIs.

In this checkout, we need:
- To send the buyer information with the payment method.
- To validate if the card is valid and who is the card issuer.
- To simulate a form of identification of the buyer that will be sent to the API.
- To return if the transaction was successful or not.
- To persist and consume the data effectively for this test.

## Level 3: The documentation

Please take a time to write a README file explaining how to:
- Run your project
- The architecture and the design adopted to solve the challenges.

# Additional informations
This challenge can be done in any language and any database.

Good luck!

*PS: Using docker is a plus! :D*

*PS2: Look to our public repository (https://github.com/wirecardBrasil). Maybe something can help you in this test :D*
