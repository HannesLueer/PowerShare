import { config } from "@/config";
import { handleResponse } from "@/services";

async function getList(): Promise<Currency[]> {
  return fetch(`${config.API_URL}/currency/all`, {
    method: "GET",
  })
    .then(handleResponse)
    .catch((error) => {
      console.error(error);
      return [];
    });
}

export class Currency {
  abbreviation!: string;
  symbol!: string;
}

export const currencyService = {
  getList,
};
