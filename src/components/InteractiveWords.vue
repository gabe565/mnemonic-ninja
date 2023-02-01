<template>
  <v-row justify="center" class="text-body-1 text-center">
    <v-col>
      <ul>
        <template v-for="(queries, queriesKey) in result">
          <template v-for="(query, queryKey) in queries">
            <li v-for="(word, key) in query.result" :key="`${queriesKey}-${queryKey}-${key}`">
              <a
                :title="query.query"
                @click.prevent="$emit('select', { word, number: query.query })"
                >{{ word }}</a
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
