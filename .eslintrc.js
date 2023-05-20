module.exports = {
  root: true,
  env: {
    browser: true,
  },
  extends: [
    "google",
    "plugin:vue/recommended",
    "plugin:vuetify/recommended",
    "prettier",
    "plugin:prettier/recommended",
  ],
  parserOptions: {
    parser: "@babel/eslint-parser",
  },
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
};
