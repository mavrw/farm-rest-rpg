export interface CreateFarmPayload {
    name: string,
}

export interface FarmResponse {
    ID: number,
    UserID: number,
    Name: string,
    CreatedAt: string,
    UpdatedAt: string,
}

export interface Farm {
    ID: number,
    UserID: number,
    Name: string,
    CreatedAt: string,
    UpdatedAt: string,
}