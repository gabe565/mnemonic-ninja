<template>
  <v-form v-model="valid">
    <v-container class="pa-0">
      <v-row v-if="$slots.default">
        <v-col>
          <slot/>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" md class="py-0 py-md-1">
          <h2 class="text-h4">{{ fromLabel }}</h2>
          <v-textarea
            clearable no-resize outlined
            :name="fromLabel" :placeholder="fromPlaceholder"
            :rules="rules"
            :height="height"
            v-model="query"
          />
        </v-col>

        <v-col
          cols="12" md="auto"
          align-self="center" class="text-center pt-0 py-md-3"
        >
          <v-btn
            large
            type="submit"
            @click.prevent="manualUpdate"
            :disabled="!valid || disabled"
            :color="error ? 'error' : 'accent'" min-width="0" class="px-3"
          >
            <v-icon v-if="loading">fas fa-fan fa-spin-2x fa-fw</v-icon>
            <v-icon v-else-if="error">fas fa-exclamation-triangle fa-fw</v-icon>
            <template v-else>
              <v-icon v-if="$vuetify.breakpoint.smAndDown">fas fa-arrow-alt-down fa-fw</v-icon>
              <v-icon v-else>fas fa-arrow-alt-right fa-fw</v-icon>
            </template>
          </v-btn>
        </v-col>

        <v-col cols="12" md class="py-0 py-md-1">
          <h2 class="text-h4">{{ toLabel }}</h2>
          <v-simple-table :height="height" class="v-data-table--outlined">
            <tbody>
            <tr v-for="(item, key) in result" :key="key">
              <td style="width: 1%">{{ item.query }}:</td>
              <td>{{ item.result }}</td>
            </tr>
            </tbody>
          </v-simple-table>
        </v-col>
      </v-row>
    </v-container>
  </v-form>
</template>

<script>
import axios from 'axios';
import { debounce } from 'lodash';
import { wait } from '@/util/helpers';

export default {
  props: {
    fromLabel: String,
    fromPlaceholder: String,
    fromValue: String,
    fromRegex: RegExp,
    toLabel: String,
    url: String,
    description: String,
  },
  data() {
    return {
      query: '',
      response: {},
      disabled: false,
      loading: false,
      error: false,
      valid: false,
    };
  },
  computed: {
    result() {
      if (this.response.result) {
        return this.response.result.map((value) => ({
          query: value.query,
          result: value.result?.join(', '),
        }));
      }
      return {};
    },
    height() {
      return this.$vuetify.breakpoint.mdAndDown ? '200px' : '250px';
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
      (v) => !v || !this.fromRegex.test(v) || 'Invalid input.',
    ];

    if (this.fromValue) {
      this.query = this.fromValue;
      await this.getResponse();
    }
  },
  methods: {
    async manualUpdate() {
      this.disabled = true;
      await this.getResponse();
      await wait(1000);
      this.disabled = false;
    },
    async getResponse() {
      if (!this.query || !this.valid) {
        this.response = {};
        this.error = null;
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
        await wait(100);
        this.loading = false;
      }
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
