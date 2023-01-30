<template>
  <v-app>
    <v-app-bar app dark color="primary" hide-on-scroll>
      <v-img
        alt="Mnemonic Ninja Logo"
        class="shrink mr-2 pa-3"
        contain
        :src="require('./assets/logo.svg')"
        transition="scale-transition"
        width="40"
      />

      <v-toolbar-title> Mnemonic Ninja </v-toolbar-title>

      <v-spacer />

      <GitHubButton />

      <template v-if="$vuetify.breakpoint.mdAndUp" #extension>
        <v-tabs centered color="white">
          <v-tab v-for="route in routes" :key="route.path" :to="route.to" :href="route.href">
            <v-icon class="pr-2"> fas {{ route.icon }} fa-fw </v-icon>
            {{ route.name }}
          </v-tab>
        </v-tabs>
      </template>
    </v-app-bar>

    <UpdateSnackbar />

    <v-main>
      <router-view class="pb-16" />
    </v-main>

    <v-bottom-navigation v-if="$vuetify.breakpoint.smAndDown" fixed background-color="primary" dark>
      <v-btn v-for="route in routes" :key="route.path" :to="route.to">
        <span>{{ route.name }}</span>
        <v-icon>fas {{ route.icon }} fa-fw</v-icon>
      </v-btn>
    </v-bottom-navigation>
  </v-app>
</template>

<script>
import GitHubButton from "./components/GitHubButton.vue";
import UpdateSnackbar from "./components/UpdateSnackbar.vue";

export default {
  name: "App",

  components: {
    GitHubButton,
    UpdateSnackbar,
  },

  data: () => ({
    routes: [
      { name: "Convert", icon: "fa-exchange-alt", to: "/convert" },
      { name: "About", icon: "fa-info-circle", to: "/about" },
    ],
  }),

  beforeMount() {
    // check for browser support
    if (window.matchMedia && window.matchMedia("(prefers-color-scheme)").media !== "not all") {
      // set to preferred scheme
      const mediaQuery = window.matchMedia("(prefers-color-scheme: dark)");
      this.$vuetify.theme.dark = mediaQuery.matches;
      // react to changes
      mediaQuery.addEventListener("change", (e) => {
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
        background-color: rgba(0, 0, 0, 0.05);
      }
    }

    @at-root #{selector-unify(".theme--dark", &)} {
      tbody tr:nth-of-type(odd) {
        background-color: rgba(255, 255, 255, 0.04);
      }
    }
  }
}

.fa-spin-2x {
  animation: fa-spin 750ms infinite linear;
}

code {
  font-family: SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", "Courier New", monospace;
}
</style>