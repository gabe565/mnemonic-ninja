<template>
  <v-container>
    <v-row class="pa-5" justify="center">
      <h1 class="display-2">
        {{ title }}
      </h1>
    </v-row>
    <slot/>
  </v-container>
</template>

<script>
export default {
  name: 'Page',

  props: {
    titleOverride: String,
  },

  data: () => ({
    title: '',
  }),

  created() {
    this.title = this.titleOverride || this.$route.name;
  },
};
</script>
