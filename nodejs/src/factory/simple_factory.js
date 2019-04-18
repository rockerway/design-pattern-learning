const Food = require('./food');

function getFood(foodType, foodName) {
  if (!Food[foodType]) {
    throw new Error(`Invaild food type: ${foodType}`);
  }

  return new Food[foodType](foodName);
}

const food = getFood('bbq', 'BBQ');
console.log(food.constructor);
console.log(food.getName());

