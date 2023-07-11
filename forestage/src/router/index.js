import { createRouter, createWebHistory } from 'vue-router'

import { useAuthStore } from '@/stores';
import usersRoutes from './users.routes';
import leaveRoutes from './leave.routes';
import errorRoutes from './error.routes';

export const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    linkActiveClass: 'active',
    routes: [
        ...usersRoutes,
        ...leaveRoutes,
        ...errorRoutes,
    ]
});

router.beforeEach(async (to) => {
    const authStore = useAuthStore();
    const publicPages = ["/login", "/401", "/404"];
    const authRequired = !publicPages.includes(to.path);
    if (authRequired && !authStore.token) {
        authStore.returnUrl = to.fullPath;
        return '/login';
    }
});