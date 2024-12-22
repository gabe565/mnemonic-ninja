<template>
  <section class="mt-6 text-center text-body-1">
    <v-row>
      <v-col cols="12">
        <h3 class="text-h4">Matches</h3>
      </v-col>
    </v-row>
    <v-row v-for="(query, key) in filtered" :key="query.query" tag="section" justify="center">
      <v-col cols="12" class="pt-1 pb-0">
        <h4 class="text-h5">{{ query.query.length === 0 ? "Filler" : query.query }}</h4>
      </v-col>
      <v-col cols="12" class="pa-1">
        <ul>
          <li v-for="(word, key) in query.result" :key="key">
            <a
              :title="query.query"
              href="#"
              @click.prevent="emit('select', { word, number: query.query })"
              >{{ word }}</a
            >&nbsp;
          </li>
        </ul>
      </v-col>
      <v-col v-if="key !== filtered.length - 1" cols="3">
        <v-divider style="opacity: 7.5%" />
      </v-col>
    </v-row>
  </section>
</template>

<script setup>
import { computed } from "vue";

const props = defineProps({
  result: {
    type: Array,
    default: () => [],
  },
});
const emit = defineEmits(["select"]);

const filtered = computed(() => props.result.map((e) => e[0]).filter((e) => e.result.length));
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
