const fs = require('fs');

const DIRECTIONS = [{
    // left
    dx: -1,
    dy: 0
  },
  {
    // bottom
    dx: 0,
    dy: -1
  },
  {
    // right
    dx: 1,
    dy: 0
  },
  {
    // top
    dx: 0,
    dy: 1
  }
];
const get_drop = (path) => path[0] - path[path.length - 1];
const find_next_steps = (point, current_path, result) => {
  current_path.push(point.elevation);

  point.neibhors.forEach((neibhor) => {
    let copy_path = current_path.slice(0);

    if (neibhor.neibhors.length === 0) {
      copy_path.push(neibhor.elevation);
      if (
        copy_path.length > result.winner.length ||
        (copy_path.length == result.winner.length && get_drop(copy_path) > get_drop(result.winner))
      ) {
        result.winner = copy_path;
      }
    } else {
      find_next_steps(neibhor, copy_path, result);
    }
  });
}

class Point {
  constructor(elevation) {
    this.elevation = elevation;
    this.neibhors = [];
    this.isPeak = false;
  }
}

const rebuildMap = (field) => {

  const maxRow = field.length;
  const maxCol = field[0].length;

  for (let rowNum = 0; rowNum < maxRow; rowNum++) {
    for (let cellNum = 0; cellNum < maxCol; cellNum++) {
      let isPeak = true;
      DIRECTIONS.forEach(({
        dx,
        dy
      }) => {
        let neighborRow = rowNum + dx;
        let neighborCell = cellNum + dy;
        if (neighborRow >= 0 && neighborRow < maxRow && neighborCell >= 0 && neighborCell < maxCol) {
          let neighbor = field[neighborRow][neighborCell]
          if (field[rowNum][cellNum].elevation <= neighbor.elevation) {
            isPeak = false
          } else {
            field[rowNum][cellNum].neibhors.push(neighbor);
          }
        }
      });
      field[rowNum][cellNum].isPeak = isPeak;
    }
  }
  return field;
};
const solve = (field) => {
  field = rebuildMap(field);
  let result = {
    winner: []
  };


  field.forEach((row) => {
    row.filter((point) => point.isPeak).forEach((point) => {
      find_next_steps(point, [], result);
    });
  });
  console.log('lenght', result.winner.length);
  console.log('drop', get_drop(result.winner));
  console.log('emal', `${result.winner.length}${get_drop(result.winner)}@redmart.com`);
};


const readFile = (file_name, cb) => {
  fs.readFile(file_name, 'utf8', (err, data) => {
    if (err) throw err;
    let game_field = [];
    let lines = data.split("\n");
    let [width, height] = lines[0].split(' ').map((n) => parseInt(n, 10));

    for (let line = 0; line < height; line++) {
      game_field[line] = lines[line + 1].split(' ').map((el) => new Point(parseInt(el))).slice(0, width);
    }
    cb(game_field);
  });
};

['./map.txt'].forEach((file_name) => readFile(file_name, solve));