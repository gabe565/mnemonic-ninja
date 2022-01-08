<template>
  <v-snackbar
    v-model="show"
    :timeout="timeout"
    bottom right
  >
    New version available!
    <template #action="{ attrs }">
      <v-btn v-bind="attrs" @click.native="refreshApp" text color="primary">
        Refresh
      </v-btn>
      <v-btn icon @click="show = false">
        <v-icon>$close</v-icon>
      </v-btn>
    </template>
  </v-snackbar>
</template>

<script>
import { wait } from '@/util/helpers';

export default {
  name: 'UpdateSnackbar',

  data: () => ({
    refreshing: false,
    registration: null,
    show: false,
    timeout: -1,
  }),

  created() {
    // Listen for swUpdated event and display refresh snackbar as required.
    document.addEventListener('swUpdated', this.showRefreshUI, { once: true });
    // Refresh all open app tabs when a new service worker is installed.
    navigator.serviceWorker.addEventListener('controllerchange', () => {
      if (this.refreshing) return;
      this.refreshing = true;
      window.location.reload();
    });
  },

  methods: {
    async showRefreshUI(e) {
      // Display a snackbar inviting the user to refresh/reload the app due
      // to an app update being available.
      // The new service worker is installed, but not yet active.
      // Store the ServiceWorkerRegistration instance for later use.
      this.registration = e.detail;
      await wait(1000);
      this.show = true;
    },

    refreshApp() {
      this.show = false;
      // Protect against missing registration.waiting.
      if (!this.registration || !this.registration.waiting) { return; }
      this.registration.waiting.postMessage('skipWaiting');
    },
  },
};
</script>
