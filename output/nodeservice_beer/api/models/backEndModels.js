'use strict';

const mongoose = require('mongoose')
const Schema = mongoose.Schema;

var ModelSchema = new Schema ({
  ingradientes : String,
sabores : String,
name : String,
grados : String,

});

module.exports = mongoose.model('Models', ModelSchema);
