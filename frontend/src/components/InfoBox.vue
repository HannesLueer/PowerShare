<template>
  <div class="InfoBox">
    <slot></slot>
  </div>
</template>

<style scoped>
div {
  background-color: var(--color-infobox-background);
  padding: 0.3em 0.5em 0.3em 0.5em;
  border-radius: 0.1em;
  margin-block: 1em;
}
</style>
