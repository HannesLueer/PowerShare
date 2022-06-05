<script setup lang="ts">
import { RouterLink } from "vue-router";
import { userService } from "@/services";

detectColorScheme();

function detectColorScheme() {
  let theme = "dark";
  if (localStorage.getItem("theme")) {
    // already stored in localstorage
    theme = localStorage.getItem("theme") || "dark";
  } else {
    if (window.matchMedia) {
      // matchMedia method supported
      theme = window.matchMedia("(prefers-color-scheme: dark)").matches
        ? "dark"
        : "light";
    } else {
      // matchMedia method not supported
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
      linksLeft: [
        { to: "/", text: "Home" },
        { to: "/charger/map", text: "Map" },
        { to: "/about", text: "About" },
      ],

      linksRight: {
        user: [
          { to: "/mycharger/all", text: "Manage chargers" },
          { to: "/account", text: "Account" },
        ],
        guest: [
          { to: "/login", text: "Login" },
          { to: "/register", text: "Register" },
        ],
      },
    };
  },
});
</script>

<template>
  <nav>
    <ul class="left">
      <li>
        <img
          alt="Vue logo"
          class="logo"
          src="@/assets/img/logo.svg"
          width="20"
          height="20"
        />
      </li>

      <li v-for="(link, index) in linksLeft" :key="index">
        <RouterLink :to="link.to">{{ link.text }}</RouterLink>
      </li>
    </ul>

    <ul class="right">
      <li v-for="(link, index) in linksRight.guest" :key="index">
        <RouterLink :to="link.to" v-if="!userService.isLoggedin.value">{{
          link.text
        }}</RouterLink>
      </li>

      <li v-for="(link, index) in linksRight.user" :key="index">
        <RouterLink :to="link.to" v-if="userService.isLoggedin.value">{{
          link.text
        }}</RouterLink>
      </li>

      <li>
        <button
          id="logout-button"
          @click="userService.logout()"
          v-if="userService.isLoggedin.value"
        >
          Logout
        </button>
      </li>

      <li>
        <button
          class="themeSwitch"
          @click="themeButtonClicked"
          aria-label="Theme Switch"
        >
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
nav {
  width: 100%;
  background-color: var(--color-navbar-background);
  color: var(--color-text);
  display: flex;
}

ul {
  list-style-type: none;
  margin: 0;
  padding: 0;
  overflow: hidden;
}

ul.left {
  margin-right: auto;
}

ul.left {
  align-items: flex-end;
}

ul * {
  display: inline-block;
  vertical-align: middle;
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

/* logout button */
button#logout-button {
  border-radius: 0.3em;
  margin: auto 0.8rem;
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
