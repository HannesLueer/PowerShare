<script setup lang="ts">
import Map from "@/components/Map.vue";

var center: [number, number] = [51.5, 10];
</script>

<template>
  <main>
    <div>
      <h1>Map</h1>
    </div>

    <Map
      :center="center"
      :use-search="true"
      :use-location="true"
      markersURL="https://localhost:5000/api/v1/charger/all"
      :markersUpdateIntervalSeconds="120"
      marker-link-to="/charger/"
    ></Map>
  </main>
</template>

<style scoped>
main {
  display: flex;
  flex-flow: column;
}

Map {
  flex-grow: 1;
}
</style>
