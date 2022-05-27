<script setup lang="ts">
import Map from "../components/Map.vue";
import { useRoute, type RouteLocationNormalizedLoaded } from "vue-router";
import { onMounted, ref, watch } from "vue";
import { chargerService, type ChargerData } from "@/services";
import InfoBox from "../components/InfoBox.vue";
import { gocardlessService, chargingService } from "@/services";
import ErrorBox from "../components/ErrorBox.vue";
import SuccessBox from "@/components/SuccessBox.vue";

let mapCenter = ref<[number, number]>([51.5, 10]);
let invalidateSizeTrigger = ref<number>(0);

var charger = ref<ChargerData>();

let newMandateURL = ref<string>("");
let isCharging = ref<boolean>(false);
let errMsg = ref<string>("");
let sucMsg = ref<string>("");

async function getChargerData(id: number): Promise<void> {
  if (isNaN(id)) {
    charger.value = undefined;
    return;
  }

  let newCharger = await chargerService.get(id);

  if (newCharger != undefined) {
    mapCenter.value = [newCharger.position.Lat, newCharger.position.Lng];
    charger.value = newCharger;
  } else {
    // TODO: display error to user
  }
}

async function onLoad(route: RouteLocationNormalizedLoaded) {
  hideMessageBoxes();

  if (typeof route.params.id == "string")
    await getChargerData(parseInt(route.params.id));

  invalidateSizeTrigger.value++;
}

async function getNewMandateURL() {
  newMandateURL.value = await gocardlessService.get_newMandateURL();
}

async function startCharging() {
  hideMessageBoxes();

  if (charger.value != undefined) {
    const data = await chargingService.start(charger.value.id);
    if (data == "") {
      displaySuccess("Charging started successfully");
    } else {
      displayError(data);
    }
  }

  updateIsCharging();
}

async function stopCharging() {
  hideMessageBoxes();
  if (charger.value != undefined) {
    const data = await chargingService.stop(charger.value.id);
    if (data == "") {
      displaySuccess("Charging stopped successfully");
    } else {
      displayError(data);
    }
  }
  updateIsCharging();
}

async function updateIsCharging() {
  if (charger.value != undefined)
    isCharging.value =
      (await chargingService.isThisUserCharging(charger.value.id)) == "true"
        ? true
        : false;
}

function displayError(text: string) {
  errMsg.value = text;
  sucMsg.value = "";
}

function displaySuccess(text: string) {
  errMsg.value = "";
  sucMsg.value = text;
}

function hideMessageBoxes() {
  errMsg.value = "";
  sucMsg.value = "";
}

onMounted(async () => {
  const route = useRoute();

  // used for classic URL navigation (someone sends link to charger)
  onLoad(route);
  getNewMandateURL();

  // used for vue router navigation
  watch(route, async () => {
    onLoad(route);
  });
});
</script>

<template>
  <main v-bind:class="[charger != undefined ? 'split50' : '']">
    <Map
      :center="mapCenter"
      :use-search="true"
      :use-location="true"
      markersApiPath="/charger/all"
      :markersUpdateIntervalSeconds="120"
      marker-link-to="/charger/"
      :use-manual-update-button="true"
      :trigger-map-invalidate-size="invalidateSizeTrigger"
      class="split50"
    >
    </Map>

    <div class="split50 textbox" v-if="charger != undefined">
      <h1>{{ charger?.title }}</h1>

      <InfoBox
        >To activate a charging station please create a mandate at gocardless
        via <a :href="newMandateURL">this link</a>. Please use the same email
        address as for PowerShare. If you have already created a mandate, simply
        click on the button to unlock the charger.</InfoBox
      >

      <button v-if="!isCharging && !charger.isOccupied" @click="startCharging">
        start charging
      </button>
      <button v-if="isCharging" @click="stopCharging">stop charging</button>

      <ErrorBox :msg="errMsg" v-if="errMsg != ''"></ErrorBox>
      <SuccessBox :msg="sucMsg" v-if="sucMsg != ''"></SuccessBox>

      <br />

      available: {{ !charger.isOccupied }}

      <h2>description</h2>
      {{ charger.description }}

      <h2>address</h2>
      Lat: {{ charger.position.Lat }} <br />
      Lng: {{ charger.position.Lng }} <br />

      <h2>cost</h2>
      {{ charger?.cost.amount }} {{ charger?.cost.currency.symbol }}
    </div>
  </main>
</template>

<style scoped>
main {
  display: flex;
  flex-flow: column;
  flex-direction: row;
  flex-grow: 1;
}

Map {
  height: fit-content;
}

main > div {
  padding-left: 2em;
  padding-right: 2em;
}

main > div {
  width: calc(100vw - 4em);
}
main.split50 > div.split50 {
  width: calc(50vw - 4em);
}

@media screen and (orientation: portrait) and (max-width: 800px) {
  main {
    flex-direction: column;
  }
  main.split50 > div.split50 {
    width: calc(100vw - 4em);
    height: 50%;
  }
}

.textbox {
  overflow-y: auto;
}

button {
  margin-bottom: 1em;
}
</style>
