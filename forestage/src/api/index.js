//api/index.js
import login from "./login/login"
import leave from "./leave/leave"
import user from "./user/user"
import attendance from "./attendance/attendance"

let api = {
    login: { ...login},
    leave: { ...leave},
    user: { ...user},
    attendance: {...attendance},
};

export default api;