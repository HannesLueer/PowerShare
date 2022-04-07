import { config } from "@/config";
import { ref } from "vue";

const isLoggedin = ref<boolean>(false);

function localStorageSaveUser(user: string) {
  localStorage.setItem("user", user);
  isLoggedin.value = true;
}

function localStorageDeleteUser() {
  localStorage.removeItem("user");
  isLoggedin.value = false;
}

function login(email: string, password: string) {
  const requestOptions = {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ email, password }),
  };

  return fetch(`${config.API_URL}/user/signin`, requestOptions)
    .then(handleResponse)
    .then((user) => {
      // login successful if there's a jwt token in the response
      if (user.token) {
        // store user details and jwt token in local storage to keep user logged in between page refreshes
        localStorageSaveUser(JSON.stringify(user));
      }

      return user;
    });
}

function logout() {
  // remove user from local storage to log user out
  localStorageDeleteUser();
}

function register(name: string, email: string, password: string) {
  const requestOptions = {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ name, email, password }),
  };

  return fetch(`${config.API_URL}/user/signup`, requestOptions)
    .then(handleResponse)
    .then((user) => {
      // login successful if there's a jwt token in the response
      if (user.token) {
        // store user details and jwt token in local storage to keep user logged in between page refreshes
        localStorageSaveUser(JSON.stringify(user));
      }

      return user;
    });
}

function update(name: string, email: string, password: string) {
  const requestOptions = {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ name, email, password }),
  };

  return fetch(`${config.API_URL}/user/`, requestOptions)
    .then(handleResponse)
    .then((user) => {
      // login successful if there's a jwt token in the response
      if (user.token) {
        // store user details and jwt token in local storage to keep user logged in between page refreshes
        localStorageSaveUser(JSON.stringify(user));
      }

      return user;
    });
}

function remove() {
  const requestOptions = {
    method: "DEL",
    headers: { "Content-Type": "application/json" },
  };

  return fetch(`${config.API_URL}/user/`, requestOptions).then(handleResponse);
}

function handleResponse(response: Response) {
  return response.text().then((text: string) => {
    if (response.ok) {
      return JSON.parse(text);
    } else {
      if (response.status === 401) {
        // auto logout if 401 response returned from api
        logout();
        location.reload();
      }
      return text || response.statusText;
    }
  });
}

export const userService = {
  login,
  logout,
  register,
  update,
  remove,
  isLoggedin,
};
