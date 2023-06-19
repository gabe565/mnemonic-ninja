<template>
  <v-tooltip location="bottom" aria-label="View on GitHub">
    <template #activator="{ props }">
      <v-btn
        icon
        href="https://github.com/gabe565/mnemonic-ninja"
        target="_blank"
        rel="noopener"
        v-bind="props"
        aria-label="View on GitHub"
      >
        <v-icon :icon="GitHubIcon" />
      </v-btn>
    </template>
    <span>View on GitHub</span>
  </v-tooltip>
</template>

<script setup>
import GitHubIcon from "~icons/simple-icons/github";
</script>
