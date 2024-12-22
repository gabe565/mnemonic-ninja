<template>
  <v-app-bar color="primary" flat :title="$route.name">
    <template #prepend>
      <v-btn v-if="isMobile" to="/" icon size="small" :active="false">
        <v-icon :icon="AppIcon" size="28" style="opacity: 0.7" aria-label="Mnemonic Ninja" />
      </v-btn>
    </template>
    <template #append>
      <GitHubButton v-if="isMobile" />
    </template>
  </v-app-bar>

  <div class="d-flex align-center" style="min-height: 100%">
    <v-container class="pt-6 pt-lg-12">
      <v-row v-if="$slots.description" tag="section">
        <h2 class="d-sr-only">Description</h2>
        <v-col>
          <slot name="description" />
        </v-col>
      </v-row>

      <slot />
    </v-container>
  </div>
</template>

<script setup>
import AppIcon from "~icons/mnemonic-ninja/logo";
import { useDisplay } from "vuetify";
import GitHubButton from "../components/NavButtons/GitHubButton.vue";

const { smAndDown: isMobile } = useDisplay();
</script>
