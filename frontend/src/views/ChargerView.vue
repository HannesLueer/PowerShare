<script setup lang="ts">
import Map from "../components/Map.vue";
import { useRoute } from "vue-router";
import { onMounted, ref, watch } from "vue";

let mapCenter = ref<[number, number]>([51.5, 10]);

class ChargerData {
  id!: number;
  title!: string;
  position!: { Lat: number; Lng: number };
  cost!: number;
  isOccupied!: boolean;
}

var charger = ref<ChargerData>();

async function getChargerData(id: number): Promise<void> {
  if (isNaN(id)) {
    charger.value = undefined;
    return;
  }
  let url = "https://localhost:5000/api/v1/charger/" + id;
  let response = await fetch(url);
  if (response.ok) {
    charger.value = await response.json();
    if (charger.value != undefined) {
      mapCenter.value = [
        charger.value?.position.Lat,
        charger.value?.position.Lng,
      ];
    }
  } else {
    // TODO: display error to user
    console.log("HTTP-Error: " + response.status);
  }
}

onMounted(async () => {
  const route = useRoute();
  // used for classic URL navigation (someone sends link to charger)
  if (typeof route.params.id == "string")
    await getChargerData(parseInt(route.params.id));
  // used for vue router navigation
  watch(route, async () => {
    if (typeof route.params.id == "string")
      await getChargerData(parseInt(route.params.id));
  });
});
</script>

<template>
  <main v-bind:class="[charger != undefined ? 'split50' : '']">
    <Map
      :center="mapCenter"
      :use-search="true"
      :use-location="true"
      markersURL="https://localhost:5000/api/v1/charger/all"
      :markersUpdateIntervalSeconds="120"
      class="split50"
    >
    </Map>

    <div class="split50 textbox" v-if="charger != undefined">
      <h1>{{ charger?.title }}</h1>
      Details text <br />
      id: {{ charger.id }} <br>
      isOccupied: {{ charger.isOccupied }}
      <h2>address</h2>
      Lat: {{ charger.position.Lat }} <br>
      Lng: {{ charger.position.Lng }} <br>
      ...

      <h2>cost</h2>
      {{ charger?.cost }} €
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

@media screen and (max-width: 800px) {
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
</style>
