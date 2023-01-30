<template>
  <v-row justify="center" class="text-body-1 text-center">
    <v-col>
      <ul>
        <template v-for="(queries, queryKey) in result">
          <template v-for="query in queries">
            <li v-for="(word, key) in query.result" :key="`${queryKey}-${key}`">
              <a @click.prevent="$emit('select', { word, number: query.query })">{{ word }}</a
              >&nbsp;
            </li>
          </template>
        </template>
      </ul>
    </v-col>
  </v-row>
</template>

<script>
export default {
  name: "InteractiveWords",

  props: {
    result: {
      type: Array,
      default: () => [],
    },
  },
};
</script>

<style scoped lang="scss">
ul {
  list-style-type: none;
  padding-left: 0;
}
li {
  display: inline-block;
}
</style>
