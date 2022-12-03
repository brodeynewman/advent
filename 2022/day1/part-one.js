const fs = require('fs');
const path = require('path');

const file = path.resolve(__dirname, 'input.txt');

function main() {
  const contents = fs.readFileSync(file);
  const res = Buffer.from(contents).toString();

  // write code below this

  let highestIndex = 0;
  let currHighestCal = 0;
  let currCalCount = 0;

  const split = res.split('\n');

  console.log('split', split.length)

  for (let i = 0; i < split.length; i++) {
    if (split[i] === '') {      
      if (currCalCount > currHighestCal) {
        highestIndex = i;
        currHighestCal = currCalCount;
      }

      currCalCount = 0;
    }

    currCalCount += Number(split[i]);
  }

  console.log('done', highestIndex, currHighestCal);
}

main()
