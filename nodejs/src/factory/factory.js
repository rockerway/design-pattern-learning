const Food = require('./food');

function BBQFactory(foodName) {
  return new Food.bbq(foodName);
}

function SoapFactory(foodName) {
  return new Food.soap(foodName);
}

function DrinkFactory(foodName) {
  return new Food.drink(foodName);
}

module.exports = {
  BBQFactory,
  SoapFactory,
  DrinkFactory,
};

