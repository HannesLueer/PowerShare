<script setup lang="ts">
import { ref } from "vue";
import DefaultLayout from "../layouts/BasicLayout.vue";
import ErrorBox from "@/components/ErrorBox.vue";
import SuccessBox from "@/components/SuccessBox.vue";
import PayPal from "@/components/PayPal.vue";
import type { OnApproveData, OnApproveActions } from "@paypal/paypal-js";
import { chargingService } from "@/services";
import { useRoute } from "vue-router";

let errMsg = ref<string>("");
let sucMsg = ref<string>("");

const route = useRoute();

function onApprove(data: OnApproveData, actions: OnApproveActions) {
  // Authorize the transaction
  return actions.order.authorize().then(function (authorization) {
    // Get the authorization id
    var authorizationID: string = authorization.purchase_units[0].payments
      .authorizations[0].id as string;

    // Call server to capture the transaction
    if (typeof route.params.id == "string")
      chargingService.start(parseInt(route.params.id), data.orderID);

    // Optional message given to user
    sucMsg.value =
      "You have authorized this transaction. Order ID:  " +
      data.orderID +
      ", Authorization ID: " +
      authorizationID;
  });
}
</script>

<template>
  <DefaultLayout>
    <h1>Charging</h1>

    Please enter your payment information using the buttons below, then the
    wallbox will be switched on.

    <PayPal :onApproveFunc="onApprove"></PayPal>

    <ErrorBox :msg="errMsg" v-if="errMsg != ''"></ErrorBox>
    <SuccessBox :msg="sucMsg" v-if="sucMsg != ''"></SuccessBox>
  </DefaultLayout>
</template>

<style scoped>
:deep() .PayPal {
  margin-top: 1em;
}
</style>
