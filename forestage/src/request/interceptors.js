import service from "./services";
import { useAuthStore } from "@/stores/auth.store";
import { storeToRefs } from "pinia";
import { router } from '@/router';

service.interceptors.request.use(
    async (config) => {
        //给请求头设置token
        const authStore = useAuthStore();
        const { token } = storeToRefs(authStore);
        if (token.value) {
            config.headers.token = `${token.value}`;
        }
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

// responce 攔截
service.interceptors.response.use(
    (response) => {
        return response;
    },
    (error) => {
        // HTTP 状态码
        const status = error.response.status;

        if (status === 401) {
            router.push('/401');
        }

        return Promise.reject(error.response);
    }
);

export default service