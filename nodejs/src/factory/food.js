class BBQ {
  constructor(foodName) {
    this.name = foodName;
  }

  getName() {
    return this.name;
  }
}

class Soap {
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
  soap: Soap,
  drink: Drink
};

