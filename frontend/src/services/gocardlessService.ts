import { config } from "@/config";
import { authHeader } from "@/helpers";
import { userService } from "@/services";

async function get_newMandateURL(): Promise<string> {
  if (!userService.isLoggedin.value) {
    return "/login";
  }

  const requestOptions = {
    method: "GET",
    headers: Object.assign(
      { "Content-Type": "application/json" },
      authHeader()
    ),
  };

  return await fetch(`${config.API_URL}/gocardless/newMandate`, requestOptions)
    .then(async (response) => response.text())
    .catch((error) => {
      console.error(error);
      return "";
    });
}

export const gocardlessService = {
  get_newMandateURL,
};
