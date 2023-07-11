<script>
import { defineComponent } from "vue";
import Navigation from "../components/Navigation.vue";
import Header from "../components/Header.vue";
export default defineComponent({
    name: "DashboardLayout",
    components: {
        Header,
        Navigation
    }
});
</script>

<template>
    <div class="h-full w-screen flex flex-col">
        <!-- Header section -->
        <header
            class="h-fit w-full basis-[5%]0 px-8"
        >
            <Header></Header>
        </header>

        <!-- Left menu -->
        <aside class="">
            <Navigation></Navigation>
        </aside>

        <!-- Main content -->
        <main class="float-left basis-[95%] w-full overflow-x-hidden">
            <slot name="default">
            </slot>
        </main>
        
        <!-- Footter -->
        <footer></footer>
    </div>
</template>