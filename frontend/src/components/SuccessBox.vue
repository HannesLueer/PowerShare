<script setup lang="ts">
defineProps<{
  msg: string;
}>();
</script>

<template>
  <div class="SuccessBox">
    <img src="@/assets/img/info.svg" alt="Info-Logo" />
    {{ msg }}
  </div>
</template>

<style scoped>
div {
  background-color: var(--color-successbox-background);
  padding: 0.3em 0.5em 0.3em 0.5em;
  border-radius: 0.1em;
  display: flex;
  align-items: center;
  gap: 0.5em;
}

img {
  height: 2em;
  width: 2em;
}
</style>
