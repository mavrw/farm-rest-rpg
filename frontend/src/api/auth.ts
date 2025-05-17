import api from '@/api'
import type { RegisterPayload, RegisterResponse, LoginPayload, AuthResponse } from '@/types/auth'

export async function register(payload: RegisterPayload): Promise<RegisterResponse> {
    const res = await api.post('/auth/register', payload, { skipAuthError: true })
    return res.data
}

export async function login(payload: LoginPayload): Promise<AuthResponse> {
    const res = await api.post('/auth/login', payload, { skipAuthError: true })
    return res.data
}

export async function refresh(): Promise<AuthResponse> {
    const res = await api.post('/auth/refresh')
    return res.data
}

export async function logout(): Promise<void> {
    await api.post('/auth/logout')
}