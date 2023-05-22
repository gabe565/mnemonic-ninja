import isEqual from "lodash/isEqual";

export default {
  props: {
    isActive: Boolean,
  },

  watch: {
    "$route.query.q": {
      handler(val) {
        if (val && this.isActive) {
          this.query = val;
        }
      },
      immediate: true,
      flush: "post",
    },
  },

  emits: ["query"],

  methods: {
    buildQueryParams() {
      const query = {};
      if (this.query) {
        query.q = this.query;
      }
      return query;
    },
    async updateUrl() {
      const query = this.buildQueryParams();
      if (!isEqual(query, this.$route.query)) {
        await this.$router.replace({ ...this.$route, query });
      }
      this.$emit("query", query);
    },
  },
};
