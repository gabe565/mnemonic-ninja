import globals from "globals";
import pluginJs from "@eslint/js";
import pluginVue from "eslint-plugin-vue";
import pluginVuetify from "eslint-plugin-vuetify";
import pluginPrettier from "eslint-plugin-prettier/recommended";

export default [
  {
    languageOptions: {
      globals: { ...globals.node, ...globals.browser },
    },
  },
  { ignores: ["dist", "dev-dist"] },
  pluginJs.configs.recommended,
  ...pluginVue.configs["flat/recommended"],
  ...pluginVuetify.configs["flat/recommended"],
  pluginPrettier,
  {
    rules: {
      "no-unused-vars": ["error", { varsIgnorePattern: "^_", argsIgnorePattern: "^_" }],
      "vue/no-template-shadow": "off",
      "vue/component-name-in-template-casing": ["error", "kebab-case"],
    },
  },
];
