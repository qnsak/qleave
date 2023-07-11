<script setup>
import { ref } from 'vue'

defineProps({
    msg: String,
})

// const counterStore = useCounterStore();
// const { count, name, doubleCount } = storeToRefs(counterStore);
const navIsDisplay = ref('false')

const setNavDisplay = () => {
    if (navIsDisplay.value) {
        navIsDisplay.value = false;
    } else {
        navIsDisplay.value = true;
    }
}
</script>

<template>
    <nav
        :class="{ navigationPosition: navIsDisplay }"
        class="navigation fixed top-11 left-0 z-[9999] flex flex-row"
    >
        <div
            :class="{ navigationHidden: navIsDisplay }"
            class="navigation w-min-h-[auto] min-w-[100px] z-20 flex shrink-0 grow-0 justify-around gap-4 bg-white border-2 border-slate-700 bg-white p-2.5 flex-col rounded-lg"
        >
            <router-link
                class="flex aspect-square min-h-[32px] w-18 flex-col items-center border-2 border-slate-700 justify-center gap-1 rounded-md p-3 text-indigo-600"
                to="/leave/apply"
            >
                <img src = "@/assets/icons/person_play_FILL0_wght100_GRAD200_opsz48.svg" title="假期管理"/>
                <small class="text-center text-[8px] font-medium"> 申請 </small>

            </router-link>
            <router-link
                class="flex aspect-square min-h-[32px] w-18 flex-col items-center border-2 border-slate-700 justify-center gap-1 rounded-md p-3 text-indigo-600"
                to="/leave/manage"
            >
                <img src = "@/assets/icons/person_play_FILL0_wght100_GRAD200_opsz48.svg" title="假期管理"/>
                <small class="text-center text-[8px] font-medium"> 管理 </small>

            </router-link>
            <router-link
                class="flex aspect-square min-h-[32px] w-18 flex-col items-center border-2 border-slate-700 justify-center gap-1 rounded-md p-3 text-indigo-600"
                to="/attendance"
            >
                <img src = "@/assets/icons/groups_FILL0_wght100_GRAD200_opsz48.svg" title="出缺席"/>
                <small class="text-center text-[8px] font-medium"> 出缺席 </small>

            </router-link>
            <router-link
                class="flex aspect-square min-h-[32px] w-18 flex-col items-center border-2 border-slate-700 justify-center gap-1 rounded-md p-3 text-indigo-600"
                to="/leave/approve"
            >
                <img src = "@/assets/icons/fact_check_FILL0_wght100_GRAD200_opsz48.svg" title="核假"/>
                <small class="text-center text-[8px] font-medium"> 核假 </small>

            </router-link>
            <hr class="dark:border-slate-700" />
            <router-link
                class="flex aspect-square min-h-[32px] w-18 flex-col items-center border-2 border-slate-700 justify-center gap-1 rounded-md p-3 text-indigo-600"
                to="/logout"
            >
                <img src = "@/assets/icons/logout_FILL0_wght100_GRAD200_opsz48.svg"/>
                <small class="text-center text-[8px] font-medium"> 登出 </small>
            </router-link>
        </div>
        <div>
            <button
                class="h-full w-8 bg-white border-2 border-slate-300 rounded-lg"
                @click="setNavDisplay"
            >
                <img
                    src = "@/assets/icons/chevron_right.svg"
                    v-if="navIsDisplay"
                />
                <img
                    src = "@/assets/icons/chevron_left.svg"
                    v-else
                />
            </button>
        </div>
    </nav>
</template>

<style scoped>
    .navigationPosition {
        left: -100px;
    }
    .navigationHidden {
        opacity: 0;
    }
    .navigation {
        transition: left 0.5s, opacity 0.5s linear;
    }
</style>
