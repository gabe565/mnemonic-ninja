import Vue from "vue";
import Vuetify from "vuetify/lib/framework";
import colors from "vuetify/lib/util/colors";

Vue.use(Vuetify);

export default new Vuetify({
  icons: {
    iconfont: "fa",
  },
  theme: {
    dark: true,
    themes: {
      light: {
        primary: colors.blueGrey.darken2,
        tertiary: colors.blueGrey.lighten5,
        accent: colors.indigo.lighten3,
        anchor: colors.blue.accent4,
        convertTab: "#1554ff",
      },
      dark: {
        primary: colors.blueGrey.darken2,
        tertiary: colors.blueGrey.darken4,
        accent: "#2d3874",
        anchor: colors.blue.accent1,
        convertTab: colors.blue.accent1,
      },
    },
  },
});
