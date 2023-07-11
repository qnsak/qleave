import { error401, error404 } from '@/views/error';

const route = [
    {
        path: '/401',
        component: error401,
        meta: {
            layout: 'AuthLayout',
        },
    },
    {
        path: '/404',
        component: error404,
        meta: {
            layout: 'AuthLayout',
        },
    },
]

export default route