const fs = require('fs');
const path = require('path');

const file = path.resolve(__dirname, 'input.txt');

const checkSlideUniqueness = (slider) => {
  if (slider.length !== 14) return false;

  const map = {};

  for (let i = 0; i < slider.length; i++) {
    map[slider[i]] = 1;
  }

  return Object.keys(map).length === 14;
}

function main() {
  const contents = fs.readFileSync(file);
  const res = Buffer.from(contents).toString();
  const split = res.split('\n');

  for (let i = 0; i < split.length; i++) {
    let ind = 0;

    const code = split[i];

    for (let j = 0; j < code.length; j++) {
      const slider = code.slice(j, j + 14)

      if (slider.length === 14) {
        const isUnique = checkSlideUniqueness(slider)

        if (isUnique) {
          ind = ind + 14;
          break;
        }
      }

      ind++;
    }

    console.log('slider position', ind)
  }
}

main()
