<script setup>
import { ref } from 'vue'
import { useCounterStore } from "../stores/hello";
import { storeToRefs } from "pinia";
import api from "../api"

defineProps({
  msg: String,
})

const counterStore = useCounterStore();
const { count, name, doubleCount } = storeToRefs(counterStore);

</script>

<template>
    HelloPinia
    store {{ counterStore.count }}
</template>

<style scoped>

</style>
