<script setup>
import { onMounted, ref, reactive } from 'vue'
import api from "@/api"

onMounted(() => {
    getLeaveList()
})

const leaveList = reactive([])

/**
 * 取得假期列表
 * 
 * @param {int} page - 第幾頁
 */
const getLeaveList = async (page = 0) => {
    let getLeaveListResult = await api.leave.getLeaveList(page)
    leaveList.length = 0
    Object.assign(leaveList, getLeaveListResult.data.data);
    let pageInfo = {
        'from': getLeaveListResult.data.meta.from,
        'to': getLeaveListResult.data.meta.to,
        'total': getLeaveListResult.data.meta.total,
        'prev': getLeaveListResult.data.links.prev,
        'next': getLeaveListResult.data.links.next,
    }
    setPagingInfo(pageInfo);
}

const pagingInfo = reactive({
    from: 0,
    to: 0,
    total: 0,
    prev: 0,
    next: 0,
})

/**
 * 設定分頁
 * 
 * @param {object} input - 分頁資料
 */
const setPagingInfo = (input) => {
    pagingInfo.from = input.from;
    pagingInfo.to = input.to;
    pagingInfo.total = input.total;
    pagingInfo.prev = input.prev;
    pagingInfo.next = input.next;
}

/**
 * 取消請假
 * 
 * @param {string} leaveId
 */
const cancelLeave =  async (leaveId) => {
    let cancelLeaveResult = await api.leave.cancelLeave(leaveId)

    if (cancelLeaveResult) {

    }

    getLeaveList()
};

/**
 * 取得假期是否能取消
 * 
 * @param {string} input - 時間
 * @return {string}
 */
const transTimeFormat = (input) => {
    let date = new Date(input);
    
    const year = date.getFullYear()
    const month = String(date.getMonth() + 1).padStart(2, '0')
    const day = String(date.getDate()).padStart(2, '0')
    const hours = String(date.getHours()).padStart(2, '0')
    const minute = String(date.getMinutes()).padStart(2, '0')

    return `${year}-${month}-${day} ${hours}:${minute}`
}

/**
 * 取得假期是否能取消
 * 
 * @param {object} leave - 假期資料
 * @return {bool}
 */
const canCancelLeave = (leave) => {
    if (!(leave.status in [0, 1])) {
        return false;
    }
    if (Date.now() > Date.parse(leave.startAt)) {
        return false;
    }
    return true;
}
</script>

<template>
    <section
        class="w-10/12 max-w-screen-xl container mx-auto pt-4 px-4"
    >
        <div
            class="flex justify-end gap-4"
        >
            <div
                class="flex items-center"
            >
                <span
                    class="w-fit material-icons mx-2 text-lg text-gray-600"
                >
                    change_circle
                </span>
                <p
                    class="w-10 inline-block text-[10px] text-left align-middle"
                >
                    等待中
                </p>
            </div>
            <div
                class="w-fit flex items-center"
            >
                <span
                    class="material-icons mx-2 text-lg text-green-400"
                >
                    check_circle_outline
                </span>
                <p
                    class="w-10 inline-block text-[10px] text-left align-middle"
                >
                    申請成功
                </p>
            </div>
            <div
                class="w-fit flex items-center"
            >
                <span
                    class="material-icons mx-2 text-lg text-red-400"
                >
                    highlight_off
                </span>
                <p
                    class="w-10 inline-block text-[10px] text-left align-middle"
                >
                    申請失敗
                </p>
            </div>
            <div
                class="w-fit border-l-2 border-gray-100 flex items-center"
            >
                <span
                    class="material-icons mx-2 text-lg text-red-400"
                >
                    delete
                </span>
                <p
                    class="w-10 inline-block text-[10px] text-left align-middle"
                >
                    取消
                </p>
            </div>
        </div>
        <div
            class="h-full w-full mx-auto"
        >
            <div class="overflow-x-auto bg-white">
                <table class="w-full">
                    <thead
                        class="bg-slate-600 border-b sticky top-0 z-50"
                    >
                        <tr>
                            <th
                                class="w-3/12 text-slate-100"
                            >
                                日期
                            </th>
                            <th
                                class="w-5/12 text-slate-100"
                            >
                                內容
                            </th>
                            <th
                                class="w-1/12 text-slate-100"
                            >
                                狀態
                            </th>
                            <th
                                class="w-3/12 text-slate-100"
                            >
                            </th>
                        </tr>
                    </thead>
                    <tbody
                        class="overflow-y-auto"
                    >
                        <tr
                            class="h-10 border-solid border-y border-gray-200 rounded-md p-0.5"
                            v-for="leave in leaveList"
                        >
                            <td
                                
                            >
                                <small class="text-left text-[10px]">
                                    開始：{{ transTimeFormat(leave.startAt) }}
                                </small>
                                <br>
                                <small class="text-left text-[10px]">
                                    結束：{{ transTimeFormat(leave.endAt) }}
                                </small>
                            </td>
                            <td>
                                <p class="text-left text-[12px]">
                                    {{ leave.reason }}
                                </p>
                            </td>
                            <td
                                class="text-center"
                            >
                                <span
                                    class="material-icons mx-2 text-lg text-gray-600"
                                    v-if="leave.status == 0"
                                >
                                    change_circle
                                </span>
                                <span
                                    class="material-icons mx-2 text-lg text-green-400"
                                    v-else-if="leave.status == 1"
                                >
                                    check_circle_outline
                                </span>
                                <span
                                    class="material-icons mx-2 text-lg text-red-400"
                                    v-else-if="leave.status == 2"
                                >
                                    highlight_off
                                </span>
                            </td>
                            <td>
                                <div
                                    class="h-full text-center flex justify-center items-center"
                                >
                                    <button
                                        class="h-full"
                                        @click.once="cancelLeave(leave.id)"
                                        v-if="canCancelLeave(leave)"
                                    >
                                        <span
                                            class="material-icons mx-2 text-lg text-red-400"
                                        >
                                            delete
                                        </span>
                                    </button>
                                </div>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
        <div class="flex flex-col items-center  mt-2 xs:mt-6">
            <span class="text-sm text-gray-700 dark:text-gray-400">
                目前為
                <span
                    class="font-semibold text-gray-900 dark:text-white"
                >
                    第{{ pagingInfo.from }} 件
                </span> 
                至 
                <span class="font-semibold text-gray-900 dark:text-white">
                    第{{ pagingInfo.to }}  件
                </span>
                    總數：
                <span class="font-semibold text-gray-900 dark:text-white">
                    {{ pagingInfo.total }}
                </span>
            </span>
            <div class="inline-flex mt-2 xs:mt-0">
                <!-- Buttons -->
                <button 
                    class="inline-flex items-center px-4 py-2 m-2 border-2 border-slate-700 text-sm font-medium"
                    @click="getLeaveList(pagingInfo.prev)"
                    v-if="pagingInfo.from > 1"
                >
                    <span
                        class="material-icons mx-2 text-xs"
                    >
                        arrow_back
                    </span>
                </button>
                <button
                    class="inline-flex items-center px-4 py-2 m-2 border-2 border-slate-700 text-sm font-medium"
                    @click="getLeaveList(pagingInfo.next)"  
                    v-if="pagingInfo.next !== 0"
                >   
                    <span
                        class="material-icons mx-2 text-xs"
                    >
                        arrow_forward
                    </span>
                </button>
            </div>
        </div>
    </section>
</template>
