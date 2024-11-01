class Animal {
  name;
  weighKg;
  heightM;
  isPredator;
  sppedKph;
  activityMultiplier;

  constructor(name, weighKg, heightM, isPredator, sppedKph) {
    this.name = name;
    this.weighKg = weighKg;
    this.heightM = heightM;
    this.isPredator = isPredator;
    this.sppedKph = sppedKph;
    this.activityMultiplier = 1.2;
  }

  getBmi() {
    const bmi = this.weighKg / this.heightM ** 2;
    return Math.floor(bmi * 100) / 100;
  }
  getDailyCalories() {
    const calories =
      70 * Math.pow(this.weighKg, 0.75) * this.activityMultiplier;
    return Math.floor(calories * 100) / 100;
  }

  isDangerous() {
    if (this.isPredator) {
      return true;
    }
    if (this.heightM >= 1.7) {
      return true;
    }
    if (this.sppedKph >= 35) {
      return true;
    }
    return false;
  }
}

const rabbit = new Animal("rabbit", 10, 0.3, false, 20);
console.log(rabbit.getBmi());
console.log(rabbit.isDangerous());

const snake = new Animal("snake", 30, 1, true, 30);
console.log(snake.isDangerous());
console.log(snake.getDailyCalories());

const elephant = new Animal("elephant", 300, 3, false, 5);
console.log(elephant.getBmi());
console.log(elephant.getDailyCalories());

const gazelle = new Animal("gazelle", 50, 1.5, false, 100);
console.log(gazelle.getDailyCalories());
console.log(gazelle.isDangerous());
