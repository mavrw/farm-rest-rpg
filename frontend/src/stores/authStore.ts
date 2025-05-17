import { defineStore } from "pinia";
import { computed, ref } from "vue";
import { login as apiLogin, logout as apiLogout, refresh as apiRefresh } from "@/api/auth";
import type { LoginPayload, AuthResponse } from "@/api/auth";

export const useAuthStore = defineStore('auth', () => {
    // State
    const accessToken = ref<string | null>(localStorage.getItem('access_token'))
    const isAuthenticated = computed(() => !!accessToken.value)
    
    // Helper Function
    const setAccessToken = (token: string) => {
        if(!token) {
            accessToken.value = null;
            localStorage.removeItem('access_token');
            return;
        }

        accessToken.value = token
        localStorage.setItem('access_token', token)
    }

    // Actions
    const login = async (payload: LoginPayload) => {
        const response = await apiLogin(payload)
        
        if(!response.access_token) {
            throw new Error('Access token missing from login response')
        }

        console.log('Login response: ', response)
        setAccessToken(response.access_token)
    }
    
    const logout = async () => {
        try {
            await apiLogout()
        } catch (err) {
            console.warn('Logout request failed, clearing local state anyway.')
        }
        
        accessToken.value = null
        localStorage.removeItem('access_token')
    }
    
    const refresh = async () => {
        const response = await apiRefresh()
        
        setAccessToken(response.access_token)
    }
    
    return {
        accessToken,
        isAuthenticated,
        login,
        logout,
        refresh,
    }
})
