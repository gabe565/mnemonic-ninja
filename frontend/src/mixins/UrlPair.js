import { castArray } from 'lodash';

export const castPair = (val) => val.map((p) => {
  const split = p.split(',');
  return {
    word: split[0],
    number: split[1],
  };
});

export default {
  data: () => ({
    pairs: [],
  }),

  watch: {
    '$route.query.pair': {
      async handler(val) {
        if (val) {
          this.$route.query.pair = val ? castArray(val) : undefined;
          this.pairs = castPair(this.$route.query.pair);
          if (this.query === '') {
            await this.updateUrl();
          }
        }
      },
      immediate: true,
    },
  },

  methods: {
    buildQueryParams() {
      const query = {};
      if (this.query) {
        query.q = this.query;
      }
      if (this.pairs?.length) {
        query.pair = this.pairs.map(
          (pair) => [pair.word, pair.number].join(','),
        );
      }
      return query;
    },
  },
};
