<template>
  <converters-page>
    <v-form v-model="valid">
      <v-container>
        <v-row>
          <v-col>
            <p>
              Enter a list of words to see corresponding numbers in table format. If a word has
              multiple pronunciations, multiple numbers will be shown.
            </p>
            <p class="mb-0">
              <span class="text-warning">Yellow</span> rows are not in the dictionary, so the output
              is approximated.
              <br />
              <span class="text-error">Red</span> rows have invalid input.
            </p>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="12" md class="py-0 py-md-3">
            <label for="word-query" class="text-h4">Word</label>
            <v-textarea
              id="word-query"
              v-model="query"
              clearable
              no-resize
              variant="filled"
              name="Word"
              placeholder="example"
              :rules="rules"
              :rows="rows"
              :loading="loading"
            />
          </v-col>

          <v-col cols="12" md="auto" align-self="center" class="text-center py-0">
            <v-icon v-if="loading" :icon="LoadingIcon" class="icon-spin" />
            <v-icon v-else :icon="arrowIcon" />
          </v-col>

          <v-col cols="12" md class="py-0 py-md-3">
            <h2 class="text-h4" aria-hidden="true">Number</h2>
            <v-table :height="height" class="v-table--variant-striped rounded border">
              <caption class="d-sr-only">
                Number
              </caption>
              <tbody>
                <tr v-for="(item, key) in result" :key="key" :class="item.class">
                  <td
                    style="width: 1%"
                    :class="{ 'text-decoration-dotted': item.title }"
                    :title="item.title"
                  >
                    {{ item.query }}:
                  </td>
                  <td>{{ item.result.join(", ") }}</td>
                </tr>
              </tbody>
            </v-table>
          </v-col>
        </v-row>
      </v-container>
    </v-form>
  </converters-page>
</template>

<script setup>
import { computed } from "vue";
import { useDisplay } from "vuetify";
import { useQueryConverter } from "../composables/query_converter.js";
import { useQueryRules } from "../composables/query_rules.js";
import ConvertersPage from "../layouts/ConvertersPage.vue";
import ArrowRightIcon from "~icons/material-symbols/chevron-right";
import ArrowDownIcon from "~icons/material-symbols/keyboard-arrow-down";
import LoadingIcon from "~icons/mdi/shuriken";

const props = defineProps({
  q: {
    type: String,
    default: "",
  },
});

const { query, result, valid, loading } = useQueryConverter("word", props);

const rules = useQueryRules(/[^A-Za-z-'\s,;.]/);

const { mdAndDown, smAndDown } = useDisplay();
const height = computed(() => (mdAndDown.value ? "212px" : "238px"));
const rows = computed(() => (mdAndDown.value ? 7 : 8));
const arrowIcon = computed(() => (smAndDown.value ? ArrowDownIcon : ArrowRightIcon));
</script>
