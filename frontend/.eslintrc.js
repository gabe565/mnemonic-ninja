module.exports = {
  root: true,
  env: {
    browser: true,
  },
  extends: [
    "airbnb",
    "plugin:vue/recommended",
    "plugin:vuetify/recommended",
    "prettier",
    "plugin:prettier/recommended",
  ],
  parserOptions: {
    parser: "@babel/eslint-parser",
  },
  rules: {
    "import/prefer-default-export": "off",
    "no-console": process.env.NODE_ENV === "production" ? "warn" : "off",
    "no-debugger": process.env.NODE_ENV === "production" ? "error" : "off",
    "no-param-reassign": ["error", { props: false }],
    "no-underscore-dangle": "off",
    "no-restricted-syntax": "off",
    "import/no-unresolved": "off",
  },
};
