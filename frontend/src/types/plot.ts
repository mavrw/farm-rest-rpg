export interface Plot {
    id: number,
    farm_id: number,
    name: string,
    crop_id: number,
    planted_at: string,
    harvest_at: string,
}

export interface PlotResponse {
    id: number,
    farm_id: number,
    name: string,
    crop_id: number,
    planted_at: string,
    harvest_at: string,
}