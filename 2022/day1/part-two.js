const fs = require('fs');
const path = require('path');

const file = path.resolve(__dirname, 'input.txt');

function main() {
  const contents = fs.readFileSync(file);
  const res = Buffer.from(contents).toString();

  // write code below this

  let highestCalories = [0];
  let currCalCount = 0;

  const split = res.split('\n');

  console.log('split', split.length)

  for (let i = 0; i < split.length; i++) {
    if (split[i] === '') {    
      const [lowestOfThree] = highestCalories;

      if (currCalCount > lowestOfThree) {
        if (highestCalories.length >= 3) {
          highestCalories.shift();
        }

        highestCalories.push(currCalCount);
        highestCalories = highestCalories.sort((a, b) => a - b);
      }

      currCalCount = 0;
    }

    currCalCount += Number(split[i]);
  }

  console.log('done', highestCalories.reduce((acc, curr) => acc + curr));
}

main()
