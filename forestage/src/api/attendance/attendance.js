import request from "../../request/index";

const BASE_URL = '/attendance'

let api = {
    getAttendanceList: async (departmentId) => {
        try {
            const url = `${BASE_URL}/${departmentId}`;
            const response =  await request.get(url);
            return response.data;
        }
        catch (error) {
            return error
        }
    },
};
export default api;
