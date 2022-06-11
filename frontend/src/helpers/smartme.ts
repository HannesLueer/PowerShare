import { config } from "@/config";
import type { ChargerData } from "@/services";

export function getSmartMeConnectLink(charger: ChargerData): string {
  const baseURL = "https://smart-me.com/api/oauth/authorize";
  const url = new URL(baseURL);
  url.searchParams.set("client_id", config.SMARTME_CLIENT_ID);
  url.searchParams.set("response_type", "code");
  url.searchParams.set("redirect_uri", window.location.href);
  url.searchParams.set("scope", "device.read");
  url.searchParams.set("state", JSON.stringify(charger));
  return url.toString();
}
