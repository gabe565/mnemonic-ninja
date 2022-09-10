<template>
  <Page>
    <template #description> Oh no! The requested file could not be located. </template>
  </Page>
</template>

<script>
import Page from "../layouts/PageLayout.vue";

export default {
  name: "NotFoundPage",
  components: { Page },
};
</script>
