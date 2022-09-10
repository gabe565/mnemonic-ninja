import debounce from "lodash/debounce";
import axios from "axios";
import { wait } from "../util/helpers";

export default (url) => ({
  data: () => ({
    query: "",
    disabled: false,
    loading: false,
    error: false,
    valid: false,
    response: {},
  }),

  watch: {
    // eslint-disable-next-line func-names
    query: debounce(function () {
      return this.apiUpdate();
    }, 200),
  },

  computed: {
    result() {
      if (this.response.result) {
        return this.response.result.map((value) => ({
          query: value.query,
          result: value.result?.join(", "),
        }));
      }
      return {};
    },
  },

  methods: {
    async manualUpdate() {
      this.disabled = true;
      await this.apiUpdate(true);
      await wait(1000);
      this.disabled = false;
    },
    apiUpdate(force = false) {
      return Promise.all([this.getResponse(force), this.updateUrl()]);
    },
    async getResponse(force = false) {
      if (!this.query || !this.valid) {
        this.response = {};
        this.error = null;
        this.loading = false;
        return;
      }
      if (!force && this.response?.query === this.query) {
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
