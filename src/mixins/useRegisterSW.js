/* eslint-disable no-unused-vars */
export default {
  name: "useRegisterSW",
  data() {
    return {
      updateSW: undefined,
      offlineReady: false,
      needRefresh: false,
      loading: false,
    };
  },
  async mounted() {
    try {
      const { registerSW } = await import("virtual:pwa-register");
      const vm = this;
      this.updateSW = registerSW({
        immediate: true,
        onOfflineReady() {
          vm.offlineReady = true;
          vm.onOfflineReadyFn();
        },
        onNeedRefresh() {
          vm.needRefresh = true;
          vm.onNeedRefreshFn();
        },
        onRegistered(swRegistration) {
          if (swRegistration) {
            vm.handleSWManualUpdates(swRegistration);
          }
        },
        onRegisterError(e) {
          vm.handleSWRegisterError(e);
        },
      });
    } catch {
      console.log("PWA disabled.");
    }
  },
  methods: {
    async closePromptUpdateSW() {
      this.offlineReady = false;
      this.needRefresh = false;
    },
    async onOfflineReadyFn() {
      console.log("onOfflineReady");
    },
    onNeedRefreshFn() {
      console.log("onNeedRefresh");
    },
    updateServiceWorker() {
      if (this.updateSW) {
        this.loading = true;
        this.updateSW(true);
      }
    },
    handleSWManualUpdates(swRegistration) {},
    handleSWRegisterError(error) {
      console.error(error);
    },
  },
};
