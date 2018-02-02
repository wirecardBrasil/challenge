# Moip's Technical Challenge

## First step: The API

Given a necessity, we need to create a new API to process payments from our clients. The API needs to accept two payment methods: Credit Card and Boleto. We need to create an endpoint to create the payment and another one to check the status of that payment. The following list describes the details of the fields:

Client:
 - ID

Buyer:
 - Name
 - Email
 - CPF

Payment:
 - Amount
 - Type
 - Card
 - Holder's Name
 - Card number
 - Card's expiration date
 - CVV (Number behind the card)

For the boleto method, there are no additional information required.

When the payment method is boleto, we only need to return the boleto's number in our response.
When the method is card, we need to return if it was successful or not.

* You don't need to worry about the acquire (Payment processors, such as Cielo, Rede, etc.), just mock the answers.

In the endpoint to check the payment status, we need to return all the information about the payment, as well as the status of that payment.

We want unit tests! :D

## Second step: The Checkout

For those clients that won't integrate with our API, we need to create a simple checkout (or a simple order completion page).

In this checkout, we need to have the option to insert the buyer's information and choose the payment method with the additional data.
We need to validate if the card is valid and what is the issuer. You need to simulate a form of identification of the client to be sent correctly to the API.

The answer should be shown on the screen for the client to know if the transaction was successful or not.

Lastly, it is necessary to persist and consume the data effectively for this test.

It can be done in any language, and any type of database.

Try to explain the architecture and the design adopted to solve the problems.

Plus - Build a docker container with the stack.
