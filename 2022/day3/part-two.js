const fs = require('fs');
const path = require('path');

const file = path.resolve(__dirname, 'input.txt');

const getUnicode = (letter) => {
  return letter.toLowerCase() === letter
    ? letter.toLowerCase().charCodeAt(0) - 97 + 1
    : letter.toLowerCase().charCodeAt(0) - 97 + (26 + 1);
}

const getIntersection = (group) => {
  const [a, b, c] = group;

  for (let i = 0; i < a.length; i++) {
    let curr = a[i];

    if (b.includes(curr) && c.includes(curr)) {
      return curr;
    }
  }
}

function main() {
  const contents = fs.readFileSync(file);
  const res = Buffer.from(contents).toString();
  const split = res.split('\n');

  // write code below this
  let currCount = 0;
  let currGroup = [];
  let groupInd = 0;

  for (let i = 0; i <= split.length; i++) {
    const curr = split[i];

    if (i > 0 && i % 3 === 0) {
      // calculate intersection between group
      const intersection = getIntersection(currGroup);
      const code = getUnicode(intersection);

      currCount += code;

      currGroup = [];
    }

    currGroup.push(curr);
  }

  console.log('final', currCount);
}

main()
