import { defineStore } from "pinia";
import { ref } from "vue";
import { 
    getAllPlots as apiGetAllPlots,
    getPlot as apiGetPlot,
    buyPlot as apiBuyPlot,
    plantPlot as apiPlantPlot,
    harvestPlot as apiHarvestPlot,
} from "@/api/plot";
import type { Plot } from "@/types/plot";

export const usePlotStore = defineStore('plot', () => {
    // State
    const plots = ref<Map<number, Plot>>(new Map());
    const error = ref<string | null>(null);

    // Actions
    const fetchPlots = async (farmID: number) => {
        try {
            const response: Plot[] = await apiGetAllPlots(farmID);
            plots.value.clear();
            response.forEach(plot => plots.value.set(plot.id, plot));
        } catch (err) {
            if (err instanceof Error) {
                error.value = err.message;
            } else {
                error.value = String(err);
            }
        }
    };

    const buy = async (farmID: number) => {
        try {
            const newPlot: Plot = await apiBuyPlot(farmID);
            plots.value.set(newPlot.id, newPlot);
        } catch(err) {
            if (err instanceof Error) {
                error.value = err.message;
            } else {
                error.value = String(err);
            }
        }
    };

    const plant = async (plotID: number, cropID: number) => {
        try {
            const updatedPlot: Plot = await apiPlantPlot(plotID, cropID);
            plots.value.set(updatedPlot.id, updatedPlot);
        } catch(err) {
            if (err instanceof Error) {
                error.value = err.message;
            } else {
                error.value = String(err);
            }
        }
    };

    const harvest = async (plotID: number) => {
        try {
            const updatedPlot: Plot = await apiHarvestPlot(plotID);
            plots.value.set(updatedPlot.id, updatedPlot);
        } catch(err) {
            if (err instanceof Error) {
                error.value = err.message;
            } else {
                error.value = String(err);
            }
        }
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
        plots,
        buy,
        plant,
        harvest,
        getPlot,
        getAllPlots,
        clearPlots,
        error,
    }
});