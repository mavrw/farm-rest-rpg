import api from '@/api'

export interface RegisterPayload {
    email: string,
    username: string,
    password: string
}

export interface LoginPayload {
    email: string,
    password: string
}

export interface RegisterResponse {
    message: string
}

export interface AuthResponse {
    access_token: string,
    refresh_token?: string
}

export async function register(payload: RegisterPayload): Promise<RegisterResponse> {
    const res = await api.post('/auth/register', payload)
    return res.data
}

export async function login(payload: LoginPayload): Promise<AuthResponse> {
    const res = await api.post('/auth/login', payload)
    return res.data
}

export async function refresh(): Promise<AuthResponse> {
    const res = await api.post('/auth/refresh')
    return res.data
}

export async function logout(): Promise<void> {
    await api.post('/auth/logout')
}