export default (failMessage = 'Invalid input.') => ({
  props: {
    queryRegex: RegExp,
  },
  async created() {
    this.rules = [
      (v) => !v || !this.queryRegex.test(v) || failMessage,
    ];
  },
});
