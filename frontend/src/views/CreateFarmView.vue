<template>
    <div class="create-farm">
        <h2>Create Your Farm</h2>

        <form @submit.prevent="handleSubmit">
            <div class="form-group">
                <label for="farmName">Farm Name</label>
                <br />
                <input
                    id="farmName"
                    v-model="name"
                    type="text"
                    placeholder="e.g. Tegridy Farms"
                    :disabled="isSubmitting"
                />
            </div>

            <div v-if="error" class="form-error">
                {{ error }}
            </div>

            <button 
                type="submit"
                class=""
                :disabled="isSubmitting"
            >
                {{ isSubmitting ? "Creating ..." : "Create Farm" }}
            </button>
        </form>
    </div>
</template>

<script setup lang="ts">
import { useFarmStore } from '@/stores/farmStore';
import type { CreateFarmPayload } from '@/types/farm';
import { ref } from 'vue';

const farmStore = useFarmStore();

const name = ref('');
const error = ref<string | null>(null);
const isSubmitting = ref(false);

const handleSubmit = async () => {
    error.value = null;

    if (!name.value.trim()) {
        error.value = "Farm name required."
        return;
    }

    isSubmitting.value = true;
    
    const payload: CreateFarmPayload = {
        name: name.value.trim(),
    };

    try {
        await farmStore.create(payload);
        name.value = "";
    } catch (err) {
        error.value = "Failed to create farm."
        console.error(err)
    }
    finally {
        isSubmitting.value = false;
    }
};
</script>

<style lang="css" scoped>

.create-farm {
  max-width: 400px;
  padding: 1rem;
  border: 1px solid #ccc;
  border-radius: 8px;
}

.form-group {
  margin-bottom: 1rem;
}

label {
  display: block;
  font-weight: bold;
  margin-bottom: 0.5rem;
}

input[type="text"] {
  width: 100%;
  padding: 0.5rem;
  font-size: 1rem;
}

button {
  padding: 0.5rem 1rem;
  font-size: 1rem;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button[disabled] {
  opacity: 0.6;
  cursor: not-allowed;
}

.form-error {
  color: red;
  margin-top: 0.5rem;
}
</style>