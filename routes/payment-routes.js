const express = require('express');
const mongoose = require('mongoose');
const router = express.Router();
const Buyer = require('../models/Buyer');
const Card = require('../models/Card');
const Payment = require('../models/Payment');

// GET Route => to get all the payments
router.get('/payment', (req, res) => {
  Payment.find()
    .then(allThePayments => {
      res.status(200).json(allThePayments);
    })
    .catch(err => {
      res.json(err);
    });
});

// GET Route => to get an specific payment by id
router.get('/payment/:id', (req, res) => {

  if (!mongoose.Types.ObjectId.isValid(req.params.id)) {
    res.status(400).json({ message: 'Specified id is not valid' });
    return;
  }

  Payment.findById(req.params.id)
    .then(payment => {
      res.status(200).json(payment);
    })
    .catch(err => {
      res.json(err);
    })
});

//POST route => to create a new payment
router.post('/payment', (req, res) => {
  const { client, buyerName, buyerEmail, buyerCpf, amount, type, cardHolderName, cardNumber, cardExpirationDate, cardCvv } = req.body;

  Buyer.findOne({ buyerCpf })
    .then(foundBuyer => {
      if (foundBuyer) {
        return;
      } 
  
      const aNewBuyer = new Buyer({
        buyerName: buyerName,
        buyerEmail: buyerEmail,
        buyerCpf: buyerCpf
      });
  
      aNewBuyer.save(err => {
        if (err) {
          res.status(400).json({ message: 'Saving Buyer to database went wrong.' });
          return;
        }
      })
    })
    .catch((err) => {
      console.log(err)
    });

  if (type.toUpperCase() === 'CARD') {

    Card.findOne({ cardNumber }) 
    .then( foundCard => {
      if (foundCard) {
        return;
      } 

      const aNewCard = new Card({
        cardNumber: cardNumber,
        cardHolderName: cardHolderName,
        cardExpirationDate: cardExpirationDate,
        cardCvv: cardCvv,
        cardIssuer: cardIssuer(cardNumber)
      });

      aNewCard.save(err => {
        if (err) {
          res.status(400).json({ message: 'Saving Card to database went wrong.' });
          return;
        }
      })

      Payment.create({
        amount: amount,
        type: type.toUpperCase(),
        cardNumber: cardNumber,
        cardHolderName: cardHolderName,
        cardExpirationDate: cardExpirationDate,
        cardCvv: cardCvv,
        cardIssuer: cardIssuer(cardNumber),
        buyerName: buyerName,
        buyerEmail: buyerEmail,
        buyerCpf: buyerCpf,
        client: client,
        status: paymentStatus(cardNumber)
      })
      .then(payment => {
        res.status(200).json(payment.status);
      })
      .catch(err => {
        res.json(err);
      })  
    })
    .catch(err => {
      console.log(err);
    })
  } else if (type.toUpperCase() === 'BOLETO') {

    const newBoleto = boleto();

    Payment.create({
      amount: amount,
      type: type.toUpperCase(),
      boleto: newBoleto,
      buyerName: buyerName,
      buyerEmail: buyerEmail,
      buyerCpf: buyerCpf,
      client: client,
      status: paymentStatus(newBoleto)
    })
    .then(payment => {
      res.status(200).json(payment.boleto);
    })
    .catch(err => {
      res.json(err);
    })  
  };
});

// boleto generator fake
const boleto = () => {
  let boletoNum = [];
  while (boletoNum.length < 48) {
    boletoNum.push(Math.round(Math.random() * 10));
  }
  return (boletoNum.join(''));
};

// Card validator
function cardIssuer(input) {
  const checkVisa =  /^(?:4[0-9]{12}(?:[0-9]{3})?)$/;
  const checkMasterCard = /^(?:5[1-5][0-9]{14})$/;
	if (checkVisa.test(input)) {
    return "Visa";
  } else if (checkMasterCard.test(input)) {
    return "MasterCard";
  } else {
    return "Invalid Card";
  }
}

// Payment status
function paymentStatus(num) {
  if (num < 48) {
    const validate = cardIssuer(num);

    if (validate === "Invalid Card") {
      return "rejected"
    } 

  } else {
    var approved = Math.floor(Math.random() * 2);
    if(approved === 0) {
      return "approved";
    } else {
      return "rejected";
    }
  }
};

module.exports = router;