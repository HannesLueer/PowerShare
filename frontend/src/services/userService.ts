import { config } from "@/config";
import { authHeader } from "@/helpers";
import { ref } from "vue";
import { handleResponse } from "@/services";

const isLoggedin = ref<boolean>(localStorageGetUser() != null);

function localStorageSaveUser(user: string) {
  localStorage.setItem("user", user);
  isLoggedin.value = true;
}

function localStorageDeleteUser() {
  localStorage.removeItem("user");
  isLoggedin.value = false;
}

function localStorageGetUser() {
  return localStorage.getItem("user");
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
    })
    .catch((error) => {
      console.error(error);
      return "a communication error occurred";
    });
}

function logout() {
  // remove user from local storage to log user out
  localStorageDeleteUser();
  location.reload();
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
    })
    .catch((error) => {
      console.error(error);
      return "a communication error occurred";
    });
}

function update(name: string, email: string, password: string) {
  const requestOptions = {
    method: "PUT",
    headers: Object.assign(
      { "Content-Type": "application/json" },
      authHeader()
    ),
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
    })
    .catch((error) => {
      console.error(error);
      return "a communication error occurred";
    });
}

function remove() {
  const requestOptions = {
    method: "DELETE",
    headers: authHeader(),
  };

  return fetch(`${config.API_URL}/user/`, requestOptions)
    .then((response) => {
      // delete successful if there's no text in response
      if (response.ok) {
        logout();
      }
      return response.text();
    })
    .catch((error) => {
      console.error(error);
      return "a communication error occurred";
    });
}

function get() {
  const requestOptions = {
    method: "GET",
    headers: Object.assign(
      { "Content-Type": "application/json" },
      authHeader()
    ),
  };

  return fetch(`${config.API_URL}/user/`, requestOptions)
    .then(handleResponse)
    .then((user) => {
      return user;
    })
    .catch((error) => {
      console.error(error);
      return "a communication error occurred";
    });
}

export const userService = {
  login,
  logout,
  register,
  update,
  remove,
  get,
  isLoggedin,
};
