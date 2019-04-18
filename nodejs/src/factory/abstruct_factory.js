const factory = require('./factory');

function getFactory(factoryName) {
  if (!factory[factoryName]) {
    throw new Error(`Invaild factory name: '${factoryName}'`);
  }

  return factory[factoryName]
}

// getFactory('fdsa');
const factory1 = getFactory('DrinkFactory');
const food = factory1('juice');
console.log(food.constructor);
console.log(food.getName());

module.exports = {
  getFactory,
}
