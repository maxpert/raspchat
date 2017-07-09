if (!Function.prototype.bind){
  Function.prototype.bind = function (scope) {
    var fn = this;
    return function () {
        var args = window.$arrayify(arguments);
        return fn.apply(scope, args);
      };
  };
}

var $arrayify = function(args) {
  return Array.prototype.slice.call(args);
};

var $glueFunctions = function (obj) {
  for (var i in obj) {
    if (obj[i] instanceof Function) {
      obj[i] = obj[i].bind(obj);
    }
  }
};

var $mix = function (){
    var ret = {};
    if (Object.assign) {
      var args = Array.prototype.slice.call(arguments);
      return Object.assign.apply(Object, args);
    }

    for(var i=0; i<arguments.length; i++)
        for(var key in arguments[i])
            if(arguments[i].hasOwnProperty(key))
                ret[key] = arguments[i][key];
    return ret;
};

module.exports = {
    'Arrayify': $arrayify,
    'GlueFunctions': $glueFunctions,
    'Mix': $mix
};
