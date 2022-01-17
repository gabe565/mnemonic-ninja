export default (queryRegex, failMessage = 'Invalid input.') => ({
  async created() {
    this.rules = [
      (v) => !v || !queryRegex.test(v) || failMessage,
    ];
  },
});
