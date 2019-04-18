const factory = require('./factory');

const food = factory.BBQFactory('BBQ');

console.log(food.constructor);
console.log(food.getName());

