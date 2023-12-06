import axios from "axios";

const DB_ADDRESS = process.env.DB_ADDRESS
const DB_PORT = process.env.DB_PORT

export const controller = axios.create({
    baseURL: `${DB_ADDRESS}:${DB_PORT}`
})