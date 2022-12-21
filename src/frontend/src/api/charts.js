const getHalfYValues = (data, x, name) => {
  return data
    ? { y: data[data.length / 2 - 1], x: x, name: name, type: "scatter" }
    : { x: [], y: [] };
};

function transpose(matrix) {
  return matrix.reduce(
    (prev, next) => next.map((item, i) => (prev[i] || []).concat(next[i])),
    []
  );
}

const pairize = (arr1, arr2) => {
  return arr1.map((el, ind) => {
    return { y: el, x: arr2[ind] };
  });
};

const getHalfXValues = (data, x, name) => {
  return data
    ? { y: transpose(data)[data.length / 2 - 1], x: x, name: name }
    : { x: [], y: [] };
};

const createAxis = (size) => {
  let h = 1 / (size - 1);
  return [...Array(size).keys()].map((el) => el * h);
};

export { getHalfYValues, getHalfXValues, createAxis };
