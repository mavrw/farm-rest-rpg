export interface InventoryItem {
    ID: number,
    UserID: number,
    ItemID: number,
    Quantity: number,
}

export type InventoryResponse = InventoryItem[]