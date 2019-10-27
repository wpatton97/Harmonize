import Vue from "vue";
import axios from "axios";

var axiosInstance = axios.create();

axios.interceptors.request.use(request => {
  return request;
});

axios.interceptors.response.use(response => {
  return response;
});

Vue.prototype.$axios = axiosInstance;
