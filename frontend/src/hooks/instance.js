import axios from "axios";
import ENV from "../.env";

const instance = (token) =>
  axios.create({
    baseURL: `${ENV.API_URL}`,
    timeout: 1000,
    headers: { authorization: "Bearer " + token },
  });

export default instance;
