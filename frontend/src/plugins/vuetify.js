import Vue from 'vue';
import Vuetify from 'vuetify/lib/framework';
import colors from 'vuetify/lib/util/colors';

Vue.use(Vuetify);

export default new Vuetify({
  icons: {
    iconfont: 'fa',
  },
  theme: {
    dark: true,
    themes: {
      dark: {
        primary: colors.blueGrey,
        secondary: '#60688b',
        accent: colors.deepPurple,
      },
      light: {
        primary: colors.blueGrey,
        secondary: '#60688b',
        accent: colors.deepPurple,
      },
    },
  },
});
