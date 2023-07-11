<script setup>
import { Form, Field } from 'vee-validate';
import * as Yup from 'yup';
import { ref } from 'vue'
import { useAuthStore } from '@/stores';
import Loading from '@/components/Laoding.vue'

const schema = Yup.object().shape({
    email: Yup.string().required('email is required'),
    password: Yup.string().required('Password is required')
});

const authStore = useAuthStore();

const loginResult = ref("")
async function onSubmit(values) {
    loadingRef.value.openLoading(true)
    const { email, password } = values;
    loginResult.value = await authStore.login(email, password);
    loadingRef.value.openLoading(false)
}

async function loginManageAccount(values) {
    loadingRef.value.openLoading(true)
    loginResult.value = await authStore.login("I0001@mail.com", "123456");
}

async function loginBasicAccount(values) {
    loadingRef.value.openLoading(true)
    loginResult.value = await authStore.login("I0002@mail.com", "123456");
}

const loginType = ref("demoLogin")
const loadingRef = ref(null);
</script>

<template>
    <Loading ref="loadingRef"/>
    <section class="container w-full mx-auto min-h-screen px-2 flex justify-center items-center flex-col gap-2">
        <div
            class="w-60 max-w-md rounded border-2 p-2"
        >
            <input
                class=""
                type="radio"
                name="loginType"
                id="standardLogin"
                value="standardLogin"
                v-model="loginType"
            />
            <label
                class=""
                for="standardLogin"
            >
                一般登入
            </label>
            <input
                class=""
                type="radio"
                name="loginType"
                id="demoLogin"
                value="demoLogin"
                v-model="loginType"
            />
            <label
                class=""
                for="demoLogin"
            >
                預設登入帳號
            </label>
        </div>
        <div>
            <div
                v-if="loginType == 'demoLogin'"
                class="w-60"
            >
                <div
                    class="w-full mx-auto p-8 rounded-lg border-2 max-w-md h-fit flex flex-col gap-3"
                >
                    <button 
                        class="w-full outline outline-slate-300 outline-offset-2 outline-1 rounded py-3 px-6 hover:bg-sky-500 ease-in"
                        @click="loginManageAccount"
                    >
                        管理
                    </button>
                    <button 
                        class="w-full outline outline-slate-300 outline-offset-2 outline-1 rounded py-3 px-6 hover:bg-sky-500 ease-in"
                        @click="loginBasicAccount"
                    >
                        一般
                    </button>
                </div>
            </div>
            <div
                v-else
            >
                <Form
                    @submit="onSubmit"
                    :validation-schema="schema" 
                    v-slot="{ errors }"
                >
                    <div class="flex flex-col mx-auto w-3/5 min-w-min max-w-md h-96 space-y-12 p-8 rounded-lg border-2">
                        <span class="text-center leading-tight tracking-tight text-gray-900">
                            登入
                            <div class="invalid-feedback">{{ loginResult }}</div>
                        </span>
                        <span class="flex dark:border bg-white border-b-2 border-b-slate-200">
                            <label class="flex-none flex-1 p-2" for="email">帳號</label>
                            <Field  
                                name="email"
                                type="text"
                                class="flex-1 p-2 text-center focus:outline-none focus:bg-zinc-100"
                                :class="{ 'is-invalid': errors.email }"
                                value="I0002@mail.com"
                            />
                            <div class="invalid-feedback">{{ errors.email }}</div>
                        </span>
                        <span class="flex bg-white  border-b-2 border-b-slate-200">
                            <label class="flex-none flex-1 p-2" for="password">密碼</label>
                            <Field
                                name="password"
                                type="password"
                                class="flex-1 p-2 text-center focus:outline-none focus:bg-zinc-100"
                                :class="{ 'is-invalid': errors.password }" 
                                value="123456"
                            />
                            <div class="invalid-feedback">{{ errors.password }}</div>
                        </span>
                        <span class="flex justify-center items-center">
                            <button
                                class="w-full outline outline-slate-300 outline-offset-2 outline-1 rounded p-2 m-1 hover:bg-sky-500 ease-in"
                                type="reset"
                            >
                                清空
                            </button>
                            <button 
                                class="w-full outline outline-slate-300 outline-offset-2 outline-1 rounded p-2 m-1 hover:bg-sky-500 ease-in"
                            >
                                確定
                            </button>
                        </span>
                    </div>
                </Form>
            </div>
        </div>
    </section>
</template>