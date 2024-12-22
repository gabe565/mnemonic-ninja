<template>
  <v-snackbar
    :model-value="offlineReady || needRefresh"
    :timeout="needRefresh ? -1 : 5000"
    location="bottom"
    class="mb-16 mb-md-2"
    theme="light"
  >
    <span v-if="needRefresh"> New content available, click on reload button to update. </span>
    <span v-else> App ready to work offline </span>
    <template #actions>
      <v-btn
        v-if="needRefresh"
        variant="text"
        color="blue-lighten-2"
        :loading="loading"
        @click="updateServiceWorker"
      >
        Refresh
      </v-btn>
      <v-btn icon aria-label="Close" @click="closePromptUpdateSW">
        <v-icon>$close</v-icon>
      </v-btn>
    </template>
  </v-snackbar>
</template>

<script setup>
import { onMounted, ref } from "vue";

const updateSW = ref(undefined);
const offlineReady = ref(false);
const needRefresh = ref(false);
const loading = ref(false);

const intervalMS = 60 * 60 * 1000;

onMounted(async () => {
  try {
    const { registerSW } = await import("virtual:pwa-register");
    updateSW.value = registerSW({
      immediate: true,
      onOfflineReady() {
        offlineReady.value = true;
        console.log("onOfflineReady");
      },
      onNeedRefresh() {
        needRefresh.value = true;
        console.log("onNeedRefresh");
      },
      onRegistered(swRegistration) {
        if (swRegistration) {
          setInterval(() => swRegistration.update(), intervalMS);
        }
      },
      onRegisterError(error) {
        console.error(error);
      },
    });
  } catch {
    console.log("PWA disabled.");
  }
});

const closePromptUpdateSW = async () => {
  offlineReady.value = false;
  needRefresh.value = false;
};

const updateServiceWorker = () => {
  if (updateSW.value) {
    loading.value = true;
    updateSW.value(true);
  }
};
</script>
