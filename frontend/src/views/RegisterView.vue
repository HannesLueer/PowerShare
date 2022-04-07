<script setup lang="ts">
import { onMounted, ref } from "vue";
import { RouterLink } from "vue-router";
import { userService } from "@/services";
import LoginLayout from "@/layouts/LoginLayout.vue";
import ErrorBox from "@/components/ErrorBox.vue";
import router from "@/router";

let name = ref<string>("");
let email = ref<string>("");
let password = ref<string>("");

let errMsg = ref<string>("");

async function register() {
  const data = await userService.register(
    name.value,
    email.value,
    password.value
  );
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
    <div>
      <h1 class="login">REGISTRATION</h1>

      <form @submit.prevent="register" accept-charset="UTF-8">
        <input
          v-model="name"
          type="text"
          name="txtUser"
          required
          placeholder="name"
          autocomplete="name"
          autofocus
        />
        <input
          v-model="email"
          type="text"
          name="txtUser"
          required
          placeholder="email"
          autocomplete="email"
        />
        <input
          v-model="password"
          type="password"
          name="txtPassword"
          required
          placeholder="password"
          autocomplete="current-password"
        />
        <button type="submit">register</button>
      </form>

      <ErrorBox :msg="errMsg" v-if="errMsg != ''"></ErrorBox>

      Do you already have an account?
      <RouterLink to="/login">Login now!</RouterLink>
    </div>
  </LoginLayout>
</template>
