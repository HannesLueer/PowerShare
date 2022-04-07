<script setup lang="ts">
defineProps<{
  msg: string;
}>();
</script>

<template>
  <div class="ErrorBox">
    <img src="@/assets/img/error.svg" alt="Error-Logo" />
    {{ msg }}
  </div>
</template>

<style scoped>
div {
  background-color: var(--color-errorbox-background);
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
