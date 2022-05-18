<template>
  <main>
    <slot></slot>
  </main>
</template>

<style scoped>
main {
  padding: 40px;
  overflow: auto;
}

main :deep() *:first-child {
  margin-top: 0;
}

@media screen and (max-width: 600px) {
  main {
    margin: 10px;
  }
}
</style>
