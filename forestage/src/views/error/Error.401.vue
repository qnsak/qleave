<script setup>
import { router } from '@/router';

const backLoginPage = () => {
    router.push('/login');
}

</script>
<template>
    <section
        class="container w-full mx-auto"
    >
        <div
            class="h-fit mt-12 flex flex-col grow-1 items-center"
        >
            <div
                class="p-4"
            >
                <h1
                    class="font-black text-8xl text-red-300"
                >
                    Error 401
                </h1>
                <hr>
                <p
                    class="m-2 text-left"
                >
                    登入逾時，請需要重新登入。
                </p>
            </div>
            <div
                class="h-fit"
            >
                <button
                    class="w-fit border-solid border-2 border-slate-900 p-4"
                    @click="backLoginPage"
                >
                    回到登入頁
                </button>
            </div>
        </div>
    </section>
</template>