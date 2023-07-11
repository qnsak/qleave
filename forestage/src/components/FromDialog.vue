<script setup>
import { ref, computed } from 'vue'
import { Form, Field } from 'vee-validate';
import * as Yup from 'yup';
import { useAuthStore } from '@/stores';
import api from "@/api"

const count = ref(0)
//const dialogShow = ref(false)
const emit = defineEmits(["doSth"]);

// console.log(1)
// const add = () => {
//     emit('doSth', 23);
//     count.value++
//     console.log(count.value)
// }

const showDialog = ref(false);

const setShowDialog = (isShow) => {
    if (isShow) {
        showDialog.value = true;
    } else {
        showDialog.value = false;
    }
}



const schema = Yup.object().shape({
    dateStar: Yup.string().required('timeStar is required'),
    dateEnd: Yup.string().required('timeEnd is required'),
    timeStar: Yup.string().required('timeStar is required'),
    timeEnd: Yup.string().required('timeEnd is required'),
    type: Yup.string().required('type is required'),
    reason: Yup.string().required('reason is required'),
});

const authStore = useAuthStore();

const loginResult = ref("")

const formStatus = ref("edit")
const sendMessega = ref("")

async function onSubmit(values) {
    formStatus.value = "send";
    sendMessega.value = "請等待";
    const { dateStar, timeStar, dateEnd, timeEnd, type, reason } = values;

    let dateTimeStar = changeDateTimeFormat(`${dateStar} ${timeStar}`);
    let dateTimeEnd = changeDateTimeFormat(`${dateEnd} ${timeEnd}`);

    let leaveInfo = {
        "typeId": parseInt(type),
        "employeeID": 2,
        "directorID": 1,
        "reason": reason,
        "startAt": dateTimeStar,
        "endAt": dateTimeEnd,
    }

    let applyLeavetResult = await api.leave.applyLeavet(leaveInfo)

    if (applyLeavetResult.statusCode == 200) {
        sendMessega.value = "成功";
        emit('doSth');
    } else {
        sendMessega.value = "失敗";
    }
}

const changeDateTimeFormat = (input) => {
    let datetimeObject = new Date(input)

    let year = datetimeObject.getFullYear();
    let month = String(datetimeObject.getMonth() + 1).padStart(2, '0');
    let day = String(datetimeObject.getDate()).padStart(2, '0');
    let hour = String(datetimeObject.getHours()).padStart(2, '0');
    let minute = String(datetimeObject.getMinutes()).padStart(2, '0');
    let seconds = String(datetimeObject.getSeconds()).padStart(2, '0');

    return `${year}-${month}-${day}T${hour}:${minute}:${seconds}Z`;
}
const timeStar = ref()

defineExpose({
    setShowDialog,
});
</script>

<template>
    <div
        class="bg-gray-100 bg-opacity-75 hs-overlay w-full h-full fixed top-0 left-0 z-[60] overflow-x-hidden overflow-y-auto"
        v-if="showDialog"
    >
        <Form
            @submit="onSubmit"
            :validation-schema="schema" 
            v-slot="{ errors }"
        >
            <div class="bg-white sm:max-w-lg sm:w-full m-3 sm:mx-auto rounded-md border-2 border-slate-700">
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
                        <span class="flex dark:border bg-white border-b-2 border-b-slate-200">
                            <label class="flex-none flex-1 p-2" for="dateStar">開始時間</label>
                            <Field  
                                name="dateStar"
                                type="date"
                                class="flex-1 p-2 text-center focus:outline-none focus:bg-zinc-100"
                                :class="{ 'is-invalid': errors.dateStar }"
                                v-model="dateStar"
                            />
                            <div class="invalid-feedback">{{ errors.dateStar }}</div>
                            <Field  
                                name="timeStar"
                                as="select"
                                class="flex-1 p-2 text-center focus:outline-none focus:bg-zinc-100"
                                :class="{ 'is-invalid': errors.timeStar }"
                            >
                                <option
                                    v-for="index in 24"
                                    :value="`${index}:00`"
                                >
                                    {{ `${index}:00` }}
                                </option>
                            </Field>
                            <div class="invalid-feedback">{{ errors.timeStar }}</div>
                        </span>
                        <span class="flex bg-white  border-b-2 border-b-slate-200">
                            <label class="flex-none flex-1 p-2" for="dateEnd">結束時間</label>
                            <Field
                                name="dateEnd"
                                type="date"
                                class="flex-1 p-2 text-center focus:outline-none focus:bg-zinc-100"
                                :class="{ 'is-invalid': errors.dateEnd }"
                                :min="dateStar"
                            />
                            <div class="invalid-feedback">{{ errors.timeEnd }}</div>
                            <Field  
                                name="timeEnd"
                                as="select"
                                class="flex-1 p-2 text-center focus:outline-none focus:bg-zinc-100"
                                :class="{ 'is-invalid': errors.timeEnd }"
                            >
                                <option
                                    v-for="index in 24"
                                    :value="`${index}:00`"
                                >
                                    {{ `${index}:00` }}
                                </option>
                            </Field>
                            <div class="invalid-feedback">{{ errors.timeEnd }}</div>
                        </span>
                        <span class="flex dark:border bg-white border-b-2 border-b-slate-200">
                            <label class="flex-none flex-1 p-2" for="type">請假類型</label>
                            <Field  
                                name="type"
                                as="select"
                                class="flex-1 p-2 text-center focus:outline-none focus:bg-zinc-100"
                                :class="{ 'is-invalid': errors.type }"
                            >
                                <option value="1">Coffee</option>
                                <option value="2">Tea</option>
                                <option value="3">Coke</option>
                            </Field>
                            <div class="invalid-feedback">{{ errors.type }}</div>
                        </span>
                        <span class="flex bg-white  border-b-2 border-b-slate-200">
                            <label class="flex-none flex-1 p-2" for="reason">原因</label>
                            <Field
                                name="reason"
                                type="text"
                                class="flex-1 p-2 text-center focus:outline-none focus:bg-zinc-100"
                                :class="{ 'is-invalid': errors.reason }"
                                maxlength="20"
                            />
                            <div class="invalid-feedback">{{ errors.reason }}</div>
                        </span>
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
                            @click="setShowDialog(false)"
                        >
                            關閉
                        </button>
                        <button
                            class="h-10 w-1/6 border-2 border-slate-700"
                            v-if="formStatus == 'edit'"
                        >
                            確定
                        </button>
                    </div>
                </div>
            </div>
        </Form>
    </div>
</template>