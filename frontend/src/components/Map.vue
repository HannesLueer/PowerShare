<script setup lang="ts">
import { onMounted, watch, toRefs } from "vue";
import L, { map } from "leaflet";
import "leaflet.locatecontrol";
import "leaflet.offline";
import "leaflet.markercluster";
import { GeoSearchControl, OpenStreetMapProvider } from "leaflet-geosearch";
import "leaflet-easybutton";
import router from "@/router";
import { chargerService, type ChargerData, type Coordinate } from "@/services";

const emit = defineEmits<{
  (e: "clickedPosition", pos: Coordinate): void;
}>();

const props = defineProps<{
  center: L.LatLngExpression;
  useLocation?: boolean;
  useSearch?: boolean;
  markersApiPath?: string;
  markersUpdateIntervalSeconds?: number;
  markerLinkTo?: string;
  useAddButton?: boolean;
  useManualUpdateButton?: boolean;
  triggerMarkerUpdate?: number;
  triggerMapInvalidateSize?: number;
}>();

const propsRef = toRefs(props);

let mapDiv: L.Map;
let markersLayer = L.markerClusterGroup({ chunkedLoading: true });

const markerImage =
  "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABkAAAApCAYAAADAk4LOAAAFgUlEQVR4Aa1XA5BjWRTN2oW17d3YaZtr2962HUzbDNpjszW24mRt28p47v7zq/bXZtrp/lWnXr337j3nPCe85NcypgSFdugCpW5YoDAMRaIMqRi6aKq5E3YqDQO3qAwjVWrD8Ncq/RBpykd8oZUb/kaJutow8r1aP9II0WmLKLIsJyv1w/kqw9Ch2MYdB++12Onxee/QMwvf4/Dk/Lfp/i4nxTXtOoQ4pW5Aj7wpici1A9erdAN2OH64x8OSP9j3Ft3b7aWkTg/Fm91siTra0f9on5sQr9INejH6CUUUpavjFNq1B+Oadhxmnfa8RfEmN8VNAsQhPqF55xHkMzz3jSmChWU6f7/XZKNH+9+hBLOHYozuKQPxyMPUKkrX/K0uWnfFaJGS1QPRtZsOPtr3NsW0uyh6NNCOkU3Yz+bXbT3I8G3xE5EXLXtCXbbqwCO9zPQYPRTZ5vIDXD7U+w7rFDEoUUf7ibHIR4y6bLVPXrz8JVZEql13trxwue/uDivd3fkWRbS6/IA2bID4uk0UpF1N8qLlbBlXs4Ee7HLTfV1j54APvODnSfOWBqtKVvjgLKzF5YdEk5ewRkGlK0i33Eofffc7HT56jD7/6U+qH3Cx7SBLNntH5YIPvODnyfIXZYRVDPqgHtLs5ABHD3YzLuespb7t79FY34DjMwrVrcTuwlT55YMPvOBnRrJ4VXTdNnYug5ucHLBjEpt30701A3Ts+HEa73u6dT3FNWwflY86eMHPk+Yu+i6pzUpRrW7SNDg5JHR4KapmM5Wv2E8Tfcb1HoqqHMHU+uWDD7zg54mz5/2BSnizi9T1Dg4QQXLToGNCkb6tb1NU+QAlGr1++eADrzhn/u8Q2YZhQVlZ5+CAOtqfbhmaUCS1ezNFVm2imDbPmPng5wmz+gwh+oHDce0eUtQ6OGDIyR0uUhUsoO3vfDmmgOezH0mZN59x7MBi++WDL1g/eEiU3avlidO671bkLfwbw5XV2P8Pzo0ydy4t2/0eu33xYSOMOD8hTf4CrBtGMSoXfPLchX+J0ruSePw3LZeK0juPJbYzrhkH0io7B3k164hiGvawhOKMLkrQLyVpZg8rHFW7E2uHOL888IBPlNZ1FPzstSJM694fWr6RwpvcJK60+0HCILTBzZLFNdtAzJaohze60T8qBzyh5ZuOg5e7uwQppofEmf2++DYvmySqGBuKaicF1blQjhuHdvCIMvp8whTTfZzI7RldpwtSzL+F1+wkdZ2TBOW2gIF88PBTzD/gpeREAMEbxnJcaJHNHrpzji0gQCS6hdkEeYt9DF/2qPcEC8RM28Hwmr3sdNyht00byAut2k3gufWNtgtOEOFGUwcXWNDbdNbpgBGxEvKkOQsxivJx33iow0Vw5S6SVTrpVq11ysA2Rp7gTfPfktc6zhtXBBC+adRLshf6sG2RfHPZ5EAc4sVZ83yCN00Fk/4kggu40ZTvIEm5g24qtU4KjBrx/BTTH8ifVASAG7gKrnWxJDcU7x8X6Ecczhm3o6YicvsLXWfh3Ch1W0k8x0nXF+0fFxgt4phz8QvypiwCCFKMqXCnqXExjq10beH+UUA7+nG6mdG/Pu0f3LgFcGrl2s0kNNjpmoJ9o4B29CMO8dMT4Q5ox8uitF6fqsrJOr8qnwNbRzv6hSnG5wP+64C7h9lp30hKNtKdWjtdkbuPA19nJ7Tz3zR/ibgARbhb4AlhavcBebmTHcFl2fvYEnW0ox9xMxKBS8btJ+KiEbq9zA4RthQXDhPa0T9TEe69gWupwc6uBUphquXgf+/FrIjweHQS4/pduMe5ERUMHUd9xv8ZR98CxkS4F2n3EUrUZ10EYNw7BWm9x1GiPssi3GgiGRDKWRYZfXlON+dfNbM+GgIwYdwAAAAASUVORK5CYII=";

