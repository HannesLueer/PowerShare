import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import { registerSW } from "virtual:pwa-register";

const app = createApp(App);

app.use(router);

app.mount("#app");

const updateSW = registerSW({
  immediate: true,
  onNeedRefresh() {
    console.log("show a prompt to the user with refresh and cancel buttons")
  },
  onOfflineReady() {
    console.log("show a ready to work offline message to the user with an OK button")
  },
});
