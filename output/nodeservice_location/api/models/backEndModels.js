'use strict';

const mongoose = require('mongoose')
const Schema = mongoose.Schema;

var ModelSchema = new Schema ({
  z : Number,
player : String,
x : Number,
y : Number,

});

module.exports = mongoose.model('Models', ModelSchema);
