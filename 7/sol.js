const fs = require("fs");

const createDir = (tree, props, dir) => {
  if (!props.length) {
    tree[dir] = { size: 0 };
  } else {
    const last = props.pop();
    createDir(tree[last], props, dir);
  }
};

const createFile = (tree, props, file, size) => {
  if (!props.length) {
    tree[file] = size;
    tree["size"] += size;
  } else {
    const last = props.pop();
    createFile(tree[last], props, file, size);
    tree["size"] += size;
  }
};

const getTree = (data) => {
  const tree = {};
  var cur = [];
  const commands = data
    .split("$")
    .slice(1)
    .map((command) => command.substring(1));

  for (var command of commands) {
    command = command.split("\n");
    command.pop();
    if (command[0].substring(0, 2) === "cd") {
      const dir = command[0].substring(3);
      if (dir === "..") {
        cur.pop();
      } else {
        createDir(tree, [...cur].reverse(), dir);
        cur.push(dir);
      }
    } else {
      for (var line of command.slice(1)) {
        line = line.split(" ");
        if (line[0] !== "dir") {
          createFile(tree, [...cur].reverse(), line[1], parseInt(line[0]));
        }
      }
    }
  }
  delete tree.size;
  return tree;
};

const getScore1 = (tree, max) => {
  var score = tree.size > max ? 0 : tree.size;
  for (const prop in tree) {
    if (tree[prop].size) {
      score += getScore1(tree[prop], max);
    }
  }
  return score;
};

const getSizes = (tree, path) => {
  var acc = { [path.length ? path : "/"]: tree.size };
  for (const prop in tree) {
    if (tree[prop].size) {
      acc = { ...acc, ...getSizes(tree[prop], path + "/" + prop) };
    }
  }
  return acc;
};

fs.readFile("input", "utf8", (err, data) => {
  if (err) {
    console.error(err);
    return;
  }

  const tree = getTree(data);
  console.log(tree);
  console.log("Score 1:", getScore1(tree["/"], 100000));

  const total = 70000000;
  const available = total - tree["/"].size;
  const required = 30000000;
  const sizes = Object.entries(getSizes(tree["/"], ""));
  const dir = sizes
    .filter(([_, size]) => size + available >= required)
    .sort(([_0, size1], [_1, size2]) => size1 - size2)[0]
  console.log("You should delete", dir[0], "to free", dir[1], "of space.");
});
