import router from '@/router';
import { useAuthStore } from '@/stores/authStore';
import axios from 'axios';

const api = axios.create({
    baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api/v1',
    withCredentials: true,
});


api.interceptors.request.use(
    (config) => {
        const authStore = useAuthStore();
        const token = authStore.accessToken;

        if (token) {
            config.headers.Authorization = `Bearer ${token}`
        }
        return config
    },
    (error) => Promise.reject(error)
);

api.interceptors.response.use(
    response => response,
    async (error) => {
        const status = error.response?.status;
        const config = error.config;

        if(status === 401 && !config?.skipAuthError) {
            const authStore = useAuthStore();

            await authStore.logout();

            router.push("/");
        }

        return Promise.reject(error);
    }
);

export default api;