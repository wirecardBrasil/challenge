require('dotenv').config();

const express = require('express');
const bodyParser = require('body-parser');
const mongoose = require('mongoose');

//connect DB
mongoose
  .connect('mongodb://localhost/payment-api', { useNewUrlParser: true })
  .then((x) => {
    console.log(`Connected to Mongo! Database name: ${x.connections[0].name}`);
  })
  .catch((err) => {
    console.error('Error connecting to mongo', err);
  });

const app = express();

app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());

app.use('/', require('./routes/payment-routes'));
const port = process.env.port || 3000;

// Starting server
app.listen(port, () => {
    console.log(`Listening on port ${port}!`);
  });
  
module.exports = app;
