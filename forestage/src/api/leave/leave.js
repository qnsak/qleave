import request from "../../request/index";

const BASE_URL = '/leave'

let api = {
    getLeaveComingList: async () => {
        try {
            const url = `${BASE_URL}/coming`

            const response =  await request.get(url, {});
            return response.data;
        }
        catch (error) {
            return error
        }
    },
    getLeaveList: async (page) => {
        try {
            const url = `${BASE_URL}?page=${page}`

            const response =  await request.get(url, {});
            return response.data;
        }
        catch (error) {
            return error
        }
    },
    getApproveLeaveList: async (page) => {
        try {
            const url = `${BASE_URL}/approve?page=${page}`
            const response = await request.get(url, {});
            return response.data;
        }
        catch (error) {
            return error
        }
    },
    applyLeavet: async (from) => {
        try {
            const url = BASE_URL;
            const response =  await request.post(url, from);
            return response.data;
        }
        catch (error) {
            return error
        }
    },
    cancelLeave: async (leaveId) => {
        try {
            const url = `${BASE_URL}/${leaveId}`
            const response =  await request.delete(url, {});
            
            if (response.data.statusCode == 200) {
                return response.data;
            } else {
                return "失敗";
            }
        }
        catch (error) {
            return error
        }
    },
    approveLeave: async (from) => {
        try {
            const url = BASE_URL;
            const response =  await request.put(url, from);
            return response.data;
        }
        catch (error) {
            return error
        }
    },
    getAttendanceByDay: async (departmentId) => {
        try {
            const url = `/attendance/${departmentId}`

            const response =  await request.get(url, {});
            return response.data;
        }
        catch (error) {
            return error
        }
    },
    getApplyLeaveInfo: async () => {
        try {
            const url = `${BASE_URL}/apply`
            const response =  await request.get(url);
            return response.data;
        }
        catch (error) {
            return error
        }
    },
};
export default api;
