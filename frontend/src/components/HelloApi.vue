<script setup lang="ts">
import { ref, onMounted } from 'vue';

const message = ref<string>('Loading...');

onMounted(async () => {
  try {
    const res = await fetch('/api/hello');
    message.value = await res.text();
  } catch (err) {
    message.value = 'Failed to fetch backend.';
    console.error(err);
  }
});
</script>

<template>
  <div class="hello-api">
    <h2>Backend Says:</h2>
    <pre>{{ message }}</pre>
  </div>
</template>

<style scoped>
.hello-api {
  font-family: monospace;
}
</style>
