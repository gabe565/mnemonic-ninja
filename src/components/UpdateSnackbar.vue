<template>
  <v-snackbar
    :value="offlineReady || needRefresh"
    :timeout="needRefresh ? -1 : 5000"
    bottom
    class="pb-14 pb-md-0"
  >
    <span v-if="needRefresh"> New content available, click on reload button to update. </span>
    <span v-else> App ready to work offline </span>
    <template #action="{ attrs }">
      <v-btn
        v-if="needRefresh"
        v-bind="attrs"
        text
        color="blue lighten-2"
        :loading="loading"
        @click.native="updateServiceWorker"
      >
        Refresh
      </v-btn>
      <v-btn icon aria-label="Close" @click="closePromptUpdateSW">
        <v-icon>$close</v-icon>
      </v-btn>
    </template>
  </v-snackbar>
</template>

<script>
import useRegisterSW from "../mixins/useRegisterSW";

const intervalMS = 60 * 60 * 1000;

export default {
  name: "UpdateSnackbar",
  mixins: [useRegisterSW],
  methods: {
    handleSWManualUpdates(r) {
      if (r) {
        setInterval(() => r.update(), intervalMS);
      }
    },
  },
};
</script>
