'use strict';

const mongoose = require('mongoose')
const Schema = mongoose.Schema;

var ModelSchema = new Schema ({
  age : Number,
color : String,
name : String,
alias : String,

});

module.exports = mongoose.model('Models', ModelSchema);
