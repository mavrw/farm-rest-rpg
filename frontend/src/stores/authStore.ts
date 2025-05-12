import { defineStore } from "pinia";
import { ref } from "vue";
import { login as apiLogin, logout as apiLogout, refresh as apiRefresh } from "@/api/auth";
import type { LoginPayload, AuthResponse } from "@/api/auth";

export const useAuthStore = defineStore('auth', () => {
    // State
    const accessToken = ref<string | null>(localStorage.getItem('access_token'))
    const isAuthenticated = ref(!!accessToken.value)
    
    // Helper Function
    const setAccessToken = (token: string) => {
        accessToken.value = token
        localStorage.setItem('access_token', token)
    }

    // Actions
    const login = async (payload: LoginPayload) => {
        const response = await apiLogin(payload)
        
        setAccessToken(response.access_token)
        isAuthenticated.value = true
    }
    
    const logout = async () => {
        try {
            await apiLogout()
        } catch (err) {
            console.warn('Logout request failed, clearing local state anyway.')
        }
        
        accessToken.value = null
        isAuthenticated.value = false
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
