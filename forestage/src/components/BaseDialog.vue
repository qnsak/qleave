<script setup>
import { ref, computed } from 'vue'

const count = ref(0)
//const dialogShow = ref(false)
const props = defineProps(['title'])
const emit = defineEmits(["doSth"]);

const add = () => {
    emit('doSth', 23);
    count.value++
    console.log(count.value)
}

const childAdd = () => {
    console.log(122222)
};

const dialogShow = computed({
    get: () => {
        if (props.title == "123456") {
            return true
        } else {
            return false
        }
    }
})

const showDialog = ref(false);

const setShowDialog = (isShow) => {
    if (isShow) {
        showDialog.value = true;
    } else {
        showDialog.value = false;
    }
}

defineExpose({
    setShowDialog,
});
</script>

<template>
    <div
        class="bg-gray-100 bg-opacity-75 hs-overlay w-full h-full fixed top-0 left-0 z-[60] overflow-x-hidden overflow-y-auto"
        v-if="dialogShow"
    >
        <div class="bg-white sm:max-w-lg sm:w-full m-3 sm:mx-auto rounded-md border-2 border-slate-700">
            <div class="flex flex-col bg-white rounded-xl">
                <div class="flex justify-between items-center py-3 px-4 border-b-2 border-slate-700">
                    <h3 class="font-bold text-gray-800 dark:text-white">
                    Modal title
                    </h3>
                </div>
                <div class="p-4 overflow-y-auto">
                    <p class="mt-1 text-gray-800 dark:text-gray-400">
                    This is a wider card with supporting text below as a natural lead-in to additional content.
                    </p>
                </div>
                <div class="flex justify-end items-center gap-x-2 py-3 px-4 border-t-2 border-slate-700">
                    <button
                        class="h-10 w-1/6 border-2 border-slate-700"
                    >
                        取消
                    </button>
                    <button
                        class="h-10 w-1/6 border-2 border-slate-700"
                    >
                        確定
                    </button>
                </div>
            </div>
        </div>
    </div>

</template>