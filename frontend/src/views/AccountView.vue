<script setup lang="ts">
import { onMounted, ref } from "vue";
import DefaultLayout from "../layouts/BasicLayout.vue";
import ErrorBox from "@/components/ErrorBox.vue";
import SuccessBox from "@/components/SuccessBox.vue";
import { userService } from "@/services";

let name = ref<string>("");
let email = ref<string>("");
let password = ref<string>("");

let errMsg = ref<string>("");
let sucMsg = ref<string>("");

async function updateUser() {
  const data = await userService.update(
    name.value,
    email.value,
    password.value
  );
  if (data.token) {
    errMsg.value = "";
    sucMsg.value = "Data was changed successfully";
  } else {
    errMsg.value = data;
    sucMsg.value = "";
  }
}

async function deleteUser() {
  const data = await userService.remove();
  if (data == "") {
    errMsg.value = "";
  } else {
    errMsg.value = data;
    sucMsg.value = "";
  }
}

async function getCurrentValues() {
  let data = await userService.get();

  if (data.name) {
    errMsg.value = "";
    name.value = data.name;
    email.value = data.email;
  } else {
    errMsg.value = data;
  }
}

onMounted(async () => {
  await getCurrentValues();
});
</script>

<template>
  <DefaultLayout>
    <h1>Account Settings</h1>

    <form @submit.prevent="updateUser" accept-charset="UTF-8">
      <label for="name">Name</label>
      <input
        v-model="name"
        type="text"
        id="name"
        required
        placeholder="name"
        autocomplete="name"
        autofocus
      />
      <br />

      <label for="email">Email</label>
      <input
        v-model="email"
        type="text"
        id="email"
        required
        placeholder="email"
        autocomplete="email"
      />
      <br />

      <label for="password">Password</label>
      <input
        v-model="password"
        type="password"
        id="password"
        required
        placeholder="password"
        autocomplete="current-password"
      />
      <br />
      <button type="submit">submit changes</button>
    </form>

    <button @click="deleteUser" class="delete">delete account</button>

    <ErrorBox :msg="errMsg" v-if="errMsg != ''"></ErrorBox>
    <SuccessBox :msg="sucMsg" v-if="sucMsg != ''"></SuccessBox>
  </DefaultLayout>
</template>

<style scoped>
label {
  font-size: 1.2em;
  font-weight: bold;
  display: block;
  margin-top: 1em;
  margin-bottom: 1em;
}

button {
  margin-top: 2em;
}

.ErrorBox,
.SuccessBox {
  margin-top: 2em;
}
</style>
</style>
