<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
import Footer from "./components/Footer.vue";
import Navbar from "./components/Navbar.vue";
</script>

<template>
  <Navbar></Navbar>

  <router-view></router-view>

  <Footer></Footer>
</template>

<style>
@import "@/assets/css/base.css";

</style>
