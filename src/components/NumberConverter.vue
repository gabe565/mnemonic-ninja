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
          <p class="v-card__text text--secondary mb-0 pa-0">
            Enter a number or a list of numbers to get a converted list of words.
            <br />
            Many words can show up for a single number. If this happens, the result box will be
            scrollable.
          </p>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" md class="py-0 py-md-1">
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
          <v-btn
            size="large"
            type="submit"
            :disabled="!valid || disabled"
            :color="error ? 'error' : 'accent'"
            elevation="0"
            min-width="0"
            class="px-3"
            aria-hidden="true"
            @click.prevent
          >
            <v-icon v-if="loading" :icon="mdiShuriken" />
            <template v-else>
              <v-icon v-if="$vuetify.display.smAndDown" :icon="mdiArrowDownBold" />
              <v-icon v-else :icon="mdiArrowRightBold" />
            </template>
          </v-btn>
        </v-col>

        <v-col cols="12" md class="py-0 py-md-1">
          <h2 class="text-h4" aria-hidden="true">Word</h2>
          <v-table :height="height" class="v-table--variant-outlined">
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
import { mdiArrowDownBold, mdiArrowRightBold, mdiShuriken } from "@mdi/js";
</script>

<script>
import ConversionApi from "../mixins/ConversionApi";
import QueryValidate from "../mixins/QueryValidate";
import QueryUrl from "../mixins/UrlQuery";

export default {
  mixins: [ConversionApi("number"), QueryValidate(/[^0-9\s,;]/), QueryUrl],
  computed: {
    height() {
      return this.$vuetify.display.mdAndDown ? "212px" : "238px";
    },
    rows() {
      return this.$vuetify.display.mdAndDown ? 7 : 8;
    },
    result() {
      if (!this.query || !this.valid) {
        return [];
      }
      return this.lookupWordlist(this.query);
    },
  },
};
</script>
