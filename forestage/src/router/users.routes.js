import { Login, Logout } from '@/views/users';

const route = [
    {
        path: '/users',
        component: Login,
        meta: {
            layout: 'AuthLayout',
        },
        children: [
            { path: '/login', component: Login },
        ]
    },
    {
        path: '/logout',
        component: Logout,
        meta: {
            layout: 'DashboardLayout',
        },
    },
]

export default route