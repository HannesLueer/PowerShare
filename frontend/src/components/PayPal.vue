<script setup lang="ts">
import { config } from "@/config";
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
  commit: false,
})
  .then((paypal) => {
    if (paypal == null) return;
    paypal
      .Buttons({
        // // Sets up the transaction when a payment button is clicked
        // createOrder: (data, actions) => {
        //   return actions.order.create({
        //     purchase_units: [
        //       {
        //         amount: {
        //           value: "1.00", // Can also reference a variable or function
        //         },
        //       },
        //     ],
        //   });
        // },
        // Finalize the transaction after payer approval
        style: {
          color: "black",
        },
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
#paypal-button-container input {
  background-color: red !important;
}
</style>
