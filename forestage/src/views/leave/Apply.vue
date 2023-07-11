<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { Form, Field, ErrorMessage } from 'vee-validate';
import * as Yup from 'yup';
import { useAuthStore } from "@/stores/auth.store";
import { storeToRefs } from "pinia";
import api from "@/api"

onMounted(() => {
    getApplyLeaveInfo()
})

const schema = Yup.object().shape({
    dateStar: Yup.string().required('開始日期不能為空'),
    dateEnd: Yup.string().required('結束日期不能為空'),
    timeStar: Yup.string().required('開始時間不能為空'),
    timeEnd: Yup.string().required('結束時間不能為空'),
    type: Yup.string().required('假期類型不能為空'),
    reason: Yup.string().required('請假理由為必填'),
});

const nationalHoliday = reactive({})
const leaveType = reactive([])

/*
 * 取得請假頁面資料
 */
const getApplyLeaveInfo = async () => {
    let leaveInfo = await api.leave.getApplyLeaveInfo()
    Object.assign(nationalHoliday, leaveInfo.data.nationalHoliday);
    Object.assign(leaveType, leaveInfo.data.leaveType);
}

/*
 * 取得 directorID
 */
const getDirectorID = () => {
    const authStore = useAuthStore();
    const { user } = storeToRefs(authStore);
    return user.value.directorID;
}

const formStatus = ref("edit");
const sendMessega = ref("");
const errorMessage = ref("");

/*
 * 送出請假資料
 */
const onSubmit = async (values) => {
    const { dateStar, timeStar, dateEnd, timeEnd, type, reason } = values;

    let dateTimeStar = formatDateTime(`${dateStar} ${timeStar}`);
    let dateTimeEnd = formatDateTime(`${dateEnd} ${timeEnd}`);
    
    if (Date.parse(dateTimeStar) > Date.parse(dateTimeEnd)) {
        errorMessage.value = "開始時間不能晚於結束時間。"
        return;
    } else {
        errorMessage.value = "";
    }

    formStatus.value = "send";
    sendMessega.value = "請等待";

    let directorID = getDirectorID();

    let leaveInfo = {
        "typeId": parseInt(type),
        "directorID": directorID,
        "reason": reason,
        "startAt": dateTimeStar,
        "endAt": dateTimeEnd,
    }

    let applyLeavetResult = await api.leave.applyLeavet(leaveInfo)

    if (applyLeavetResult.statusCode == 200) {
        sendMessega.value = "假單申請成功";
    } else {
        sendMessega.value = "假單申請失敗";
    }
}

/*
 * 轉換時間格式
 */
const formatDateTime = (input) => {
    let datetimeObject = new Date(input)

    let year = datetimeObject.getFullYear();
    let month = String(datetimeObject.getMonth() + 1).padStart(2, '0');
    let day = String(datetimeObject.getDate()).padStart(2, '0');
    let hour = String(datetimeObject.getHours()).padStart(2, '0');
    let minute = String(datetimeObject.getMinutes()).padStart(2, '0');
    let seconds = String(datetimeObject.getSeconds()).padStart(2, '0');

    return `${year}-${month}-${day}T${hour}:${minute}:${seconds}Z`;
}

/*
 * 重新整理頁面 
 */
const reloadPage = () => {
    location.reload();
}
</script>

