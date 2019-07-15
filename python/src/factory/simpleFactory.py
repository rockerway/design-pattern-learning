class BBQ:
    def __init__(self, name):
        self.name = name

    def cook(self):
        return {
            'name': self.name,
            'type': 'BBQ'
        }

class Stew:
    def __init__(self, name):
        self.name = name

    def cook(self):
        return {
            'name': self.name,
            'type': 'Stew'
        }

class Drink:
    def __init__(self, name):
        self.name = name

    def cook(self):
        return {
            'name': self.name,
            'type': 'Drink'
        }

def factory(foodType):
    foodTypes = {
        'BBQ': BBQ,
        'Stew': Stew,
        'Drink': Drink
    }
    return foodTypes[foodType]

Food = factory('BBQ')
food = Food('bot')
print(food.cook())
