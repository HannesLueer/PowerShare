<script setup lang="ts">
import { RouterLink } from "vue-router";
</script>

<script lang="ts">
import { defineComponent } from "vue";
export default defineComponent({
  name: "NavBar",
  components: {},
  data() {
    return {
      links: [
        { to: "/", text: "Home" },
        { to: "/about", text: "About" },
      ],
    }
  }
});
</script>

<template>
  <nav>
    <ul>
      <li>
        <img
          alt="Vue logo"
          class="logo"
          src="@/assets/img/logo.svg"
          width="20"
          height="20"
        />
      </li>

      <li v-for="(link, index) in links" :key="index">
        <RouterLink :to="link.to">{{ link.text }}</RouterLink>
      </li>
    </ul>
  </nav>
</template>

<style scoped>
ul {
  width: 100%;
  background-color: var(--color-navbar-background);
  color: var(--color-text);
}

/* ul * {
  margin: 0;
  display: inline-block;
} */


ul {
  list-style-type: none;
  margin: 0;
  padding: 0;
  overflow: hidden;
}

li {
  float: left;
}

li a {
  display: block;
  color: var(--color-text);
  text-align: center;
  padding: 14px 16px;
  text-decoration: none;
}

li img {
  display: block;
  width: auto;
  height: calc(1em + 20px);
  padding: 8px;
  object-fit: contain;
}

li a:hover {
  background-color: var(--color-navbar-hover-background);
}
</style>
