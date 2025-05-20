import api from '@/api';
import type { PlotResponse } from '@/types/plot';

export async function getPlot(plotID: number): Promise<PlotResponse> {
    const res = await api.get(`/plots/${plotID}`);
    if (res.status != 200) throw new Error('error fetching plot');

    return res.data;
}

export async function getAllPlots(farmID: number): Promise<PlotResponse[]> {
    const res = await api.get(`/farm/${farmID}/plots`);
    if (res.status != 200) throw new Error('error fetching plots');

    return res.data;
}

export async function buyPlot(farmID: number): Promise<PlotResponse> {
    const res = await api.post(`/farm/${farmID}/plots`);
    if (res.status != 200) throw new Error('error purchasing new plots');

    return res.data;
}

export async function plantPlot(plotID: number, cropID: number): Promise<PlotResponse> {
    const res = await api.post(`/plots/${plotID}/plant/${cropID}`);
    if (res.status != 200) throw new Error('error planting crop in plot');

    return res.data;
}

export async function harvestPlot(plotID: number): Promise<PlotResponse> {
    const res = await api.post(`/plots/${plotID}/harvest`);
    if (res.status != 200) throw new Error('error harvesting plot');

    return res.data;
}