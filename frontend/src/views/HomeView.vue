<template>
    <div class="home">
        <h1>Welcome to FarmRestRPG</h1>
        <div v-if="isAuthenticated">
            <p>Glad to have you back!</p>
            <RouterLink to="/farm">Go to your farm</RouterLink>
        </div>
        <div v-else>
            <p>
                Please <RouterLink to="/login">log in</RouterLink> or <RouterLink to="/register">register</RouterLink>.
            </p>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useAuthStore } from "@/stores/authStore";
import { storeToRefs } from "pinia";

const authStore = useAuthStore();
const { isAuthenticated } = storeToRefs(authStore);
</script>

<style scoped>
.home {
    padding: 20px;
}
</style>