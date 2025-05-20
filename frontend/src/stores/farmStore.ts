import { defineStore } from "pinia";
import { get as apiGetFarm, create as apiCreateFarm } from '@/api/farm';
import { computed, ref } from "vue";

import type { CreateFarmPayload, FarmResponse, Farm } from "@/types/farm";

export const useFarmStore = defineStore('farm', () => {
    // State
    const farm = ref<Farm | null>(null);
    const hasFarm = computed(() => !!farm.value);
    const error = ref<string | null>(null);

    // Actions
    const get = async () => {
        try {
            const response: FarmResponse = await apiGetFarm();
            
            console.log("Farm API response: ", response);

            farm.value = response;
            error.value = null;
        } catch (err) {
            farm.value = null;

            if (err instanceof Error) {
                error.value = err.message;
            } else {
                error.value = String(err);
            }
        }

    };

    const create = async (payload: CreateFarmPayload) => {
        try {
            const response: FarmResponse = await apiCreateFarm(payload);

            farm.value = response;
            error.value = null;
        } catch (err) {
            farm.value = null;

            if (err instanceof Error) {
                error.value = err.message;
            } else {
                error.value = String(err);
            }
        }
    };

    return {
        farm,
        hasFarm,
        error,
        get,
        create,
    }
});