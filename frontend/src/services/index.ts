import { userService } from "@/services";

export * from "./chargerService";
export * from "./chargingService";
export * from "./currencyService";
export * from "./gocardlessService";
export * from "./smartmeService";
export * from "./userService";

export function handleResponse(response: Response) {
  return response.text().then((text: string) => {
    if (
      response.ok &&
      response.headers.get("content-type") == "application/json"
    ) {
      return JSON.parse(text);
    } else {
      if (response.status === 401) {
        // auto logout if 401 response returned from api
        userService.logout();
      }
      return text || response.statusText;
    }
  });
}
