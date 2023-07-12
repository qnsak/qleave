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
        setUserInfo(getUserInformation.data)
    }

    onMounted(() => {
        getUserInfo();
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

