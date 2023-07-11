import request from "../../request/index";

const BASE_URL = '/login'

let api = {
    getUserInformation: async (email, password) => {
        try {
            let params = { 
                email: email,
                password: password,
            };
            const response =  await request.post(BASE_URL, params);
            return response.data;
        }
        catch (error) {
            return error
        }
    },
};
export default api;
