<script setup>
    import { onMounted, reactive } from 'vue'
    import api from "@/api"

    const userInfo = reactive({
        name: '',
        email: '',
        department: '',
        paidLeaveTotal: '',
        paidLeaveUse: '',
    })

    const setUserInfo = (userData) => {
        userInfo.name = userData.name;
        userInfo.email = userData.email;
        userInfo.department = userData.departmentTitle;
        userInfo.paidLeaveTotal = userData.paidLeaveTotal;
        userInfo.paidLeaveUse = userData.paidLeaveUse;
    };

    const getUserInfo = async () => {
        let getUserInformation = await api.user.getUserInformation();
        console.log(getUserInformation.data)
        setUserInfo(getUserInformation.data)
    }

    const leavespanst = reactive([])
    const getLeavespanst = async () => {
        let getLeavespanstResult = await api.leave.getLeaveComingspanst()
        leavespanst.length = 0
        Object.assign(leavespanst, getLeavespanstResult.data);
    }

    const getDate = (input) => {
        let date = new Date(input);

        const year = date.getFullYear()
        const month = String(date.getMonth() + 1).padStart(2, '0')
        const day = String(date.getDate()).padStart(2, '0')
        const hours = String(date.getHours()).padStart(2, '0')
        const minute = String(date.getMinutes()).padStart(2, '0')

        return `${year}-${month}-${day} ${hours}:${minute}`
    }

    onMounted(() => {
        getUserInfo();
        getLeavespanst()
    })
    
</script>

<template>
    <div
        class="flex flex-row float-right max-h-5 text-[8px] text-xs align-middle text-center"
    >
        <span
            class="mx-2 flex"
        >
            <p
                class="mx-2 align-middle leading-loose"
            >
                部門：{{ userInfo.department }}
            </p>
        </span>
        <span
            class="mx-2 flex"
        >
            <p
                class="mx-2 align-middle leading-loose"
            >
                職員：{{ userInfo.name }}
            </p>
        </span>
    </div>
</template>

