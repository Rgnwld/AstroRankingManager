import axios from "axios";

export const controller = axios.create({
    baseURL: '18.228.153.168:8080'
})