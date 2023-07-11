import { LeaveManage, LeaveApprove, LeaveApply, Attendance } from '@/views/leave';

const leaveRoutes =[
    {
        name: "leaveManage",
        path: '/leave/manage',
        component: LeaveManage, 
        meta: {
            layout: 'DashboardLayout',
        },
    },
    {
        name: "leaveApprove",
        path: "/leave/approve",
        meta: {
            layout: 'DashboardLayout',
        },
        component: LeaveApprove 
    },
    {
        name: "LeaveApply",
        path: "/leave/apply",
        meta: {
            layout: 'DashboardLayout',
        },
        component: LeaveApply 
    },
    {
        name: "attendance",
        path: '/attendance',
        meta: {
            layout: 'DashboardLayout',
        },
        component: Attendance,
    },
];

export default leaveRoutes