// return authorization header with jwt token
export function authHeader(): Record<string, string> {
  const user = JSON.parse(localStorage.getItem("user") ?? "");

  if (user && user.token) {
    return { Token: user.token };
  } else {
    return {};
  }
}
