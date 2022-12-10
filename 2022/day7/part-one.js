const fs = require('fs');
const path = require('path');

const file = path.resolve(__dirname, 'input.txt');

const buildGraph = (split) => {
  let graph = {
    size: null,
    name: '/',
    children: [],
  };

  let walk = [graph];

  for (let i = 0; i < split.length; i++) {
    const cmd = split[i];
    const parts = cmd.split(' ');

    if (parts[0] === '$' && parts[1] === 'cd') {
      const [_x, _y, loc] = parts;

      if (loc === '/') {
        // shift pointer back to the top of the graph;
        walk = [graph];
      } else if (loc === '..') {
        // walk back up one vertex in the graph
        walk.pop()
      } else {
        // location switches context in the current dir
        const curr = walk[walk.length - 1];
        const [_x, _y, loc] = parts;

        const found = curr.children.find(n => n.name === loc);

        if (!found) {
          throw new Error(`Directory: ${loc} does not exist in current vertex: ${curr.name}`)
        }

        walk.push(found);
      }
    }

    if (parts[0] === 'dir' && parts[1]) {
      const [_x, dirname] = parts;

      const curr = walk[walk.length - 1];

      const vertex = {
        name: dirname,
        size: null,
        children: []
      }

      curr.children.push(vertex);
    }

    if (!isNaN(parts[0])) {
      const [size, filename] = parts;

      const curr = walk[walk.length - 1];

      curr.children.push({
        name: filename,
        size: Number(size),
        children: null,
      })
    }
  }

  return graph;
}

const calculateSize = (graph) => {
  const neededSizes = [graph];

  const sizes = [];  

  while (neededSizes.length) {
    const toSize = neededSizes.pop();

    const walk = [];
    
    if (toSize.children) {
      neededSizes.push(...toSize.children);
      walk.push(...toSize.children);
    }

    let track = 0;

    while (walk.length) {
      const curr = walk.pop();

      if (curr.children && curr.children.length) {
        // console.log('CHILDREN', curr.name, curr.children);
        walk.push(...curr.children);
      }

      if (curr.size) {
        track += curr.size;
      }
    }

    sizes.push(track);
  }

  let totals = 0;

  sizes.forEach((size) => {
    if (size <= 100000) {
      totals += size;
    }
  });

  console.log(totals)
}

function main() {
  const contents = fs.readFileSync(file);
  const res = Buffer.from(contents).toString();
  const split = res.split('\n');

  const graph = buildGraph(split);
  calculateSize(graph);
}

main()
