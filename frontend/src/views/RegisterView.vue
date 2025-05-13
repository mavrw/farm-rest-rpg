<template>
    <div class="register">
        <h1>Register</h1>
        <form @submit.prevent="handleRegister">
            <div>
                <label for="email">Email</label>
                <input v-model="email" type="email" id="email" required />
            </div>
            <div>
                <label for="username">Username</label>
                <input v-model="username" type="text" id="username" required />
            </div>
            <div>
                <label for="password">Password</label>
                <input v-model="password" type="password" id="password" required />
            </div>
            <button type="submit" :disabled="isSubmitting">Register</button>
        </form>
    </div>
</template>

<script setup lang="ts">

import { ref } from "vue";
import { useRouter } from "vue-router";
import { register } from "@/api/auth";
import type { RegisterPayload } from "@/api/auth";

const email = ref('');
const username =  ref('');
const password = ref('');
const isSubmitting = ref(false);

const router = useRouter();

const handleRegister = async () => {
    isSubmitting.value = true;

    const payload: RegisterPayload = {
        email: email.value,
        username: username.value,
        password: password.value,
    };

    try {
        await register(payload);
        router.push("/login");
    } catch (err) {
        console.error("Registration failed: ", err);
    } finally {
        isSubmitting.value = false;
    }
};  

</script>

<style scoped>

.register {
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
  background-color: #4CAF50;
  color: white;
  border: none;
  cursor: pointer;
}
button:disabled {
  background-color: #ccc;
}

</style>