<script setup lang="ts">
import { RouterLink } from "vue-router";

detectColorScheme();

function detectColorScheme() {
  let theme = "dark";
  if (localStorage.getItem("theme")) { // already stored in localstorage
    theme = localStorage.getItem("theme") || "dark";
  } else {
    if (window.matchMedia) {  // matchMedia method supported
      theme = window.matchMedia("(prefers-color-scheme: dark)").matches
        ? "dark"
        : "light";
    } else {  // matchMedia method not supported
      theme = "dark";
    }
  }

  changeColorSchemeTo(theme);
}

function changeColorSchemeTo(theme: string) {
  localStorage.setItem("theme", theme);
  document.querySelector(":root")?.setAttribute("theme", theme);
}

function themeButtonClicked() {
  if (localStorage.getItem("theme") == "dark") {
    localStorage.setItem("theme", "light");
  } else {
    localStorage.setItem("theme", "dark");
  }

  changeColorSchemeTo(localStorage.getItem("theme") || "dark");
}
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
    };
  },
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

      <li class="right">
        <button class="themeSwitch" @click="themeButtonClicked">
          <svg id="sun-icon" viewBox="0 0 100 100">
            <use href="@/assets/img/sun-icon.svg#sun"></use>
          </svg>
          <svg id="moon-icon" viewBox="0 0 100 100">
            <use href="@/assets/img/moon-icon.svg#moon"></use>
          </svg>
        </button>
      </li>
    </ul>
  </nav>
</template>

<style scoped>
ul {
  width: 100%;
  background-color: var(--color-navbar-background);
  color: var(--color-text);

  list-style-type: none;
  margin: 0;
  padding: 0;
  overflow: hidden;
}

ul * {
  margin: auto;
  display: inline-block;

  vertical-align: middle;
}

li {
  float: left;
}

li.right {
  float: right;
}

li a {
  display: block;
  color: var(--color-text);
  text-align: center;
  padding: 0.8rem 1rem;
  text-decoration: none;
}

li > img {
  display: block;
  width: auto;
  height: calc(1rem + 1.1rem);
  padding: 0.5rem;
  object-fit: contain;
}

li a:hover {
  background-color: var(--color-navbar-hover-background);
}

/* theme switch */
button.themeSwitch {
  border-radius: 50%;
  width: 2.1rem;
  height: 2.1rem;
  margin: 0.5rem;
}

[theme="dark"] svg#moon-icon {
  display: none;
}

[theme="light"] svg#sun-icon {
  display: none;
}

button.themeSwitch svg {
  fill: var(--color-button-text);
}

button.themeSwitch svg:hover {
  animation-name: wiggle;
  animation-duration: 0.5s;
}

@keyframes wiggle {
  0% {
    transform: rotate(0deg);
  }
  40% {
    transform: rotate(-16deg);
  }
  80% {
    transform: rotate(8deg);
  }
  100% {
    transform: rotate(0deg);
  }
}
</style>
