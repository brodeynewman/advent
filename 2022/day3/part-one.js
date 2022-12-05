const fs = require('fs');
const path = require('path');

const file = path.resolve(__dirname, 'input.txt');

const getUnicode = (letter) => {
  return letter.toLowerCase() === letter
    ? letter.toLowerCase().charCodeAt(0) - 97 + 1
    : letter.toLowerCase().charCodeAt(0) - 97 + (26 + 1);
}

const getIntersection = (a, b) => {
  for (let i = 0; i < a.length; i++) {
    if (b.includes(a[i])) return a[i];
  }
}

function main() {
  const contents = fs.readFileSync(file);
  const res = Buffer.from(contents).toString();
  const split = res.split('\n');

  // write code below this
  let currCount = 0;

  for (let i = 0; i < split.length; i++) {
    const curr = split[i];
    const len = curr.length;

    const a = curr.slice(0, len / 2);
    const b = curr.slice(len / 2, len);

    // find the intersection
    const intersection = getIntersection(a, b);

    const code = getUnicode(intersection);
    
    currCount += code;
  }

  console.log('final', currCount);
}

main()
