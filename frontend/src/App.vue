<template>
  <v-app>
    <v-app-bar app dark color="primary" hide-on-scroll>
      <v-img
        alt="Mnemonic Ninja Logo"
        class="shrink mr-2 pa-3"
        contain
        :src="require('@/assets/logo.svg')"
        transition="scale-transition"
        width="40"
      />

      <v-toolbar-title>
        Mnemonic Ninja
      </v-toolbar-title>

      <v-spacer/>

      <GitHubButton/>

      <template #extension v-if="$vuetify.breakpoint.mdAndUp">
        <v-tabs centered color="white">
          <v-tab v-for="route in routes" :key="route.path" :to="route.path">
            <v-icon class="pr-2">fas {{ route.meta.icon }} fa-fw</v-icon>
            {{ route.name }}
          </v-tab>
        </v-tabs>
      </template>
    </v-app-bar>

    <UpdateSnackbar/>

    <v-main>
      <router-view class="pb-16"/>
    </v-main>

    <v-bottom-navigation
      fixed
      v-if="$vuetify.breakpoint.smAndDown"
      background-color="primary"
      dark
    >
      <v-btn
        v-for="route in routes"
        :key="route.path"
        :to="route.path"
        :value="route.name"
      >
        <span>{{ route.name }}</span>
        <v-icon>fas {{ route.meta.icon }} fa-fw</v-icon>
      </v-btn>
    </v-bottom-navigation>
  </v-app>
</template>

<script>
import GitHubButton from '@/components/GitHubButton.vue';
import UpdateSnackbar from '@/components/UpdateSnackbar.vue';

export default {
  name: 'App',

  components: {
    GitHubButton,
    UpdateSnackbar,
  },

  computed: {
    routes() {
      return this.$router.options.routes.filter((route) => route.meta?.showInNav);
    },
  },

  beforeMount() {
    // check for browser support
    if (window.matchMedia && window.matchMedia('(prefers-color-scheme)').media !== 'not all') {
      // set to preferred scheme
      const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)');
      this.$vuetify.theme.dark = mediaQuery.matches;
      // react to changes
      mediaQuery.addEventListener('change', (e) => {
        this.$vuetify.theme.dark = e.matches;
      });
    }
  },
};
</script>

<style lang="scss">
/* Outlined Data Tables */
.v-data-table {
  &--outlined {
    @at-root #{selector-unify(".theme--light", &)} {
      border: 1px solid rgba(0, 0, 0, 0.42);
    }

    @at-root #{selector-unify(".theme--dark", &)} {
      border: 1px solid rgba(255, 255, 255, 0.24);
    }
  }
}

/* Striped Data Tables */
.v-data-table {
  &--striped {
    @at-root #{selector-unify(".theme--light", &)} {
      tbody tr:nth-of-type(odd) {
        background-color: rgba(0, 0, 0, .05);
      }
    }

    @at-root #{selector-unify(".theme--dark", &)} {
      tbody tr:nth-of-type(odd) {
        background-color: rgba(255, 255, 255, .04);
      }
    }
  }
}

h1, h2, h3, h4, h5 {
  text-align: center;
}

.fa-spin-2x {
  animation: fa-spin 750ms infinite linear;
}
</style>
