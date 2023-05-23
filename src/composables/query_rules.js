export const useQueryRules = (queryRegex, failMessage = "Invalid input.") => [
  (v) => !v || !queryRegex.test(v) || failMessage,
];
