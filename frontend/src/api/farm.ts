import api from "@/api";
import type { CreateFarmPayload, FarmResponse } from "@/types/farm";



export async function create(payload: CreateFarmPayload): Promise<FarmResponse> {
    const res = await api.post('/farm/create', payload);
    if (res.status != 201) throw new Error('Failed to create farm');
    
    return res.data;
}

export async function get(): Promise<FarmResponse> {
    const res = await api.get('/farm/get');
    if (res.status != 200) throw new Error('Failed to fetch farm');
    
    return res.data;
}