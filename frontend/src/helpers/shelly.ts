import { config } from "@/config";

export function getShellyConnectLink(): string {
  const baseURL = "https://my.shelly.cloud/integrator.html";
  const url = new URL(baseURL);
  url.searchParams.set("itg", config.SHELLY_INTEGRATOR_TAG);
  url.searchParams.set("cb", `${config.API_URL}/shelly/callback`);
  return url.toString();
}
