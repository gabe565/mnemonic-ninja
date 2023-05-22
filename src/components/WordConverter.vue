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
          <p class="v-card__text text--secondary mb-0 pa-0">
            Enter a word or a list of words to get a converted list of numbers.
            <br />
            More than one number may show up for a single word. This means there is more than one
            pronunciation!
          </p>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" md class="d-flex flex-column py-0 py-md-1">
          <label for="word-query" class="text-h4">Word</label>
          <v-textarea
            id="word-query"
            v-model="query"
            class="flex-grow-1"
            clearable
            no-resize
            variant="outlined"
            name="Word"
            placeholder="example"
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
          <h2 class="text-h4" aria-hidden="true">Number</h2>
          <v-table :height="height" class="v-table--variant-outlined">
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
</script>

<script>
import ConversionApi from "../mixins/ConversionApi";
import QueryValidate from "../mixins/QueryValidate";
import QueryUrl from "../mixins/UrlQuery";

export default {
  mixins: [ConversionApi("word"), QueryValidate(/[^A-Za-z-'\s,;.]/), QueryUrl],
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
