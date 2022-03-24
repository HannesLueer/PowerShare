<!-- Based on https://codepen.io/1832Manaswini/pen/Vwezyjx -->

<template>
  <main>
    <div class="container">
      <h1 class="first-number">{{ errNo[0] }}</h1>
      <div class="cog-wheel">
        <div class="cog">
          <div class="top"></div>
          <div class="down"></div>
          <div class="left-top"></div>
          <div class="left-down"></div>
          <div class="right-top"></div>
          <div class="right-down"></div>
          <div class="left"></div>
          <div class="right"></div>
        </div>
      </div>
      <h1 class="third-number">{{ errNo[2] }}</h1>
      <p class="wrong-para">Uh Oh! {{ errStatusText }}!</p>
    </div>
  </main>  
</template>

<script setup lang="ts">
import { onMounted, ref } from "vue";
import { getReasonPhrase } from "http-status-codes";
import { useRoute } from "vue-router";

let errNo = ref<string[]>(["4", "0", "4"]);
let errStatusText = ref<string>("Something is wrong");

const getCode = () => {
  const route = useRoute()
  const errCode = route.query.err?.toString();

  if (errCode) {
    try {
      const reasonPhrase = getReasonPhrase(errCode);
      console.log("hallo")
      errStatusText.value = reasonPhrase;
      errNo.value = errCode.split("", 3);
    } catch (error) {
      // invalid error code
    }
  }
}

onMounted(() => {
  getCode();
});
</script>

<style scoped>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.container {
  margin-top: calc(10vh - 3rem);
  width: 100vw;
  height: auto;
  display: flex;
  justify-content: center;
  font-family: "Poppins", sans-serif;
  position: relative;
  align-items: center;
}

.cog-wheel {
  transform: scale(0.7);
}

.cog {
  width: 40vmin;
  height: 40vmin;
  border-radius: 50%;
  border: 6vmin solid var(--aqua-green);
  position: relative;
  top: 2vh;

  animation-name: spin;
  animation-duration: 10s;
  animation-iteration-count: infinite;
  animation-timing-function: linear;
}

@keyframes spin {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}

div.cog div {
  width: 10vmin;
  height: 10vmin;
  background-color: var(--aqua-green);
  position: absolute;
  border-radius: 5%;
}

.top {
  top: -14vmin;
  left: 9vmin;
}

.down {
  bottom: -14vmin;
  left: 9vmin;
}

.left {
  left: -14vmin;
  top: 9vmin;
}

.right {
  right: -14vmin;
  top: 9vmin;
}

.left-top {
  transform: rotateZ(-45deg);
  left: -7.5vmin;
  top: -7.5vmin;
}

.left-down {
  transform: rotateZ(45deg);
  left: -7.5vmin;
  bottom: -7.5vmin;
}

.right-top {
  transform: rotateZ(45deg);
  right: -7.5vmin;
  top: -7.5vmin;
}

.right-down {
  transform: rotateZ(-45deg);
  right: -7.5vmin;
  bottom: -7.5vmin;
}

h1 {
  color: var(--color-button-border);
}

.first-number {
  position: relative;
  top: 6vh;
  left: 7vmin;
  z-index: 1;
  font-size: 40vmin;
}

.third-number {
  position: relative;
  top: -2vh;
  right: 7vmin;
  z-index: -1;
  font-size: 40vmin;
}

.wrong-para {
  font-family: "Montserrat", sans-serif;
  position: absolute;
  bottom: -10vh;
  padding: 3vmin;
  font-weight: 600;
  color: var(--color-text);
  text-align: center;
}
</style>
