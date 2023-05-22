<template>
  <v-form v-model="valid" @submit.prevent>
    <v-container>
      <v-row>
        <v-col>
          <h2 class="text-h4 text-center">Interactive</h2>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <p class="v-card__text text--secondary mb-0 pa-0">
            Enter a number to get a word cloud of available words. Select one of these words to
            start your phrase. The corresponding numbers will be filtered out of your query,
            updating the word cloud to show words for the unselected part. Keep selecting words to
            build a phrase!
          </p>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <v-text-field
            v-model="query"
            rounded
            variant="filled"
            clearable
            label="Query"
            placeholder="70395"
            :loading="loading"
            :rules="rules"
            @keydown.backspace="reset"
            @select="resetAll"
          >
            <template #prepend-inner>
              <div style="margin-top: 20px">
                <span v-for="(pair, key) in pairs" :key="key">{{ pair.number }}</span>
              </div>
            </template>
          </v-text-field>
        </v-col>
      </v-row>
      <template v-if="query || pairs.length">
        <v-row>
          <v-divider />
        </v-row>
        <interactive-toolbar :pairs="pairs" @go-back="goBack" @go-back-to="goBackTo" />
        <v-row>
          <v-divider />
        </v-row>
        <interactive-words v-if="result" :result="result" @select="select" />
      </template>
      <v-row v-if="allUsed">
        <v-col>
          <interactive-result v-model="phrase" />
        </v-col>
      </v-row>
    </v-container>
  </v-form>
</template>

<script>
import ConversionApi from "../mixins/ConversionApi";
import QueryValidate from "../mixins/QueryValidate";
import QueryUrl from "../mixins/UrlQuery";
import QueryUrlPair from "../mixins/UrlPair";

export default {
  mixins: [ConversionApi("interactive"), QueryValidate(/[^0-9\s,;]/), QueryUrl, QueryUrlPair],
  data() {
    return {
      pairs: [],
      showCopiedSnackbar: false,
    };
  },
  computed: {
    usedNumbers() {
      return this.pairs.map((e) => e.number).join("");
    },
    phrase() {
      return this.pairs.map((e) => e.word).join(" ");
    },
    allUsed() {
      return !this.query && this.pairs.length;
    },
    result() {
      if (!this.query || !this.valid) {
        return [];
      }

      const results = [];
      for (let i = 0; i < this.query.length; i += 1) {
        const query = this.query.slice(0, this.query.length - i);
        results.push(this.lookupWordlist(query));
      }

      return results;
    },
  },
  methods: {
    async reset() {
      if (this.pairs.length && this.query.length <= 1) {
        await this.goBack();
      }
    },
    resetAll() {
      return this.goBack(this.pairs.length);
    },
    async select(pair) {
      this.pairs.push(pair);
      this.query = this.query.slice(pair.number.length);
      await this.updateUrl();
    },
    async goBack(n = 1) {
      let { query } = this;
      for (let i = 0; i < n; i += 1) {
        query = this.pairs.pop().number + query;
      }
      this.query = query;
      await this.updateUrl();
    },
    async goBackTo(key) {
      if (this.pairs.length) {
        await this.goBack(this.pairs.length - key);
      }
    },
  },
};
</script>
