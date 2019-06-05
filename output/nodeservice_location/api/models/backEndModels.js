'use strict';

const mongoose = require('mongoose')
const Schema = mongoose.Schema;

var ModelSchema = new Schema ({
  player : String,
x : Number,
y : Number,
z : Number,

});

module.exports = mongoose.model('Models', ModelSchema);
