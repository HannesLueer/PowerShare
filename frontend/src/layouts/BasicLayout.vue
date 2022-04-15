<template>
  <main>
    <slot></slot>
  </main>
</template>

<style scoped>
main {
  margin: 40px;
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
