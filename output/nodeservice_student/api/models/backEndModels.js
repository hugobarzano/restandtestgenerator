'use strict';

const mongoose = require('mongoose')
const Schema = mongoose.Schema;

var ModelSchema = new Schema ({
  name : String,

});

module.exports = mongoose.model('Models', ModelSchema);
