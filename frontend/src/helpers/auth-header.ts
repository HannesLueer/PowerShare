// return authorization header with jwt token
export function authHeader(): Record<string, string> {
  const userStr = localStorage.getItem("user");
  const user = userStr ? JSON.parse(userStr) : "";

  if (user && user.token) {
    return { Token: user.token };
  } else {
    return {};
  }
}
