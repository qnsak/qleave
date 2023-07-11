import { defineStore, getActivePinia  } from 'pinia';

import { router } from '@/router';
import api from "../api"

export const useAuthStore = defineStore({
    id: 'auth',
    state: () => ({
        user: JSON.parse(localStorage.getItem('user')),
        token: null
    }),
    actions: {
        async login(username, password) {
            try {
                let loginResult = await api.login.getUserInformation(username, password)

                if (loginResult.statusCode == 200) {
                    this.user = {
                        ID: loginResult.data.ID,
                        name: loginResult.data.name,
                        email: loginResult.data.email,
                        departmentID: loginResult.data.departmentID,
                        departmentTitle: loginResult.data.departmentTitle,
                        directorID: loginResult.data.directorID,
                    }
                    this.token = loginResult.data.accessToken
                    router.push('/leave/apply');
                    return null;
                } else {
                    return "登入失敗，請檢查帳號或者密碼是否正確。"
                }
            } catch (error) {            
                return "登入失敗，請檢查帳號或者密碼是否正確。"
            }
        },
        logout() {
            getActivePinia()._s.forEach(store => store.$reset());
            router.push('/login');
        }
    },
    persist: {
        enabled: true,
        strategies: [
            { storage: sessionStorage, paths: ['token', 'user'] },
        ],
    }
});