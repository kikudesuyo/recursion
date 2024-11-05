class Rgb {
  red: number;
  green: number;
  blue: number;

  constructor(red: number, green: number, blue: number) {
    this.red = red;
    this.green = green;
    this.blue = blue;
  }

  getHexCode() {
    return `#${this.red.toString(16).padStart(2, "0")}${this.green
      .toString(16)
      .padStart(2, "0")}${this.blue.toString(16).padStart(2, "0")}`;
  }

  getBits() {
    const hexCode = this.getHexCode();
    return hexCode
      .split("")
      .slice(1)
      .map((char) => parseInt(char, 16).toString(2).padStart(4, "0"))
      .join("")
      .replace(/^0+/, "");
  }

  getColorShade() {
    if (this.red === this.green && this.green === this.blue) {
      return "grayscale";
    }
    const maxColor = Math.max(this.red, this.green, this.blue);
    if (maxColor === this.red) {
      return "red";
    } else if (maxColor === this.green) {
      return "green";
    } else {
      return "blue";
    }
  }
}

const color1 = new Rgb(0, 153, 255);
console.log(color1.getHexCode()); // #0099ff
console.log(color1.getBits()); // 000000001001100011111111
console.log(color1.getColorShade()); // blue

const color2 = new Rgb(255, 255, 204);
console.log(color2.getHexCode()); // #ffffcc
console.log(color2.getBits()); // 111111111111111100110011
console.log(color2.getColorShade()); // red

const color3 = new Rgb(0, 87, 0);
console.log(color3.getHexCode()); // #005700
console.log(color3.getBits()); // 000000000101011100000000
console.log(color3.getColorShade()); // green

const gray = new Rgb(123, 123, 123);
console.log(gray.getHexCode()); // #7b7b7b
console.log(gray.getBits()); // 011110110111101101111011
console.log(gray.getColorShade()); // grayscale
