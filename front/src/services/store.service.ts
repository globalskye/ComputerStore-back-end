import axios, { Axios } from "axios";


const API_URL = "http://localhost:8080/";

export const getAllProductItems = () => {
  
  return axios.get(API_URL +"product/");
  
};
export const getAllCategories = () => {
    return axios.get(API_URL + "product/categories")
};