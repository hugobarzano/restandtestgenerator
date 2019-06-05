'use strict';

const mongoose = require('mongoose')
const Schema = mongoose.Schema;

var ModelSchema = new Schema ({
  name : String,
alias : String,
age : Number,
color : String,

});

module.exports = mongoose.model('Models', ModelSchema);
