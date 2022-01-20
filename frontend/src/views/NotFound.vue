<template>
  <Page>
    <v-row>
      <v-col>
        Oh no! The requested file could not be located.
      </v-col>
    </v-row>
  </Page>
</template>

<script>
import Page from '@/layouts/Page.vue';

export default {
  name: 'NotFoundPage',
  components: { Page },
};
</script>
