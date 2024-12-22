<template>
  <PageLayout>
    <v-row v-if="error">
      <v-col>
        <v-alert type="error">{{ error }}</v-alert>
      </v-col>
    </v-row>

    <v-row>
      <v-col>
        <slot />
      </v-col>
    </v-row>
  </PageLayout>
</template>

<script setup>
import { error } from "../data/wordlist.js";
import PageLayout from "./PageLayout.vue";
</script>
