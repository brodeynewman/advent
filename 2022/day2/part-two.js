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

const oppRules = {
  // rock
  A: (y) => {
    // lose
    if (y === 'X') return 'Z';
    // draw
    if (y === 'Y') return 'X';
    // win
    return 'Y';
  },
  // paper
  B: (y) => {
    // lose
    if (y === 'X') return 'X';
    // draw
    if (y === 'Y') return 'Y';
    // win
    return 'Z';
  },
  // scissors
  C: (y) => {
    // lose
    if (y === 'X') return 'Y';
    // draw
    if (y === 'Y') return 'Z';
    // win
    return 'X';
  },
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

    const result = oppRules[x](y);
    const gameRes = rules[x](result)
    const roundPoints = resVals[gameRes] + pointVals[result];

    currPoints += roundPoints;
  }

  console.log('final', currPoints)
}

main()
