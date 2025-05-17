<template>
    <div class="login">
        <h1>Login</h1>
        <form @submit.prevent="handleLogin">
            <div>
                <label for="email">Email</label>
                <input v-model="email" type="email" id="email" required />
            </div>
            <div>
                <label for="password">Password</label>
                <input v-model="password" type="password" id="password" required />
            </div>
            <button type="submit" :disabled="isSubmitting">Login</button>
        </form>
    </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import { useAuthStore } from "@/stores/authStore";
import type { LoginPayload } from "@/types/auth";

const email = ref('');
const password = ref('');
const isSubmitting = ref(false);

const router = useRouter();
const authStore = useAuthStore();

const handleLogin = async () => {
    isSubmitting.value = true;

    const payload: LoginPayload = {
        email: email.value,
        password: password.value,
    };

    try {
        await authStore.login(payload);
        router.push("/");
    } catch (err) {
        console.error("Login failed: ", err);
    } finally {
        isSubmitting.value = false;
    }
};
</script>

<style scoped>
.login {
  max-width: 400px;
  margin: 0 auto;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 8px;
}
form {
  display: flex;
  flex-direction: column;
}
input {
  margin: 10px 0;
  padding: 8px;
  font-size: 14px;
}
button {
  padding: 10px;
  background-color: #2196F3;
  color: white;
  border: none;
  cursor: pointer;
}
button:disabled {
  background-color: #ccc;
}
</style>