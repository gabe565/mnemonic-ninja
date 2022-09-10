<template>
  <v-snackbar
    :value="offlineReady || needRefresh"
    :timeout="needRefresh ? -1 : 5000"
    bottom
    class="pb-14 pb-md-0"
    @input="close"
  >
    <span v-if="needRefresh">
      New content available, click on reload button to update.
    </span>
    <span v-else>
      App ready to work offline
    </span>
    <template #action="{ attrs }">
      <v-btn
        v-if="needRefresh"
        v-bind="attrs"
        text
        color="primary"
        @click.native="updateServiceWorker"
      >
        Refresh
      </v-btn>
      <v-btn
        icon
        @click="closePromptUpdateSW"
      >
        <v-icon>$close</v-icon>
      </v-btn>
    </template>
  </v-snackbar>
</template>

<script>
import useRegisterSW from '../mixins/useRegisterSW';

const intervalMS = 60 * 60 * 1000;

export default {
  name: 'UpdateSnackbar',
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
