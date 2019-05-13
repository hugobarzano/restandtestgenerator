'use strict';

module.exports = function(backEndCore){
  var backEndDrivers=require('../drivers/backEndDrivers');

  //backEndROUTES
  //var route='/models'
  //<<API_PLACEHOLDER>>
  backEndCore.route(route)
    .get(backEndDrivers.list_all_models)
    .post(backEndDrivers.create_model);

  backEndCore.route(route+'/:modelId')
     .get(backEndDrivers.read_model)
     .put(backEndDrivers.update_model)
     .delete(backEndDrivers.delete_model);  


};
