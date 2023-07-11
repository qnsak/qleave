import axios from "axios";

const domain = "/api";

const service = axios.create({
    baseURL: `${domain}`,
    headers: {
        "Content-Type": "application/json",
        "accept": "application/json",
    },
});

export default service