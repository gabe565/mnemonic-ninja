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
          <p class="mb-0">
            Enter a number or a list of numbers to get a converted list of words.
            <br />
            Many words can show up for a single number. If this happens, the result box will be
            scrollable.
          </p>
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
          />
        </v-col>

        <v-col cols="12" md="auto" align-self="center" class="text-center py-0">
          <v-btn icon disabled aria-hidden="true">
            <v-icon :icon="arrow" />
          </v-btn>
        </v-col>

        <v-col cols="12" md class="py-0 py-md-3">
          <h2 class="text-h4" aria-hidden="true">Word</h2>
          <v-table :height="height" class="v-table--variant-outlined v-table--variant-striped">
            <caption class="d-sr-only">
              Word
            </caption>
            <tbody>
              <tr v-for="(item, key) in result" :key="key">
                <td style="width: 1%">{{ item.query }}:</td>
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
import { mdiArrowDownBold, mdiArrowRightBold } from "@mdi/js";
import { useQueryConverter } from "../composables/query_converter";
import { useQueryRules } from "../composables/query_rules";
import { useDisplay } from "vuetify";

const props = defineProps({
  isActive: Boolean,
});
const emit = defineEmits(["query"]);

const { query, result, valid } = useQueryConverter("number", props, emit);

const rules = useQueryRules(/[^0-9\s,;]/);

const display = useDisplay();
const height = computed(() => (display.mdAndDown.value ? "212px" : "238px"));
const rows = computed(() => (display.mdAndDown.value ? 7 : 8));
const arrow = computed(() => (display.smAndDown.value ? mdiArrowDownBold : mdiArrowRightBold));
</script>
