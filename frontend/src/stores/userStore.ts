import { defineStore } from "pinia";
import { ref } from "vue";
import type { User } from "@/types/user";

export const useUserStore = defineStore('user', () => {
    // State
    const user = ref<User | null>(null);
    const error = ref<string | null>(null);

    // Actions
    const setUser = (data: User) => {
        user.value = data;
    };

    const clearUser = () => {
        user.value = null;
    };

    return {
        user,
        error,
        setUser,
        clearUser,
    }
});