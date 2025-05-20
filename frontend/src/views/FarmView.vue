<template>
    <div class="farm-container">
        <h1 class="farm-title">Your Farm</h1>
        
        <div v-if="farmStore.error" class="farm-error">
            {{ farmStore.error }}
        </div>

        <CreateFarmView v-if="!farmStore.hasFarm" />

        <div v-else class="farm-details">
            <p><strong>Name:</strong> {{ farmStore.farm?.Name }}</p>
            <p><strong>Created:</strong> {{ farmStore.farm?.CreatedAt }}</p>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useFarmStore } from '@/stores/farmStore';
import { onMounted } from 'vue';
import CreateFarmView from './CreateFarmView.vue';

const farmStore = useFarmStore();

onMounted(async () => {
    await farmStore.get();
});
</script>

<style scoped>
.farm-container {
  padding: 1rem;
}

.farm-title {
  font-size: 1.5rem;
  font-weight: bold;
  margin-bottom: 1rem;
}

.farm-error {
  color: red;
  margin-bottom: 1rem;
}

.farm-details {
  border: 1px solid #ccc;
  padding: 1rem;
  border-radius: 6px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}
</style>