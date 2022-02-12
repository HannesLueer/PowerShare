<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
import HelloWorld from "@/components/HelloWorld.vue";

</script>

<template>
  <header>
    <img
      alt="Vue logo"
      class="logo"
      src="@/assets/img/logo.svg"
      width="125"
      height="125"
    />

    <div class="wrapper">
      <HelloWorld msg="You did it!" />

      <nav>
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/about">About</RouterLink>
      </nav>

    </div>
  </header>
      <router-view></router-view>
</template>

<style>
@import "@/assets/css/base.css";

</style>
