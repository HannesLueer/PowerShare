<script setup lang="ts">
import { onMounted } from "vue";
import L, { map } from "leaflet";
import "leaflet.locatecontrol";
import "leaflet.offline";
import "leaflet.markercluster";
import { GeoSearchControl, OpenStreetMapProvider } from "leaflet-geosearch";

const props = defineProps<{
  center: L.LatLngExpression;
  useLocation?: boolean;
  useSearch?: boolean;
  markersURL?: string;
  markersUpdateIntervalSeconds?: number;
}>();

let mapDiv: L.Map;

// methods
function setupLeafletMap() {
  mapDiv = L.map("mapContainer").setView(props.center, 6);
  L.tileLayer(
    // "https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png",
    "http://{s}.google.com/vt/lyrs=m&x={x}&y={y}&z={z}",
    {
      maxZoom: 22,
      subdomains: ["mt0", "mt1", "mt2", "mt3"],
    }
  ).addTo(mapDiv);
}

function setupLocation() {
  L.control
    .locate({
      position: "topleft",
      strings: {
        title: "Show location",
      },
      flyTo: true,
      keepCurrentZoomLevel: true,
      clickBehavior: {
        inView: "setView",
        outOfView: "setView",
        inViewNotFollowing: "inView",
      },
    })
    .addTo(mapDiv);
}

function setupSearch() {
  const provider = new OpenStreetMapProvider();

  const searchControl: L.Control = GeoSearchControl({
    provider: provider,
    position: "topright",
    style: "bar",
    autoCompleteDelay: 200,
    showMarker: false,
    searchLabel: "Enter location",
  });

  mapDiv.addControl(searchControl);
}

async function setupMarkers(url: string) {
  const timeout = (props.markersUpdateIntervalSeconds ?? (5 * 60 * 60)) * 1000;

  class MarkerData {
    id!: number;
    title!: string;
    position!: { Lat: number; Lng: number };
    cost!: number;
    isOccupied!: boolean;
  }

  let markersData: MarkerData[];
  let markersLayer = L.markerClusterGroup({ chunkedLoading: true });

  mapDiv.addLayer(markersLayer);

  updateMarkers(url);
  setInterval(updateMarkers, timeout, url);

  async function updateMarkers(url: string) {
    let tmpMarkers = await getMarkerLocations(url);
    if (tmpMarkers) {
      markersData = tmpMarkers;
      deleteAllMarkers();
      drawMarkers(markersData);
    }
  }

  function deleteAllMarkers() {
    markersLayer.clearLayers();
  }

  async function getMarkerLocations(url: string) {
    let response = await fetch(url);
    if (response.ok) {
      return await response.json();
    } else {
      console.log("HTTP-Error: " + response.status);
      return undefined;
    }
  }

  function drawMarkers(markersData: MarkerData[]) {
    let markers: L.Layer[] = [];
    markersData.forEach((m) => {
      let marker = L.marker([m.position.Lat, m.position.Lng]).bindPopup(
        '<a href="/location/' +
          m.id +
          '" target="_blank" rel="noopener">' +
          m.title +
          "</a>"
      );
      markers.push(marker);
    });
    markersLayer.addLayers(markers)
  }
}

onMounted(async () => {
  setupLeafletMap();
  if (props.useLocation) setupLocation();
  if (props.useSearch) setupSearch();
  if (props.markersURL) setupMarkers(props.markersURL);
});
</script>

<template>
  <div id="mapContainer"></div>
</template>

<style>
@import "leaflet.locatecontrol/dist/L.Control.Locate.min.css";
@import "leaflet/dist/leaflet.css";
@import "leaflet-geosearch/dist/geosearch.css";
@import "leaflet.markercluster/dist/MarkerCluster.Default.css";
@import "leaflet.markercluster/dist/MarkerCluster.css";

/* map */
.leaflet-container {
  background: rgb(163, 163, 163);
}

:root[theme="dark"] .leaflet-container {
  background: rgb(73, 73, 73);
}

:root[theme="dark"] div#mapContainer .leaflet-tile {
  /* mapbox */
  /* filter: brightness(0.6) invert(1) contrast(3) hue-rotate(200deg) saturate(0.5) brightness(0.6); */

  /* google maps */
  filter: invert(0.8) hue-rotate(175deg) saturate(0.8) contrast(1.05) brightness(0.9);
}

/* location */
:root[theme="dark"] .leaflet-control-container a[href="#"] {
  background-color: var(--color-button-background);
  color: var(--color-button-text);
}

/* attribution */
:root[theme="dark"] .leaflet-control-attribution {
  background-color: var(--color-background);
  opacity: 0.6;
  color: var(--color-text);
}

.leaflet-control-attribution a {
  color: var(--color-link);
}

/* search */
:root[theme="dark"] .leaflet-control-geosearch form,
:root[theme="dark"] .leaflet-control-geosearch input,
:root[theme="dark"] .leaflet-control-geosearch .results {
  background-color: var(--color-button-background);
}

:root[theme="dark"] .leaflet-control-geosearch .results.active {
  color: var(--color-button-text);
  border-top: 1px solid var(--color-button-border);
}

:root[theme="dark"] .leaflet-control-geosearch .results div:hover {
  background-color: var(--color-button-background-hover);
}

:root[theme="dark"] .leaflet-control-geosearch input::placeholder {
  color: var(--color-button-text);
}
</style>

<style scoped>
div {
  width: 100%;
  height: 100%;
}
</style>