<template>
  <v-form v-model="valid">
    <v-container>
      <v-row>
        <v-col>
          <h2 class="text-h4 text-center">Word to Number</h2>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <p>
            Enter a word or a list of words to get a converted list of numbers.
            <br />
            More than one number may show up for a single word. This means there is more than one
            pronunciation!
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
            variant="outlined"
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
          <v-table :height="height" class="v-table--variant-outlined v-table--variant-striped">
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
</template>

<script setup>
import ArrowRightIcon from "~icons/material-symbols/chevron-right";
import ArrowDownIcon from "~icons/material-symbols/keyboard-arrow-down";
import LoadingIcon from "~icons/mdi/shuriken";
import { useQueryConverter } from "../composables/query_converter";
import { useQueryRules } from "../composables/query_rules";
import { useDisplay } from "vuetify";

const props = defineProps({
  isActive: Boolean,
});
const emit = defineEmits(["query"]);

const { query, result, valid, loading } = useQueryConverter("word", props, emit);

const rules = useQueryRules(/[^A-Za-z-'\s,;.]/);

const { mdAndDown, smAndDown } = useDisplay();
const height = computed(() => (mdAndDown.value ? "212px" : "238px"));
const rows = computed(() => (mdAndDown.value ? 7 : 8));
const arrowIcon = computed(() => (smAndDown.value ? ArrowDownIcon : ArrowRightIcon));
</script>
