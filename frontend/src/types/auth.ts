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
