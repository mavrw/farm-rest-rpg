import api from '@/api';
import type { InventoryItem, InventoryResponse } from '@/types/inventory';

export async function getItem(item_id: number): Promise<InventoryItem> {
    const res = await api.get(`/inventory/item/${item_id}`);
    if (res.status != 200) throw new Error('error fetching inventory item');

    return res.data;
};

export async function listInventoryItems(): Promise<InventoryResponse> {
    const res = await api.get(`/inventory/items/all`);
    if (res.status != 200) throw new Error('error fetching inventory');

    return res.data;
};
