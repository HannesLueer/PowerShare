<template>
  <main>
    <slot></slot>
  </main>
</template>

<style scoped>
main {
  margin-left: 40px;
  margin-right: 40px;
  margin-bottom: 40px;
}

@media screen and (max-width: 600px) {
  main {
    margin-left: 10px;
    margin-right: 10px;
    margin-bottom: 10px;
  }
}
</style>
