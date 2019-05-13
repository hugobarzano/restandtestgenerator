//Tools

const express = require('express');
const bodyParser = require('body-parser');
const mongoose = require('mongoose')

//Server app
const backEndCore = express();



//back end models
var Model = require('./api/models/backEndModels');

//persistency layer
mongoose.Promise = global.Promise;
mongoose.connect("mongodb://cesar01:cesar01@ds155516.mlab.com:55516/cesarcorplab01",{useNewUrlParser: true});

//Parse content-type x-Www
backEndCore.use(bodyParser.urlencoded({ extended: true }))

//Parse content-type - Json
backEndCore.use(bodyParser.json())


// Back End routes
var routes = require('./api/routes/backEndRoutes'); //importing route
routes(backEndCore); //register the route


//middleware

backEndCore.use(function(req, res) {
  res.status(404).send({url: req.originalUrl + ' not found'})
});

//RUN
var port = process.env.PORT || 3006;
backEndCore.listen(port);
console.log(new Date()+ ' -- RESTFUL API server started on: '+port)
