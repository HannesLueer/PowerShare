<script setup lang="ts">
import { onMounted, ref } from "vue";
import { RouterLink } from "vue-router";
import { userService } from "@/services";
import LoginLayout from "@/layouts/LoginLayout.vue";
import ErrorBox from "@/components/ErrorBox.vue";
import router from "@/router";

let email = ref<string>("");
let password = ref<string>("");

let errMsg = ref<string>("");

async function login() {
  const data = await userService.login(email.value, password.value);
  if (data.token) {
    errMsg.value = "";
    router.back();
  } else {
    errMsg.value = data;
  }
}

onMounted(() => {
  if (userService.isLoggedin.value) {
    router.back();
  }
});
</script>

<template>
  <LoginLayout>
    <h1 class="login">LOGIN</h1>

    <form @submit.prevent="login" accept-charset="UTF-8">
      <input
        v-model="email"
        type="text"
        name="txtUser"
        required
        placeholder="email"
        autocomplete="email"
        autofocus
      />
      <input
        v-model="password"
        type="password"
        name="txtPassword"
        required
        placeholder="password"
        autocomplete="current-password"
      />
      <button type="submit">login</button>
    </form>

    <ErrorBox :msg="errMsg" v-if="errMsg != ''"></ErrorBox>

    Don't you have an account yet?
    <RouterLink to="/register">Register now!</RouterLink>
  </LoginLayout>
</template>
