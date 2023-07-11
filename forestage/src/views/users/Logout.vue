<script setup>
import { useAuthStore } from '@/stores';

const authStore = useAuthStore();
authStore.logout()
</script>

<template>
    <div class="container">
        <div class="row">
            登出
        </div>
    </div>
</template>