<template>
    <section
        class="overflow-hidden px-8"
    >
        <div
            class="container mx-auto p-4 max-w-3xl"
        >
            <div>
                <p
                    class="text-left font-bold text-lg"
                >
                    國定假日
                </p>
            </div>
            <div
                class="flex flex-wrap gap-1 justify-items-center text-xs"
            >
                <div
                    class="inline-block w-fit border-2 border-green-400 px-2"
                    v-for="(item, index) in nationalHoliday"
                >
                    {{ item.title }}
                    {{ item.starTime }} ~ {{ item.endTime }}
                </div>
            </div>
        </div>
        <div
            class="container mx-auto p-4 max-w-3xl"
        >
            <div>
                <p>
                    {{ errorMessage }}
                </p>
            </div>
            <Form
                @submit="onSubmit"
                :validation-schema="schema" 
                v-slot="{ errors }"
            >
                <div class="bg-white w-full m-3 sm:mx-auto rounded-md border-2 border-slate-700">
                    <div class="flex flex-col bg-white rounded-xl">
                        <div class="flex justify-between items-center py-3 px-4 border-b-2 border-slate-700">
                            <h3 class="font-bold text-gray-800 dark:text-white">
                                請假
                            </h3>
                        </div>
                        <div 
                            class="p-4 overflow-y-auto"
                            v-if="formStatus == 'edit'"
                        >
                            <div
                                class="invalid-feedback text-xs"
                            >
                                <div
                                    class="w-fit float-left mx-2 text-rose-400"
                                    v-if="errors.dateStar"
                                >
                                    <span class="material-icons mx-2 text-xs">
                                        error
                                    </span>
                                    {{ errors.dateStar }}
                                </div>
                                <div
                                    class="w-fit float-left text-rose-400"
                                    v-if="errors.timeStar"
                                >
                                    <span class="material-icons text-xs">
                                        error
                                    </span>
                                    {{ errors.timeStar }}
                                </div>
                            </div>
                            <div class="w-full dark:border bg-white border-b-2 border-b-slate-200 flex">
                                <label
                                    class="flex-none flex-1 p-1"
                                    for="dateStar"
                                >
                                    開始日期
                                </label>
                                <Field  
                                    name="dateStar"
                                    type="date"
                                    class="flex-1 p-2 text-center focus:outline-none focus:bg-zinc-100"
                                    :class="{ 'is-invalid': errors.dateStar }"
                                    v-model="dateStar"
                                />
                                <label
                                    class="flex-none flex-1 p-1"
                                    for="dateStar"
                                >
                                    開始時間
                                </label>
                                <Field  
                                    name="timeStar"
                                    as="select"
                                    class="bg-white w-2/12 flex-none flex-1 p-1 text-center focus:outline-none"
                                    :class="{ 'is-invalid': errors.timeStar }"
                                >
                                    <option
                                        v-for="index in 24"
                                        :value="`${index}:00`"
                                    >
                                        {{ `${index}:00` }}
                                    </option>
                                </Field>
                            </div>
                            <div
                                class="invalid-feedback text-xs"
                            >
                                <div
                                    class="w-fit float-left mx-2 text-rose-400"
                                    v-if="errors.dateEnd"
                                >
                                    <span class="material-icons mx-2 text-xs">
                                        error
                                    </span>
                                    {{ errors.dateEnd }}
                                </div>
                                <div
                                    class="w-fit float-left text-rose-400"
                                    v-if="errors.timeEnd"
                                >
                                    <span class="material-icons text-xs">
                                        error
                                    </span>
                                    {{ errors.timeEnd }}
                                </div>
                            </div>
                            <div class="w-full dark:border bg-white border-b-2 border-b-slate-200 flex">
                                <label class="flex-none flex-1 p-2" for="dateEnd">結束日期</label>
                                <Field
                                    name="dateEnd"
                                    type="date"
                                    class=" bg-white flex-1 p-2 text-center focus:outline-none"
                                    :class="{ 'is-invalid': errors.dateEnd }"
                                    :min="dateStar"
                                />
                                <label
                                    class="flex-none flex-1 p-1"
                                    for="dateStar"
                                >
                                    結束時間
                                </label>
                                <Field  
                                    name="timeEnd"
                                    as="select"
                                    class=" bg-white w-2/12 flex-none flex-1 p-1 text-center focus:outline-none"
                                    :class="{ 'is-invalid': errors.timeEnd }"
                                >
                                    <option
                                        v-for="index in 24"
                                        :value="`${index}:00`"
                                    >
                                        {{ `${index}:00` }}
                                    </option>
                                </Field>
                            </div>
                            <div
                                class="invalid-feedback text-xs"
                            >
                                <div
                                    class="w-fit float-left mx-2 text-rose-400"
                                    v-if="errors.type"
                                >
                                    <span class="material-icons mx-2 text-xs">
                                        error
                                    </span>
                                    {{ errors.type }}
                                </div>
                            </div>
                            <div class="w-full flex dark:border bg-white border-b-2 border-b-slate-200">
                                <label class="flex-none flex-1 p-2" for="type">請假類型</label>
                                <Field  
                                    name="type"
                                    as="select"
                                    class=" bg-white flex-1 p-2 text-center focus:outline-none"
                                    :class="{ 'is-invalid': errors.type }"
                                >
                                    <option
                                        v-for="(item, index) in leaveType"
                                        :value="item.id"
                                    >
                                        {{ item.title }}
                                    </option>
                                </Field>
                            </div>
                            <div
                                class="invalid-feedback text-xs"
                            >
                                <div
                                    class="w-fit float-left mx-2 text-rose-400"
                                    v-if="errors.reason"
                                >
                                    <span class="material-icons mx-2 text-xs">
                                        error
                                    </span>
                                    {{ errors.reason }}
                                </div>
                            </div>
                            <div class="w-full flex bg-white  border-b-2 border-b-slate-200">
                                <label
                                    class="w-2/12 flex-none flex-1 p-2"
                                    for="reason"
                                >
                                    原因
                                </label>
                                <Field
                                    name="reason"
                                    as="textarea"
                                    class="w-10/12 float-right text-left focus:outline-none focus:bg-zinc-100"
                                    :class="{ 'is-invalid': errors.reason }"
                                    maxlength="100"
                                    rows="3" cols="100"
                                    style="resize: none;"
                                />
                            </div>
                        </div>
                        <div
                            class="p-4 overflow-y-auto"
                            v-if="formStatus == 'send'"
                        >
                            {{ sendMessega }}
                        </div>
                        <div class="flex justify-end items-center gap-x-2 py-3 px-4 border-t-2 border-slate-700">
                            <button
                                class="h-10 w-1/6 border-2 border-slate-700"
                                v-if="formStatus == 'edit'"
                            >
                                確定
                            </button>
                            <button
                                class="h-10 w-1/6 border-2 border-slate-700"
                                v-if="formStatus == 'send'"
                                @click="reloadPage"
                            >
                                確定
                            </button>
                        </div>
                    </div>
                </div>
            </Form>
        </div>
    </section>
</template>

<style scoped>

</style>
