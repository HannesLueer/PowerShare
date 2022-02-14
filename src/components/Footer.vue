<script setup lang="ts">

</script>

<template>
  <div class="footer">
    <p>Author: Hannes Lüer</p>
    <hr>
    <p>
      Mail:
      <a href="mailto:hannes-lueer+PowerShare@outlook.de">
        hannes-lueer@outlook.de
      </a>
    </p>
    <hr>
    <p>{{ new Date().getFullYear() }}</p>
  </div>
</template>

<style scoped>
.footer {
  position: fixed;
  left: 0;
  bottom: 0;
  width: calc(100% - 1em);
  background-color: var(--color-footer-background);
  color: var(--color-text);
  text-align: center;
  padding: 0.5em;
}

.footer * {
  margin: 0;
  display: inline-block;
}

.footer hr {
  border: none;
  border-top: 0.7em dotted var(--gray500);
  width: 0.7em;
  margin: auto 20px auto 20px;
}
</style>
