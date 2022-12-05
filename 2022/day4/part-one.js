const fs = require('fs');
const path = require('path');

const file = path.resolve(__dirname, 'input.txt');

const buildPairsAndCheckRedunancy = (p1, p2) => {
  let x = new Map();
  let y = new Map();

  let yContainsAll = true;
  let xContainsAll = true;

  const [x1, y1] = p1.split('-');
  const [x2, y2] = p2.split('-');

  for (let i = Number(x1); i <= Number(y1); i++) {
    x.set(i, 1);
  }

  for (let i = Number(x2); i <= Number(y2); i++) {
    y.set(i, 1);

    if (!x.has(i)) {
      yContainsAll = false;
    }
  }

  for (let id of x.keys()) {
    if (!y.has(id)) {
      xContainsAll = false;
    }
  }

  return xContainsAll || yContainsAll;
}

function main() {
  const contents = fs.readFileSync(file);
  const res = Buffer.from(contents).toString();
  const split = res.split('\n');

  let redundantPairs = 0;

  for (let i = 0; i < split.length; i++) {
    const pairs = split[i];
    const [p1, p2] = pairs.split(',');

    const isRedundant = buildPairsAndCheckRedunancy(p1, p2);
    
    if (isRedundant) redundantPairs++;
  }

  console.log('final', redundantPairs);
}

main()
