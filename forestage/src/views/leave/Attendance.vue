<script setup>
import { onMounted, ref, reactive } from 'vue'
import { useAuthStore } from '@/stores';
import { storeToRefs } from "pinia";
import api from "@/api"

defineProps({
    msg: String,
})

onMounted(() => {
    getattendanceList()
})

const leaveList = reactive({})

/*
 * 取得請假日曆
 */
const getattendanceList = async () => {
    let departmentID = getDepartmentID();
    let attendanceList = await api.leave.getAttendanceByDay(departmentID)
    Object.assign(leaveList, attendanceList.data);
}

/*
 * 取得部門ID
 */
const getDepartmentID = () => {
    const authStore = useAuthStore();
    const { user } = storeToRefs(authStore);
    return user.value.departmentID;
}
</script>

<template>
    <section
        class="container mx-auto my-4 max-w-3xl p-8 overflow-x-hidden overflow-y-auto rounded-md"
    >
        <div
            class="p-2 text-lg text-center"
        >
            最近 10 天休假清單
        </div>
        <div
            class="w-full flex flex-col gap-2 justify-items-center text-xs"
        >
            <div
                class="bg-white inline-block w-full h-fit border-y-2 border-y-blue-100 py-4 flex flex-row gap-2"
                v-for="(leave, index) in leaveList"
            >
                <div
                    class="w-fit flex items-center justify-center m-1 p-2 rounded-md border-emerald-800/0.9"
                    
                >
                    <time
                        class="text-base text-amber-900 text-center font-black"
                    > 
                        {{index}}
                    </time>
                </div>
                <div
                    class="float-lef flex flex-wrap gap-2"
                >
                    <div
                        class="float-left w-fill h-fit max-h-[40px] m-2 p-2 border-2 rounded-md border-green-800 flex flex-col"
                        v-for="item in leave"
                    >
                        <span
                            class="text-orange-950 font-semibold"
                        >
                            {{ item.name }}
                        </span>
                    </div>
                </div>
            </div>
        </div>
    </section>
</template>
