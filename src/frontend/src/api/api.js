const apiEndpoint = ""; //leave empty for prod

import axios from "axios";

const getTasks = () => {
  return axios.get(apiEndpoint + "/tasks");
};

const getSchemas = () => {
  return axios.get(apiEndpoint + "/schemas");
};

const cleanParams = (params) => {
  let p = {};
  p.n = parseInt(params.n);
  p.theta = parseFloat(params.theta);
  p.tau = parseFloat(params.tau);
  p.task = params.task;
  return p;
};

const solve = (p) => {
  return axios.post(apiEndpoint + "/solve", p);
};

export { getSchemas, getTasks, cleanParams, solve };
