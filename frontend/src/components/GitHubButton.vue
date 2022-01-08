<template>
  <v-tooltip bottom>
    <template #activator="{ on, attrs }">
      <v-btn
        icon
        href="https://github.com/gabe565/mnemonic-ninja" target="_blank"
        v-bind="attrs" v-on="on">
        <v-icon>fab fa-github</v-icon>
      </v-btn>
    </template>
    <span>View on GitHub</span>
  </v-tooltip>
</template>

<script>
export default {
  name: 'GitHubButton',
};
</script>
