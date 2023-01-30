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
        <v-col cols="12" md class="py-0 py-md-1">
          <h2 class="text-h4">Word</h2>
          <v-textarea
            v-model="query"
            clearable
            no-resize
            outlined
            name="Word"
            placeholder="example"
            :rules="rules"
            :height="height"
          />
        </v-col>

        <v-col cols="12" md="auto" align-self="center" class="text-center py-0">
          <v-btn
            large
            type="submit"
            :disabled="!valid || disabled"
            :color="error ? 'error' : 'accent'"
            elevation="0"
            min-width="0"
            class="px-3"
            @click.prevent
          >
            <v-icon v-if="loading"> fas fa-fan fa-spin-2x fa-fw </v-icon>
            <v-icon v-else-if="error"> fas fa-exclamation-triangle fa-fw </v-icon>
            <template v-else>
              <v-icon v-if="$vuetify.breakpoint.smAndDown"> fas fa-arrow-alt-down fa-fw </v-icon>
              <v-icon v-else> fas fa-arrow-alt-right fa-fw </v-icon>
            </template>
          </v-btn>
        </v-col>

        <v-col cols="12" md class="py-0 py-md-1">
          <h2 class="text-h4">Number</h2>
          <v-simple-table :height="height" class="v-data-table--outlined">
            <tbody>
              <tr v-for="(item, key) in result" :key="key">
                <td style="width: 1%">{{ item.query }}:</td>
                <td>{{ item.result.join(", ") }}</td>
              </tr>
            </tbody>
          </v-simple-table>
        </v-col>
      </v-row>
    </v-container>
  </v-form>
</template>

<script>
import ConversionApi from "../mixins/ConversionApi";
import QueryValidate from "../mixins/QueryValidate";
import QueryUrl from "../mixins/UrlQuery";

export default {
  mixins: [ConversionApi("word"), QueryValidate(/[^A-Za-z-'\s,;.]/), QueryUrl],
  computed: {
    height() {
      return this.$vuetify.breakpoint.mdAndDown ? "200px" : "250px";
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

<style scoped lang="scss">
.v-data-table {
  @at-root #{selector-unify(".theme--light", &)} {
    background-color: #f6f6f6 !important;
  }
  @at-root #{selector-unify(".theme--dark", &)} {
    background-color: #1b1b1b !important;
  }
}
</style>