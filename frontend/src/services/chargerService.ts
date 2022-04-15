import { config } from "@/config";
import { authHeader } from "@/helpers";

function get(id: number): Promise<ChargerData> {
  const requestOptions = {
    method: "GET",
    headers: Object.assign(
      { "Content-Type": "application/json" },
      authHeader()
    ),
  };

  return fetch(`${config.API_URL}/charger/${id}`, requestOptions)
    .then(handleResponse)
    .then((charger) => {
      return charger;
    })
    .catch((error) => {
      console.error(error);
      return "a communication error occurred";
    });
}

function create(charger: ChargerData) {
  const requestOptions = {
    method: "POST",
    headers: Object.assign(
      { "Content-Type": "application/json" },
      authHeader()
    ),
    body: JSON.stringify(charger),
  };

  return fetch(`${config.API_URL}/charger/`, requestOptions)
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

function update(charger: ChargerData) {
  const requestOptions = {
    method: "PUT",
    headers: Object.assign(
      { "Content-Type": "application/json" },
      authHeader()
    ),
    body: JSON.stringify(charger),
  };

  return fetch(`${config.API_URL}/charger/`, requestOptions)
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

function remove(id: number) {
  const requestOptions = {
    method: "DELETE",
    headers: authHeader(),
  };

  return fetch(`${config.API_URL}/charger/${id}`, requestOptions)
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

function handleResponse(response: Response) {
  return response.text().then((text: string) => {
    if (
      response.ok &&
      response.headers.get("content-type") == "application/json"
    ) {
      return JSON.parse(text);
    } else {
      return text || response.statusText;
    }
  });
}

export class ChargerData {
  id!: number;
  title!: string;
  position!: Coordinate;
  cost!: number;
  isOccupied!: boolean;
}

export class Coordinate {
  Lat!: number;
  Lng!: number;
}

export const chargerService = {
  get,
  create,
  update,
  remove,
};
