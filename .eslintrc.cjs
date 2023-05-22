/* eslint-env node */
require("@rushstack/eslint-patch/modern-module-resolution");

module.exports = {
  root: true,
  extends: [
    "google",
    "plugin:vue/vue3-recommended",
    "eslint:recommended",
    "@vue/eslint-config-prettier",
    "./.eslintrc-auto-import.json",
  ],
  rules: {
    "object-curly-spacing": ["error", "always"],
    "require-jsdoc": "off",
    indent: ["error", 2, { SwitchCase: 1 }],
    "no-unused-vars": [
      "error",
      { varsIgnorePattern: "^_", argsIgnorePattern: "^_" },
    ],
    "valid-jsdoc": "off",
    "new-cap": "off",
    "vue/no-template-shadow": "off",
  },
  parserOptions: {
    ecmaVersion: "latest",
  },
};
