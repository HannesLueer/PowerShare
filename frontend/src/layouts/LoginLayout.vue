<template>
  <main>
    <div class="login">
      <div class="image-wrapper">
        <img src="@/assets/img/user.svg" class="login" alt="User-Logo" />
      </div>

      <div class="form">
        <slot></slot>
      </div>
    </div>
  </main>
</template>

<style scoped>
main {
  display: flex;
  justify-content: center;
  align-items: center;
  overflow-y: auto;
}
div.login {
  background-color: var(--color-login-box-background);
  width: var(--wrapper-width);
  margin-top: 6rem;
  padding: var(--wrapper-padding);

  /* vars */
  --image-height: 12rem;
  --wrapper-padding: 40px;
  --wrapper-width: 33%;
  --image-wrapper-padding: 2.5rem;
}

@media (max-height: 800px) {
  main {
    align-items: flex-start;
  }

  div.login {
    margin-top: 9rem;
    margin-bottom: 3rem;
  }
}

@media (max-width: 900px) {
  div.login {
    /* vars */
    --wrapper-width: 66%;
  }
}

@media (max-width: 500px) {
  div.login {
    /* vars */
    --wrapper-padding: 20px;
    --wrapper-width: calc(100% - 2 * var(--wrapper-padding));
  }
}

div.image-wrapper {
  background-color: var(--color-login-image-background);
  border-radius: 50%;

  width: calc(var(--image-height) - 2 * var(--image-wrapper-padding));
  height: calc(var(--image-height) - 2 * var(--image-wrapper-padding));
  margin-top: calc(
    -1 * (var(--image-height) / 2 + var(--wrapper-padding))
  ); /* circle should overhang half out of the login wrapper */
  margin-left: auto;
  margin-right: auto;
  padding: var(--image-wrapper-padding);
  padding-bottom: 3rem;
  padding-top: 2rem;

  box-shadow: 0px 6px 30px -12px #000;
}

div.image-wrapper img {
  width: 100%;
  height: 100%;
}

.form :deep() h1 {
  font-weight: 100;
  text-align: center;
  font-size: 3em;
  margin: 1.5rem;
}

.form :deep() input {
  width: calc(100% - 2 * (0.3em + 2px));
  font-size: 1.4em;
  margin-bottom: 1em;
  border: 2px solid var(--color-login-input-border);
}

.form :deep() button {
  width: 100%;
  padding: 0.3em;
  font-size: 1.2em;
  font-size: 1.4em;
  margin-top: 1em;
  margin-bottom: 1em;
}

.form :deep() div.ErrorBox {
  margin-bottom: 1em;
}
</style>
