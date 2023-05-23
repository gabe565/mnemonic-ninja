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
          <p class="mb-0">
            Enter a word or a list of words to get a converted list of numbers.
            <br />
            More than one number may show up for a single word. This means there is more than one
            pronunciation!
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
          <v-btn icon disabled aria-hidden="true" :loading="loading">
            <template #loader>
              <v-icon :icon="mdiShuriken" class="icon-spin" />
            </template>
            <v-icon :icon="arrow" />
          </v-btn>
        </v-col>

        <v-col cols="12" md class="py-0 py-md-3">
          <h2 class="text-h4" aria-hidden="true">Number</h2>
          <v-table :height="height" class="v-table--variant-outlined v-table--variant-striped">
            <caption class="d-sr-only">
              Number
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
import { mdiArrowDownBold, mdiArrowRightBold, mdiShuriken } from "@mdi/js";
import { useQueryConverter } from "../composables/query_converter";
import { useQueryRules } from "../composables/query_rules";
import { useDisplay } from "vuetify";

const props = defineProps({
  isActive: Boolean,
});
const emit = defineEmits(["query"]);

const { query, result, valid, loading } = useQueryConverter("word", props, emit);

const rules = useQueryRules(/[^A-Za-z-'\s,;.]/);

const display = useDisplay();
const height = computed(() => (display.mdAndDown.value ? "212px" : "238px"));
const rows = computed(() => (display.mdAndDown.value ? 7 : 8));
const arrow = computed(() => (display.smAndDown.value ? mdiArrowDownBold : mdiArrowRightBold));
</script>