const markerShadowImage =
  "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACkAAAApCAQAAAACach9AAACMUlEQVR4Ae3ShY7jQBAE0Aoz/f9/HTMzhg1zrdKUrJbdx+Kd2nD8VNudfsL/Th///dyQN2TH6f3y/BGpC379rV+S+qqetBOxImNQXL8JCAr2V4iMQXHGNJxeCfZXhSRBcQMfvkOWUdtfzlLgAENmZDcmo2TVmt8OSM2eXxBp3DjHSMFutqS7SbmemzBiR+xpKCNUIRkdkkYxhAkyGoBvyQFEJEefwSmmvBfJuJ6aKqKWnAkvGZOaZXTUgFqYULWNSHUckZuR1HIIimUExutRxwzOLROIG4vKmCKQt364mIlhSyzAf1m9lHZHJZrlAOMMztRRiKimp/rpdJDc9Awry5xTZCte7FHtuS8wJgeYGrex28xNTd086Dik7vUMscQOa8y4DoGtCCSkAKlNwpgNtphjrC6MIHUkR6YWxxs6Sc5xqn222mmCRFzIt8lEdKx+ikCtg91qS2WpwVfBelJCiQJwvzixfI9cxZQWgiSJelKnwBElKYtDOb2MFbhmUigbReQBV0Cg4+qMXSxXSyGUn4UbF8l+7qdSGnTC0XLCmahIgUHLhLOhpVCtw4CzYXvLQWQbJNmxoCsOKAxSgBJno75avolkRw8iIAFcsdc02e9iyCd8tHwmeSSoKTowIgvscSGZUOA7PuCN5b2BX9mQM7S0wYhMNU74zgsPBj3HU7wguAfnxxjFQGBE6pwN+GjME9zHY7zGp8wVxMShYX9NXvEWD3HbwJf4giO4CFIQxXScH1/TM+04kkBiAAAAAElFTkSuQmCC";

const markerIcon = new L.Icon({
  iconUrl: markerImage,
  iconAnchor: [12, 41],
  shadowUrl: markerShadowImage,
});

// methods
function setupLeafletMap() {
  mapDiv = L.map("mapContainer")
    .setView(props.center, 6)
    .on("click", function (ev: L.LeafletMouseEvent) {
      emit("clickedPosition", { Lat: ev.latlng.lat, Lng: ev.latlng.lng });
    });
  L.tileLayer(
    // "https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png",
    "https://{s}.google.com/vt/lyrs=m&x={x}&y={y}&z={z}",
    {
      maxZoom: 22,
      subdomains: ["mt0", "mt1", "mt2", "mt3"],
    }
  ).addTo(mapDiv);

  watch(propsRef.center, async () => {
    mapDiv.setView(props.center);
  });
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

async function setupMarkers(apiPath: string) {
  const timeout = (props.markersUpdateIntervalSeconds ?? 5 * 60 * 60) * 1000;

  mapDiv.addLayer(markersLayer);

  updateMarkers(apiPath);
  setInterval(updateMarkers, timeout, apiPath);
}

async function updateMarkers(apiPath: string) {
  let markersData: ChargerData[];
  let tmpMarkers = await chargerService.getList(apiPath);
  if (tmpMarkers != []) {
    markersData = tmpMarkers;
    deleteAllMarkers();
    drawMarkers(markersData);
  }
}

function deleteAllMarkers() {
  markersLayer.clearLayers();
}

function drawMarkers(chargerArray: ChargerData[]) {
  let markers: L.Layer[] = [];
  chargerArray.forEach((m) => {
    let marker = L.marker([m.position.Lat, m.position.Lng], {
      icon: markerIcon,
    }).on("click", () => {
      if (props.markerLinkTo) router.push(props.markerLinkTo + m.id);
    });
    markers.push(marker);
  });
  markersLayer.addLayers(markers);
}

async function setupAddButton() {
  L.easyButton(
    '<img src="/src/assets/img/new-location.svg" width="22">',
    function (btn, map) {
      router.push("/mycharger/new");
    }
  ).addTo(mapDiv);
}

async function setupManualUpdateButton(apiPath: string) {
  L.easyButton(
    '<img src="/src/assets/img/reload.svg" width="22">',
    function (btn, map) {
      updateMarkers(apiPath);
    }
  ).addTo(mapDiv);
}

onMounted(async () => {
  setupLeafletMap();
  if (props.useLocation) setupLocation();
  if (props.useSearch) setupSearch();
  if (props.markersApiPath) setupMarkers(props.markersApiPath);
  if (props.useAddButton) setupAddButton();
  if (props.useManualUpdateButton && props.markersApiPath)
    setupManualUpdateButton(props.markersApiPath);

  if (propsRef.triggerMarkerUpdate) {
    watch(propsRef.triggerMarkerUpdate, async () => {
      if (props.markersApiPath) updateMarkers(props.markersApiPath);
    });
  }

  if (propsRef.triggerMapInvalidateSize) {
    watch(propsRef.triggerMapInvalidateSize, async () => {
      setTimeout(() => {
        mapDiv.invalidateSize();
      }, 300);
    });
  }
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
  filter: invert(0.8) hue-rotate(175deg) saturate(0.8) contrast(1.05)
    brightness(0.9);
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

/* leaflet-easybutton */
.easy-button-button {
  padding: 4px;
  border: none;
  background-color: white;
}

.easy-button-button:hover {
  background-color: #f4f4f4;
}

:root[theme="dark"] .easy-button-button {
  background-color: var(--color-button-background);
}
</style>

<style scoped>
div {
  width: 100%;
  height: 100%;
}
</style>
