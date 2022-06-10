import { config } from "@/config";
import { authHeader } from "@/helpers";
import { userService } from "@/services";

function start(chargerID: number): Promise<string> {
  const requestOptions = {
    method: "POST",
    headers: Object.assign(
      { "Content-Type": "application/json" },
      authHeader()
    ),
  };

  return fetch(`${config.API_URL}/charging/start/${chargerID}`, requestOptions)
    .then((response) => {
      if (!response.ok) {
        console.error(response.statusText);
      }
      return response.text();
    })
    .catch((error) => {
      console.error(error);
      return "a communication error occurred";
    });
}

function stop(chargerID: number): Promise<string> {
  const requestOptions = {
    method: "POST",
    headers: Object.assign(
      { "Content-Type": "application/json" },
      authHeader()
    ),
  };

  return fetch(`${config.API_URL}/charging/stop/${chargerID}`, requestOptions)
    .then((response) => {
      if (!response.ok) {
        console.error(response.statusText);
      }
      return response.text();
    })
    .catch((error) => {
      console.error(error);
      return "a communication error occurred";
    });
}

async function isThisUserCharging(chargerID: number): Promise<string> {
  if (!userService.isLoggedin.value) {
    return "false";
  }

  const requestOptions = {
    method: "GET",
    headers: Object.assign(authHeader()),
  };

  return fetch(`${config.API_URL}/charging/is/${chargerID}`, requestOptions)
    .then((response) => {
      return response.text();
    })
    .catch((error) => {
      console.error(error);
      return "a communication error occurred";
    });
}

export const chargingService = {
  start,
  stop,
  isThisUserCharging,
};
