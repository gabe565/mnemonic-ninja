<template>
  <v-tooltip location="top" text="View on GitHub">
    <template #activator="{ props }">
      <v-btn
        v-bind="props"
        href="https://github.com/gabe565/mnemonic-ninja"
        target="_blank"
        rel="noopener"
        color="transparent"
        variant="flat"
        aria-label="View on GitHub"
        :icon="GitHubIcon"
      />
    </template>
    <span>View on GitHub</span>
  </v-tooltip>
</template>

<script setup>
import GitHubIcon from "~icons/simple-icons/github";
</script>
