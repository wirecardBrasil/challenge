const mongoose = require('mongoose');

const Schema = mongoose.Schema;

const paymentSchema = new Schema({
  amount: {type: Number, required: true},
  type: {
    type: String,
    required: true,
    enum: ['CARD', 'BOLETO']
  }, 
  boleto: String,
  cardHolderName: String,
  cardIssuer: String,
  cardNumber: String,
  cardExpirationDate: String,
  cardCvv: String, 
  buyerName: {type: String, required: true}, 
  buyerEmail: {type: String, required: true},
  buyerCpf: {type: Number, required: true},
  client: { type: String, required: true},
  status: String
}, { 
  timestamps: true,
});

const Payment = mongoose.model('payment', paymentSchema);

module.exports = Payment;
