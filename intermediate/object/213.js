class Dog {
  name;
  size;
  age;

  constructor(name, size, age) {
    this.name = name;
    this.size = size;
    this.age = age;
  }

  bark() {
    if (this.size >= 50) {
      return "Wooof! Woof!";
    }
    if (this.size >= 20) {
      return "Ruff! Ruff!";
    } else {
      return "Yip! Yip!";
    }
  }
  calcHumanAge() {
    return 12 + (this.age - 1) * 7;
  }
}

const goldenRetriever = new Dog("Golden Retriever", 60, 10);
console.log(goldenRetriever.bark());
console.log(goldenRetriever.calcHumanAge());

const siberianHusky = new Dog("Siberian Husky", 55, 6);
console.log(siberianHusky.bark());
console.log(siberianHusky.calcHumanAge());

const poodle = new Dog("poodle", 10, 1);
console.log(poodle.bark());
console.log(poodle.calcHumanAge());

const shibaInu = new Dog("shibaInu", 35, 4);
console.log(shibaInu.bark());
console.log(shibaInu.calcHumanAge());
