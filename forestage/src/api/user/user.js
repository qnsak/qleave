import request from "../../request/index";

const BASE_URL = '/user'

let api = {
    getUserInformation: async () => {
        try {
            const url = `${BASE_URL}`
            const response =  await request.get(url, {});
            return response.data;
        }
        catch (error) {
            return error
        }
    },
};
export default api;
