'use strict';

const mongoose = require('mongoose')
const Schema = mongoose.Schema;

var ModelSchema = new Schema ({
  job : String,
city : String,
name : String,
company : String,

});

module.exports = mongoose.model('Models', ModelSchema);
