<script setup lang="ts">
import { onMounted, ref, watch } from "vue";
import Map from "@/components/Map.vue";
import {
  chargerService,
  currencyService,
  type Coordinate,
  type Cost,
  type Currency,
  type ChargerData,
  smartmeService,
} from "@/services";
import { useRoute, type LocationQuery } from "vue-router";
import ErrorBox from "@/components/ErrorBox.vue";
import SuccessBox from "@/components/SuccessBox.vue";
import isEqual from "lodash.isequal";
import { getShellyConnectLink, getSmartMeConnectLink } from "@/helpers";

const defaultCharger = <ChargerData>{
  id: -1,
  title: "",
  position: { Lat: 0, Lng: 0 },
  cost: <Cost>{
    amount: 0,
    currency: <Currency>{ abbreviation: "EUR" },
  },
  isOccupied: false,
  technicalData: {
    shellyDeviceId: 0,
  },
};

let mapCenter = ref<[number, number]>([51.5, 10]);

let charger = ref<ChargerData>(defaultCharger);
let currencies = ref<Currency[]>([]);

let errMsg = ref<string>("");
let sucMsg = ref<string>("");

let markerTrigger = ref<number>(0);

function setLocation(coordinate: Coordinate) {
  charger.value.position = coordinate;
}

async function getCurrencies(): Promise<void> {
  let newCurrenciesList = await currencyService.getList();

  if (newCurrenciesList && newCurrenciesList != []) {
    currencies.value = newCurrenciesList;
  } else {
    displayError("An error occurred while loading the currencies");
  }
}

async function getChargerData(id: number): Promise<void> {
  if (isNaN(id)) {
    charger.value = defaultCharger;
    return;
  }

  let newCharger = await chargerService.get(id);

  if (newCharger.id != undefined) {
    mapCenter.value = [newCharger.position.Lat, newCharger.position.Lng];
    charger.value = newCharger;
  } else {
    displayError("An error occurred while loading the charger data");
  }
}

function submit() {
  if (charger.value.id == defaultCharger.id) createCharger();
  else updateCharger();
}

async function createCharger() {
  const data = await chargerService.create(charger.value);
  if (!isNaN(Number(data))) {
    displaySuccess("Charger was added successfully");
    updateMarkers();
  } else {
    displayError(data);
  }
}

async function updateCharger() {
  const data = await chargerService.update(charger.value);
  if (data == charger.value.id.toString()) {
    displaySuccess("Data was changed successfully");
    updateMarkers();
  } else {
    displayError(data);
  }
}

async function deleteCharger() {
  const data = await chargerService.remove(charger.value.id);
  if (data == "") {
    displaySuccess("Charger was deleted successfully");
    updateMarkers();
  } else {
    displayError(data);
  }
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

function updateMarkers() {
  markerTrigger.value++;
}

function handleSmartMeAuth(params: LocationQuery) {
  // send code to backend
  if (params.code != undefined) {
    smartmeService.post_authCode(params.code.toString());
  }

  // use state to keep previously entered but not submitted values
  if (params.state != undefined) {
    charger.value = JSON.parse(params.state.toString());
  }
}

onMounted(async () => {
  const route = useRoute();
  // used for classic URL navigation (e.g saved link)
  if (typeof route.params.id == "string")
    await getChargerData(parseInt(route.params.id));
  getCurrencies();
  handleSmartMeAuth(route.query);

  // used for vue router navigation
  watch(route, async () => {
    hideMessageBoxes();
    if (typeof route.params.id == "string")
      await getChargerData(parseInt(route.params.id));
  });
});
</script>

<template>
  <main>
    <Map
      :center="mapCenter"
      :use-search="true"
      :use-location="true"
      markersApiPath="/charger/my"
      :markersUpdateIntervalSeconds="120"
      marker-link-to="/mycharger/"
      :useAddButton="true"
      :use-manual-update-button="true"
      v-on:clicked-position="setLocation"
      class="split50"
      :trigger-marker-update="markerTrigger"
    >
    </Map>

    <div class="split50 textbox">
      <form @submit.prevent="submit" accept-charset="UTF-8">
        <label for="chargerName">Name</label>
        <input
          v-model="charger.title"
          type="text"
          id="chargerName"
          required
          placeholder="name"
        />
        <br />

        <label for="description">Description</label>
        <textarea
          v-model="charger.description"
          id="description"
          placeholder="description"
        />

        <label for="latitude">Latitude</label>
        <input
          v-model="charger.position.Lat"
          type="number"
          step="any"
          id="latitude"
          required
          placeholder="latitude"
        />
        <br />

        <label for="longitude">Longitude</label>
        <input
          v-model="charger.position.Lng"
          type="number"
          step="any"
          id="longitude"
          required
          placeholder="longitude"
        />
        <br />

        <label for="cost">Cost</label>
        <input
          v-model="charger.cost.amount"
          type="number"
          step="0.01"
          id="cost"
          required
          placeholder="cost"
        />
        <select v-model="charger.cost.currency.abbreviation">
          <option v-for="currency in currencies" :value="currency.abbreviation">
            {{ currency.symbol }} ({{ currency.abbreviation }})
          </option>
        </select>
        <br />

        <div v-if="charger.technicalData != undefined">
          <label>Shelly Account</label>
          Connect your Shelly account with PowerShare and grant access to the
          device <br />
          <a :href="getShellyConnectLink()"> Shelly &#8599;</a>

          <label for="shellyDeviceId">Shelly Device-ID</label>
          <input
            v-model="charger.technicalData.shellyDeviceId"
            type="number"
            step="1"
            id="shellyDeviceId"
            required
            placeholder="shelly device id"
          />
          <br />

          <label>smart-me Account</label>
          Connect your Smart-Me account with PowerShare <br />
          <a :href="getSmartMeConnectLink(charger)"> Smart-Me &#8599;</a>

          <label for="smartmeSerialNumber">smart-me serial number</label>
          The Serial is the part before the "-". <br />
          <input
            v-model="charger.technicalData.smartmeSerialNumber"
            type="text"
            id="smartmeSerialNumber"
            required
            placeholder="serial number"
          />
          <br />
        </div>

        <button type="submit">submit</button>
      </form>

      <button
        @click="deleteCharger"
        v-if="!isEqual(charger, defaultCharger)"
        class="delete"
      >
        delete charger
      </button>

      <ErrorBox :msg="errMsg" v-if="errMsg != ''"></ErrorBox>
      <SuccessBox :msg="sucMsg" v-if="sucMsg != ''"></SuccessBox>
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

main > div.split50 {
  padding: 2em;
  width: calc(50vw - 4em);
  height: calc(100% - 4em);
}

div.textbox form > label:first-child {
  margin-top: 0;
}

@media screen and (orientation: portrait) and (max-width: 800px) {
  main {
    flex-direction: column;
  }
  main > div.split50 {
    width: calc(100vw - 4em);
    height: calc(50% - 4em);
  }
}

.textbox {
  overflow-y: auto;
}

label {
  font-size: 1.2em;
  font-weight: bold;
  display: block;
  margin-top: 1em;
  margin-bottom: 1em;
}

button.delete,
button[type="submit"] {
  margin-top: 2em;
}

.ErrorBox,
.SuccessBox {
  margin-top: 2em;
}
</style>
