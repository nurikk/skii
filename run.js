const fs = require('fs');

const get_position_value = (x, y, field) => {
  if (typeof field[x] !== 'undefined' && typeof field[x][y] !== 'undefined') {
    return field[x][y];
  }
  return -1;
}
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

const get_posible_directions = (x, y, field, start_value) => get_neibhors(x, y, field).filter(({elevation}) => elevation < start_value);


const find_next_steps = (start_x, start_y, field, current_path, done_cycle) => {
  let start_value = get_position_value(start_x, start_y, field);
  current_path.push(start_value);
  let possible_directions = get_posible_directions(start_x, start_y, field, start_value);
  possible_directions.forEach(({x: move_x, y: move_y, elevation: move_elevation}) => {
    let copy_path = current_path.slice(0);
    let next_moves = get_posible_directions(move_x, move_y, field, move_elevation);
    if (next_moves.length === 0) {
      copy_path.push(move_elevation);
      done_cycle.push(copy_path);
    } else {
      find_next_steps(move_x, move_y, field, copy_path, done_cycle);
    }
  });
}

const get_neibhors = (x, y, field) => {
  let neibhors = [];
  DIRECTIONS.forEach(({ dx, dy}) => {
    let xx = x + dx;
    let yy = y + dy;
    let elevation = get_position_value(xx, yy, field);
    if (elevation !== -1) {
      neibhors.push({
        x: xx,
        y: yy,
        elevation: elevation
      });
    }
  });
  return neibhors;
};

const find_peaks = (field) => {
  let peaks = [];
  for (let row_num = 0; row_num < field.length; row_num++) {
    for (let cell_num = 0; cell_num < field[0].length; cell_num++) {
      let current_elevation = get_position_value(row_num, cell_num, field);
      let neibhors = get_neibhors(row_num, cell_num, field);

      if (neibhors.every((neibhor) => neibhor.elevation < current_elevation)) {
        peaks.push({
          x: row_num,
          y: cell_num,
          elevation: current_elevation
        })
      }
    }
  }
  return peaks;
};


const solve = (field) => {
  const get_drop = (path) => path[0] - path[path.length - 1];
  let results = [];
  let peaks = find_peaks(field);

  peaks.forEach(({x, y}) => {
    find_next_steps(x, y, field, [], results);
  });

  const result = results.sort((a, b) => b.length - a.length || get_drop(b) - get_drop(a))[0];
  console.log('lenght', result.length);
  console.log('drop', get_drop(result));
  console.log('emal', `${result.length}${get_drop(result)}@redmart.com`);
};


const readFile = (file_name, cb) => {
  fs.readFile(file_name, 'utf8', (err, data) => {
    if (err) throw err;
    let game_field = [];
    let lines = data.split("\n");
    let [width, height] = lines[0].split(' ').map((n) => parseInt(n, 10));

    for (let line = 0; line < height; line++) {
      game_field[line] = lines[line + 1].split(' ').map((el) => parseInt(el)).slice(0, width);
    }
    cb(game_field);
  });
};

['./map.txt', './in.txt'].forEach((file_name) => readFile(file_name, solve));



