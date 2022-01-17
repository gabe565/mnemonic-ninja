<template>
  <v-form v-model="valid">
    <v-container>
      <v-row>
        <v-col>
          <h2 class="text-h4 text-center">
            Interactive
          </h2>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <p class="v-card__text text--secondary mb-0 pa-0">
            Enter a number to get a word cloud of available words.
            Select one of these words to start your phrase.
            The corresponding numbers will be filtered out of your query,
            updating the word cloud to show words for the unselected part.
            Keep selecting words to build a phrase!
          </p>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <v-text-field
            rounded filled clearable label="Query" placeholder="70395" v-model="query"
            @keydown.backspace="reset" @select="resetAll" :loading="loading"
            :rules="rules"
          >
            <template #prepend-inner>
              <div style="margin-top: 13px">
                <span
                  v-for="(pair, key) in pairs" :key="key"
                >{{ pair.number }}</span>
              </div>
            </template>
          </v-text-field>
        </v-col>
      </v-row>
      <template v-if="query || pairs.length">
        <v-row>
          <v-divider/>
        </v-row>
        <interactive-toolbar :pairs="pairs" @go-back="goBack" @go-back-to="goBackTo"/>
        <v-row>
          <v-divider/>
        </v-row>
        <interactive-words v-if="response.result" :response="response" @select="select"/>
      </template>
      <v-row v-if="allUsed">
        <v-col>
          <interactive-result v-model="phrase"/>
        </v-col>
      </v-row>
    </v-container>
  </v-form>
</template>

<script>
import InteractiveToolbar from '@/components/InteractiveToolbar.vue';
import InteractiveWords from '@/components/InteractiveWords.vue';
import InteractiveResult from '@/components/InteractiveResult.vue';
import ConversionApi from '@/mixins/ConversionApi';
import QueryValidate from '@/mixins/QueryValidate';
import QueryUrl from '@/mixins/UrlQuery';
import QueryUrlPair from '@/mixins/UrlPair';

export default {
  components: {
    InteractiveResult,
    InteractiveWords,
    InteractiveToolbar,
  },
  mixins: [
    ConversionApi('/api/interactive'),
    QueryValidate(/[^0-9\s,;]/),
    QueryUrl,
    QueryUrlPair,
  ],
  data() {
    return {
      pairs: [],
      showCopiedSnackbar: false,
    };
  },
  computed: {
    usedNumbers() {
      return this.pairs.map((e) => e.number).join('');
    },
    phrase() {
      return this.pairs.map((e) => e.word).join(' ');
    },
    allUsed() {
      return !this.query && this.pairs.length;
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
      if (!this.loading) {
        this.loading = true;
        this.pairs.push(pair);
        this.query = this.query.slice(pair.number.length);
        await this.apiUpdate();
      }
    },
    async goBack(n = 1) {
      let { query } = this;
      for (let i = 0; i < n; i += 1) {
        query = this.pairs.pop().number + query;
      }
      this.query = query;
      await this.apiUpdate();
    },
    goBackTo(key) {
      return this.goBack(this.pairs?.length - key);
    },
  },
};
</script>
