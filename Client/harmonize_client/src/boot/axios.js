import Vue from "vue";
import axios from "axios";

import config from "../../config";

var axiosInstance = axios.create({
  baseURL: config.API_URL
});

axios.interceptors.request.use(request => {
  return request;
});

axios.interceptors.response.use(response => {
  return response;
});

Vue.prototype.$axios = axiosInstance;
