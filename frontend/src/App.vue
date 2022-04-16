<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
import Footer from "./components/Footer.vue";
import Navbar from "./components/Navbar.vue";
</script>

<template>
  <div class="page-wrap">
    <Navbar></Navbar>

    <router-view></router-view>

    <Footer></Footer>
  </div>
</template>

<style>
@import "@/assets/css/base.css";
@import "@/assets/css/scrollbar.css";
</style>

<style scoped>
div.page-wrap {
  height: 100%;
  display: flex;
  flex-flow: column;
}
</style>
