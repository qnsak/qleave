import service from "./interceptors"

const httpObj = {
    get : async (url, params) => {
        try {
            const response = await service.get(url, params);
            return response;
        }
        catch (error) {
            return Promise.reject(error);
        }
    },
    post : async (url, params, config = null) => {
        try {
            const response = await service.post(url, params, config);
            return response;
        }
        catch (error) {
            return Promise.reject(error);
        }
    },
    put : async (url, params, config) => {
        try {
            const response = await service.put(url, params, config);
            return response;
        }
        catch (error) {
            return Promise.reject(error);
        }
    },
    delete : async (url, config) => {
        try {
            const response = await service.delete(url, config);
            return response;
        }
        catch (error) {
            return Promise.reject(error);
        }
    },
};

export default httpObj;