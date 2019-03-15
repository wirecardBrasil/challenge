const mongoose = require('mongoose');

const Schema = mongoose.Schema;

const cardSchema = new Schema({
  cardHolderName: {type: String, required: true },
  cardIssuer: String,
  cardNumber: {type: Number, required: true },
  cardExpirationDate: {type: Number, required: true },
  cardCvv: {type: Number, required: true }, 
});

const Card = mongoose.model('card', cardSchema);

module.exports = Card;
