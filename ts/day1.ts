import { createReadStream } from "fs";
const readline = require("readline");
const fileStream = createReadStream("../example.txt");
const rl = readline.createInterface({
  input: fileStream,
  crlfDelay: Infinity,
});
let sum = 0;
rl.on("line", function (line: string) {
  const numbers: number[] = [];
  let lineNum = 0;
  [...line].forEach((char) => {
    if (Number.isInteger(parseInt(char))) {
      numbers.push(parseInt(char));
      //console.log(char)
    }
  });
  if (numbers.length > 1) {
    sum = sum + parseInt(`${numbers[0]}${numbers[numbers.length - 1]}`);
    lineNum = parseInt(`${numbers[0]}${numbers[numbers.length - 1]}`);
  } else if (numbers.length == 1) {
    sum = sum + parseInt(`${numbers[0]}${numbers[0]}`);
    lineNum = parseInt(`${numbers[0]}${numbers[0]}`);
  }

  console.log(`Sum: ${sum}`);
  console.log(`Line from file: ${line} calib: ${lineNum}`);
});
