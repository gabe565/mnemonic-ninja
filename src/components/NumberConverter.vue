<template>
  <v-form v-model="valid">
    <v-container>
      <v-row>
        <v-col>
          <h2 class="text-h4 text-center">Number to Word</h2>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <p>
            Enter a number or a list of numbers to get a converted list of words.
            <br />
            Many words can show up for a single number. If this happens, the result box will be
            scrollable.
          </p>
          <p class="mb-0"><span class="text-error">Red</span> rows have invalid input.</p>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" md class="py-0 py-md-3">
          <label for="number-query" class="text-h4">Number</label>
          <v-textarea
            id="number-query"
            v-model="query"
            clearable
            no-resize
            variant="outlined"
            name="Number"
            placeholder="70395"
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
          <h2 class="text-h4" aria-hidden="true">Word</h2>
          <v-table :height="height" class="v-table--variant-outlined v-table--variant-striped">
            <caption class="d-sr-only">
              Word
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

const { query, result, valid, loading } = useQueryConverter("number", props, emit);

const rules = useQueryRules(/[^0-9\s,;]/);

const { mdAndDown, smAndDown } = useDisplay();
const height = computed(() => (mdAndDown.value ? "212px" : "238px"));
const rows = computed(() => (mdAndDown.value ? 7 : 8));
const arrowIcon = computed(() => (smAndDown.value ? ArrowDownIcon : ArrowRightIcon));
</script>
