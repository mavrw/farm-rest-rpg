import { defineStore } from "pinia";
import { ref } from "vue";
import type { InventoryItem } from "@/types/inventory";
import { getItem as apiGetItem, listInventoryItems as apiListInventoryItems } from "@/api/inventory";

export const useInventoryStore = defineStore('inventory', () => {
    const items = ref<InventoryItem[]>([]);
    const error = ref<string | null>(null);

    const fetchAll = async () => {
        try {
            items.value = await apiListInventoryItems();
            error.value = null;
        } catch (err: unknown) {
            error.value = err instanceof Error ? err.message : String(err);
        }
    };

    const fetchOne = async (id: number) => {
        try {
            const item = await apiGetItem(id);
            const idx = items.value.findIndex(i => i.ID === item.ID);
            
            if (idx >= 0) items.value[idx] = item
            else items.value.push(item)
            
            error.value = null
        } catch (err: unknown) {
            error.value = err instanceof Error ? err.message : String(err);
        }
    };

    const clear = () => {
        items.value = [];
        error.value = null;
    };

    return {
        items,
        error,
        fetchAll,
        fetchOne,
        clear
    }
});