const Food = require('./food');

function BBQFactory(foodName) {
  return new Food.bbq(foodName);
}

function StewFactory(foodName) {
  return new Food.stew(foodName);
}

function DrinkFactory(foodName) {
  return new Food.drink(foodName);
}

module.exports = {
  BBQFactory,
  StewFactory,
  DrinkFactory,
};

