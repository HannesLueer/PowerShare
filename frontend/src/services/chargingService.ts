import { config } from "@/config";
import { authHeader } from "@/helpers";
import { handleResponse } from "@/services";

function start(chargerID: number, paypalOrderID: string): Promise<void> {
  const requestOptions = {
    method: "POST",
    headers: Object.assign(
      { "Content-Type": "application/json" },
      authHeader()
    ),
  };

  return fetch(
    `${config.API_URL}/charging/start/${chargerID}/${paypalOrderID}`,
    requestOptions
  )
    .then(handleResponse)
    .then((charger) => {
      return charger;
    })
    .catch((error) => {
      console.error(error);
      return "a communication error occurred";
    });
}

function stop(chargerID: number, paypalOrderID: string): Promise<void> {
  const requestOptions = {
    method: "POST",
    headers: Object.assign(
      { "Content-Type": "application/json" },
      authHeader()
    ),
  };

  return fetch(
    `${config.API_URL}/charging/stop/${chargerID}/${paypalOrderID}`,
    requestOptions
  )
    .then(handleResponse)
    .then((charger) => {
      return charger;
    })
    .catch((error) => {
      console.error(error);
      return "a communication error occurred";
    });
}

export const chargingService = {
  start,
  stop,
};
