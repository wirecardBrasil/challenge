const mongoose = require('mongoose');

const Schema = mongoose.Schema;

const buyerSchema = new Schema({
  buyerName: {type: String, required: true}, 
  buyerEmail: {type: String, required: true},
  buyerCpf: {type: Number, required: true}
});

const Buyer = mongoose.model('buyer', buyerSchema);

module.exports = Buyer;
