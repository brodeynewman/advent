const fs = require('fs');
const path = require('path');

const file = path.resolve(__dirname, 'input.txt');

const numRegex = /\d/g;

function buildStacks(split) {
  const stacks = new Map();

  let numRow;
  let numRowInd;

  // captures the number row in our stack
  for (let i = 0; i < split.length; i++) {
    const row = split[i];

    if (numRegex.test(row)) {
      numRow = split[i];
      numRowInd = i;
      break;
    }
  }

  let bufferAmt = 0;

  for (let i = 0; i < numRow.length; i++) {
    if (numRegex.test(numRow[i])) {
      const currNum = numRow[i];

      // create our initial stack
      if (!stacks.has(currNum)) {
        stacks.set(currNum, []);
      }

      // we now know where to start tracking stacked letters
      for (let j = 0; j < numRowInd; j++) {
        const target = split[j][bufferAmt];

        if (target && target !== ' ') {
          stacks.set(currNum, [...stacks.get(currNum), target]);
        }
      }
    }

    bufferAmt++;
  }

  return stacks;
}

const moveCrates = (split, stacks) => {
  for (let i = 0; i < split.length; i++) {
    if (split[i].includes('move')) {
      const [amt, x, y] = split[i].replace(/[^0-9 ]+/g, "").split(' ').filter(Boolean);

      const taken = stacks.get(x).slice(0, amt);

      stacks.get(x).slice(amt, stacks.get(x).length);
      stacks.get(x).splice(0, amt)
      stacks.get(y).unshift(...taken)
    }
  }

  stacks.forEach((value) => {
    console.log(value[0])
  })
}

function main() {
  const contents = fs.readFileSync(file);
  const res = Buffer.from(contents).toString();
  const split = res.split('\n');

  const stacks = buildStacks(split);
  moveCrates(split, stacks);
}

main()
