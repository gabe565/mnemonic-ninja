import { createVuetify } from "vuetify";
import { aliases, mdi } from "vuetify/iconsets/mdi-svg";
import colors from "vuetify/util/colors";

import "vuetify/styles";

export default createVuetify({
  icons: {
    defaultSet: "mdi",
    aliases,
    sets: {
      mdi,
    },
  },
  theme: {
    defaultTheme: "dark",
    themes: {
      light: {
        colors: {
          primary: colors.blueGrey.darken2,
          tertiary: colors.blueGrey.lighten5,
          accent: colors.indigo.lighten3,
          anchor: colors.blue.accent4,
          convertTab: "#1554ff",
        },
      },
      dark: {
        dark: true,
        colors: {
          primary: colors.blueGrey.darken2,
          tertiary: colors.blueGrey.darken4,
          accent: "#2d3874",
          anchor: colors.blue.accent1,
          convertTab: colors.blue.accent1,
        },
      },
    },
  },
});
