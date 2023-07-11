<script setup>
import { ref } from 'vue'

const show = ref(false)
const openLoading = (input) => {
    show.value = input;
}

defineExpose({
    openLoading
})

</script>
<template>
    <section
        class="loading-mask"
        v-if="show"
    >
        <div>
            <h1 class="title">
                <span class="title-word">
                    L
                </span>
                <span class="title-word">
                    o
                </span>
                <span class="title-word">
                    a
                </span>
                <span class="title-word">
                    d
                </span>
                <span class="title-word">
                    i
                </span>
                <span class="title-word">
                    n
                </span>
                <span class="title-word">
                    g
                </span>
            </h1>
        </div>
    </section>
</template>
<style scoped>
.loading-mask {
    background-color: rgba(255, 255, 255, 1);
    position: fixed;
    width: 100%;
    height: 100%;
    z-index: 999;
    display: flex;
    align-items: center;
    justify-content: center;
}
.title-word {
    font-size: 32px;
    animation: color-animation 6s ease-in-out infinite;
    text-shadow: 1px 1px 40px #000000;
}

.title-word:nth-child(1) {
    --color-1: rgb(245, 245, 245);
    --color-2: rgb(255, 232, 164);
    --color-3: rgb(84, 78, 78);
    --color-4: rgb(145, 220, 247);
}

.title-word:nth-child(2) {
    --color-1: rgb(245, 245, 245);
    --color-2: rgb(254, 192, 185);
    --color-3: rgb(232, 207, 250);
    --color-4: rgb(162, 253, 129);
}

.title-word:nth-child(3) {
    --color-1: rgb(245, 245, 245);
    --color-2: rgb(115, 115, 115);
    --color-3: rgb(239, 140, 140);
    --color-4: rgb(163, 163, 163);
}

.title-word:nth-child(4) {
    --color-1: rgb(245, 245, 245);
    --color-2: rgb(251, 2, 2);
    --color-3: rgb(165, 161, 161);
    --color-4: rgb(245, 242, 195);
}

.title-word:nth-child(5) {
    --color-1: rgb(245, 245, 245);
    --color-2: rgb(191, 241, 242);
    --color-3: rgb(163, 163, 163);
    --color-4: rgb(115, 115, 115);
}

.title-word:nth-child(6) {
    --color-1: rgb(245, 245, 245);
    --color-2: rgb(163, 163, 163);
    --color-3: rgb(115, 115, 115);
    --color-4: rgb(210, 206, 243);
}

.title-word:nth-child(7) {
    --color-1: rgb(245, 245, 245);
    --color-2: rgb(146, 216, 248);
    --color-3: rgb(28, 26, 26);
    --color-4: rgb(115, 115, 115);
}

@keyframes color-animation {
    0%    {
        color: var(--color-1)
    }
    10%   {
        color: var(--color-4)
    }
    30%   {
        color: var(--color-4)
    }
    50%   {
        color: var(--color-3)
    }
    75%   {
        color: var(--color-2)
    }
    85%   {
        color: var(--color-2)
    }
    100%  {
        color: var(--color-1)
    }
}
</style>