import { defineStore } from "pinia";
import { computed, ref } from "vue";
import { login as apiLogin, logout as apiLogout, refresh as apiRefresh } from "@/api/auth";
import { getMe as apiGetMe } from "@/api/user";
import type { LoginPayload, AuthResponse } from "@/types/auth";
import type { UserResponse, User } from "@/types/user";
import { useUserStore } from "./userStore";

export const useAuthStore = defineStore('auth', () => {
    // State
    const accessToken = ref<string | null>(localStorage.getItem('access_token'));
    const isAuthenticated = computed(() => !!accessToken.value);
    // TODO: add `ready` flag to indicate store hydration

    // Helper Function
    const setAccessToken = (token: string) => {
        if(!token) {
            accessToken.value = null;
            localStorage.removeItem('access_token');
            return;
        }

        accessToken.value = token;
        localStorage.setItem('access_token', token);
    };

    // Actions
    const login = async (payload: LoginPayload) => {
        const response: AuthResponse = await apiLogin(payload);
        
        if(!response.access_token) {
            throw new Error('Access token missing from login response');
        }

        console.log('Login response: ', response);
        setAccessToken(response.access_token);
    };
    
    const logout = async () => {
        try {
            await apiLogout();
        } catch (err) {
            console.warn('Logout request failed, clearing local state anyway.');
        }
        
        accessToken.value = null;
        localStorage.removeItem('access_token');
        
        const userStore = useUserStore();
        userStore.clearUser();
    };
    
    const refresh = async () => {
        const response: AuthResponse = await apiRefresh();
        
        setAccessToken(response.access_token);
    };

    const fetchCurrentUser = async () => {
        try {
            const response: UserResponse = await apiGetMe();

            const userStore = useUserStore();
            const user: User = {
                id: response.id,
                username: response.username,
                email: response.email,
            }

            userStore.setUser(user);

        } catch {
            // invalidate authStore state via logout since 
            // access_token is no longer valid
            await logout();
        }
    };
    
    return {
        accessToken,
        isAuthenticated,
        login,
        logout,
        refresh,
        fetchCurrentUser,
    }
});
