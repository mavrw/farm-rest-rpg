import type { Plot } from "@/types/plot";
import { defineStore } from "pinia";
import { ref } from "vue";

export const usePlotStore = defineStore('plot', () => {
    // State
    const plots = ref<Map<number, Plot>>(new Map());
    const error = ref<string | null>(null);

    // Actions
    const fetchPlots = async () => {
        // TODO: fetch plots from backend using API
    };

    const plant = async (plotID: number, cropID: number) => {
        // TODO: plant plot from backend using API
    };

    const harvest = async (plotID: number) => {
        // TODO: harvest plot from backend using API
    };

    const getPlot = (id: number) => {
        return plots.value.get(id);
    };

    const getAllPlots = () => {
        return Array.from(plots.value.values());
    };

    const clearPlots = () => {
        plots.value.clear();
    };

    return {
        fetchPlots,
        plant,
        harvest,
        getPlot,
        getAllPlots,
        clearPlots,
        error,
    }
});