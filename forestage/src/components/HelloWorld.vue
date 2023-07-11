<script setup>
import { ref } from 'vue'
import { useCounterStore } from "../stores/hello";
import { storeToRefs } from "pinia";

defineProps({
  msg: String,
})

const counterStore = useCounterStore();
console.log(counterStore.count)
counterStore.increment()
console.log(counterStore.count)
const { count, name, doubleCount } = storeToRefs(counterStore);
</script>

<template>
  HelloWorld
  <h1>{{ msg }}</h1>

  <div class="card">
    <button type="button" @click="count++">count is {{ count }}</button>
    <p>
      store {{ counterStore.doubleCount }}
    </p>
  </div>
</template>

<style scoped>
.read-the-docs {
  color: #888;
}
</style>
