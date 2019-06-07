'use strict';

const mongoose = require('mongoose')
const Schema = mongoose.Schema;

var ModelSchema = new Schema ({
  name : String,
company : String,
job : String,
city : String,

});

module.exports = mongoose.model('Models', ModelSchema);
