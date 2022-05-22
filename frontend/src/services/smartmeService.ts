import { config } from "@/config";
import { authHeader } from "@/helpers";

function post_authCode(code: string) {
  const requestOptions = {
    method: "POST",
    headers: Object.assign(
      { "Content-Type": "application/json" },
      authHeader()
    ),
  };

  return fetch(`${config.API_URL}/authcode/${code}`, requestOptions)
    .then((response) => {
      if (!response.ok) console.error(response.body);
    })
    .catch((error) => {
      console.error(error);
    });
}

export const smartmeService = {
  post_authCode,
};
