class BBQ {
  constructor(foodName) {
    this.name = foodName;
  }

  getName() {
    return this.name;
  }
}

class Stew {
  constructor(foodName) {
    this.name = foodName;
  }

  getName() {
    return this.name;
  }
}

class Drink {
  constructor(foodName) {
    this.name = foodName;
  }

  getName() {
    return this.name;
  }
}

module.exports = {
  bbq: BBQ,
  stew: Stew,
  drink: Drink
};

