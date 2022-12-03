const fs = require('fs');
const path = require('path');

const file = path.resolve(__dirname, 'input.txt');

const pointVals = {
  // rock
  X: 1,
  // paper
  Y: 2,
  // scissors
  Z: 3,
}

const resVals = {
  W: 6,
  L: 0,
  D: 3,
}

const rules = {
  // rock
  A: (y) => {
    if (y === 'Y') return 'W';
    if (y === 'X') return 'D';
    return 'L';
  },
  // paper
  B: (y) => {
    if (y === 'Z') return 'W';
    if (y === 'Y') return 'D';
    return 'L';
  },
  // scissors
  C: (y) => {
    if (y === 'X') return 'W';
    if (y === 'Z') return 'D';
    return 'L';
  },
}

function main() {
  const contents = fs.readFileSync(file);
  const res = Buffer.from(contents).toString();
  const split = res.split('\n');

  // write code below this
  let currPoints = 0;

  for (let i = 0; i < split.length; i++) {
    const [x, y] = split[i].split(' ');

    const result = rules[x](y);
    const roundPoints = resVals[result] + pointVals[y];

    currPoints += roundPoints;
  }

  console.log('final', currPoints)
}

main()
