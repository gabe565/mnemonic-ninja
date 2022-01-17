import { debounce } from 'lodash';
import axios from 'axios';
import { wait } from '@/util/helpers';

export default (url) => ({
  props: {
    queryValue: String,
  },

  data: () => ({
    query: '',
    disabled: false,
    loading: false,
    error: false,
    valid: false,
    response: {},
  }),

  watch: {
    // eslint-disable-next-line func-names
    query: debounce(async function () {
      await this.getResponse();
      if (this.queryValue !== this.query) {
        let query;
        if (this.query) {
          query = { q: this.query };
        }
        await this.$router.replace({ ...this.$route, query });
      }
    }, 200),
    queryValue(newVal) {
      if (newVal) {
        this.query = newVal;
      }
    },
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
  },

  mounted() {
    if (this.queryValue) {
      this.query = this.queryValue;
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
        const { data } = await axios.get(`${url}/${encodeURIComponent(this.query)}`);
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
});
