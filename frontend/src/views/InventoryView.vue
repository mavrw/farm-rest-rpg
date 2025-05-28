<template>
    <div class="p-4">
        <h1 class="text-2xl font-bold mb-4">Inventory</h1>

        <div v-if="error" class="text-red-500 mb-4">
            Error: {{  error }}
        </div>
        
        <div v-if="inventory.items.length === 0 && !error" class="text-gray-500">
            No items in your inventory.
        </div>

        <ul v-else class="space-y-2">
            <li
                v-for="item in inventory.items"
                :key="item.ID"
                class="p-2 border rounded shadow-sm bg-white"
            >
                <div><strong>Item ID:</strong> {{ item.ItemID }}</div>
                <div><strong>Quantity:</strong> {{  item.Quantity }}</div>
            </li>
        </ul>
    </div>
</template>

<script setup lang="ts">

import { onMounted } from 'vue';
import { useInventoryStore } from '@/stores/inventoryStore';

const inventory = useInventoryStore();

onMounted(() => {
    inventory.fetchAll();
});

const items = inventory.items;
const error = inventory.error;

</script>