import api from '@/api';
import type { UserResponse } from '@/types/user';

export async function getMe(): Promise<UserResponse> {
    const res = await api.get('/users/me');
    if (res.status != 200) throw new Error('Failed to fetch current user');
    
    return res.data;
}

export async function updateMe(): Promise<UserResponse> {
    throw new Error('Not Implemented');
};

export async function deleteMe(): Promise<UserResponse> {
    throw new Error('Not Implemented');
};