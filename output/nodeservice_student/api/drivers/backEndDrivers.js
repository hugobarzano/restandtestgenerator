'use strict';


var mongoose = require('mongoose');
var Models = mongoose.model('Models');

exports.list_all_models = function(req, res) {
  console.log("LIST MODELS::");

  Models.find({}, function(err, model_find) {
    if (err)
      res.send(err);
    res.json(model_find);
  });
};

exports.create_model = function(req, res) {
  console.log("POST MODEL::");
  var new_model = new Models(req.body);
  console.log(new_model)
  new_model.save(function(err, model_save) {
    if (err)
      res.send(err);
    res.json(model_save);
  });
};


exports.read_model = function(req, res) {
  console.log("GET MODEL:: \n")
  Models.findById(req.params.modelId, function(err, model_read) {
    if (err)
      res.send(err);
    if (model_read == null)
      res.status(404);
    res.json(model_read);
  });
};

exports.update_model = function(req, res) {
  console.log("PUT MODEL::");

  Models.findOneAndUpdate({_id: req.params.modelId}, req.body, {new: true}, function(err, model_updated) {
    if (err)
      res.send(err);
    res.json(model_updated);
  });
};


exports.delete_model = function(req, res) {
  console.log("DELETE MODEL::");

  Models.deleteOne({
    _id: req.params.modelId
  }, function(err, model_deleted) {
    if (err)
      res.send(err);
    res.json({ message: 'Model successfully deleted' });
  });
};
