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
        primary: colors.blueGrey,
        tertiary: colors.blueGrey.lighten4,
        accent: colors.indigo.lighten3,
      },
      dark: {
        primary: colors.blueGrey,
        tertiary: colors.blueGrey.darken4,
        accent: "#2d3874",
        anchor: colors.blueGrey.lighten3,
      },
    },
  },
});
