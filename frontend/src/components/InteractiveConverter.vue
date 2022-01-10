<template>
  <v-form v-model="valid">
    <v-container>
      <v-row>
        <v-col>
          <h2 class="text-h4 text-center">
            {{ title }}
          </h2>
        </v-col>
      </v-row>
      <v-row v-if="$slots.default">
        <v-col>
          <p class="v-card__text text--secondary mb-0 pa-0">
            <slot/>
          </p>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <v-text-field
            rounded filled label="Query" placeholder="70395" v-model="query"
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
import axios from 'axios';
import { debounce } from 'lodash';
import { wait } from '@/util/helpers';
import InteractiveToolbar from '@/components/InteractiveToolbar.vue';
import InteractiveWords from '@/components/InteractiveWords.vue';
import InteractiveResult from '@/components/InteractiveResult.vue';

export default {
  components: {
    InteractiveResult,
    InteractiveWords,
    InteractiveToolbar,
  },
  props: {
    title: String,
    queryPlaceholder: String,
    queryRegex: RegExp,
    url: String,
  },
  data() {
    return {
      query: '',
      response: {},
      loading: false,
      error: false,
      valid: false,
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
  watch: {
    // eslint-disable-next-line func-names
    query: debounce(async function () {
      await this.getResponse();
    }, 200),
  },
  async created() {
    this.rules = [
      (v) => !v || !this.queryRegex.test(v) || 'Invalid input.',
    ];

    if (this.fromValue) {
      this.query = this.fromValue;
      await this.getResponse();
    }
  },
  methods: {
    async manualUpdate() {
      await this.getResponse();
      await wait(1000);
    },
    debouncedResponse: debounce(() => this.getResponse(), 1000, { leading: true }),
    async getResponse() {
      if (!this.query || !this.valid) {
        this.response = {};
        this.error = null;
        this.loading = false;
        return;
      }
      this.loading = true;
      try {
        const { data } = await axios.get(`${this.url}/${encodeURIComponent(this.query)}`);
        this.response = data;
        this.error = false;
        this.valid = true;
      } catch (err) {
        this.error = err;
      } finally {
        this.loading = false;
      }
    },
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
        await this.manualUpdate();
      }
    },
    async goBack(n = 1) {
      let { query } = this;
      for (let i = 0; i < n; i += 1) {
        query = this.pairs.pop().number + query;
      }
      this.query = query;
      await this.manualUpdate();
    },
    goBackTo(key) {
      return this.goBack(this.pairs?.length - key);
    },
  },
};
</script>
