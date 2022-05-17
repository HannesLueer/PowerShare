<script setup lang="ts">
import { config } from "@/config";
import { chargingService } from "@/services";
import {
  type OnApproveData,
  loadScript,
  type OnApproveActions,
} from "@paypal/paypal-js";

const props = defineProps<{
  onApproveFunc: (
    data: OnApproveData,
    actions: OnApproveActions
  ) => Promise<void>;
}>();

loadScript({
  "client-id": config.PAYPAL_CLIENT_ID,
  intent: "authorize",
  // intent: "capture",
  // commit: false,
  // vault: true,
})
  .then((paypal) => {
    if (paypal == null) return;
    paypal
      .Buttons({
        // Sets up the transaction when a payment button is clicked
        createOrder: (data, actions) => {
          // return actions.order.create({
          //   purchase_units: [
          //     {
          //       amount: {
          //         value: "10.00",
          //       },
          //     },
          //   ],
          // });

          return chargingService
            .newOrder(1) // TODO: parameter: chargerID
            .then((response) => response.json())
            .then((order) => order.id);
        },
        style: {
          color: "black",
        },
        // Finalize the transaction after payer approval
        onApprove: props.onApproveFunc,
      })
      .render("#paypal-button-container")
      .catch((error) => {
        console.error("failed to render the PayPal Buttons", error);
      });
  })
  .catch((error) => {
    console.error("failed to load the PayPal JS SDK script", error);
  });
</script>

<template>
  <div class="PayPal">
    <div id="paypal-button-container"></div>
  </div>
</template>

<style scoped>
#paypal-button-container {
  filter: invert(1) hue-rotate(180deg) brightness(0.9);
}

:root[theme="dark"] #paypal-button-container {
  filter: invert(1) hue-rotate(180deg);
}
</style>
