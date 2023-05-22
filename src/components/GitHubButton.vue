<template>
  <v-tooltip location="bottom">
    <template #activator="{ props }">
      <v-btn
        icon
        href="https://github.com/gabe565/mnemonic-ninja"
        target="_blank"
        rel="noopener"
        v-bind="props"
        aria-label="View on GitHub"
      >
        <v-icon :icon="GitHubIcon" style="fill: currentColor" />
      </v-btn>
    </template>
    <span>View on GitHub</span>
  </v-tooltip>
</template>

<script setup>
import GitHubIcon from "simple-icons/icons/github.svg";
</script>
