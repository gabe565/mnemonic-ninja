<template>
  <v-app>
    <v-app-bar app theme="dark" color="primary">
      <v-btn to="/" class="text-body-2 text-none px-2" size="x-large">
        <template #prepend>
          <img alt="Mnemonic Ninja Logo" :src="logo" style="height: 40px" />
        </template>

        <v-app-bar-title>Mnemonic Ninja</v-app-bar-title>
      </v-btn>

      <v-spacer />

      <GitHubButton />

      <template v-if="$vuetify.display.mdAndUp" #extension>
        <v-tabs :model-value="currentRoute" align-tabs="center" color="white" style="width: 100%">
          <v-tab v-for="route in routes" :key="route.path" :to="route.to">
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

    <v-bottom-navigation
      v-if="$vuetify.display.smAndDown"
      position="fixed"
      bg-color="primary"
      theme="dark"
      grow
    >
      <v-btn
        v-for="(route, key) in routes"
        :key="route.path"
        :active="currentRoute === key"
        :to="route.to"
      >
        <v-icon>fas {{ route.icon }} fa-fw</v-icon>
        <span>{{ route.name }}</span>
      </v-btn>
    </v-bottom-navigation>
  </v-app>
</template>

<script setup>
import logo from "./assets/logo.svg";
</script>

<script>
export default {
  name: "App",

  data: () => ({
    routes: [
      { name: "Convert", icon: "fa-exchange-alt", to: "/convert" },
      { name: "About", icon: "fa-info-circle", to: "/about" },
    ],
  }),

  computed: {
    currentRoute() {
      return this.routes.findIndex((e) => this.$route.path.startsWith(e.to));
    },
  },

  beforeMount() {
    // check for browser support
    if (window.matchMedia && window.matchMedia("(prefers-color-scheme)").media !== "not all") {
      // set to preferred scheme
      const mediaQuery = window.matchMedia("(prefers-color-scheme: dark)");
      this.$vuetify.theme.name = mediaQuery.matches ? "dark" : "light";
      // react to changes
      mediaQuery.addEventListener("change", (e) => {
        this.$vuetify.theme.name = e.matches ? "dark" : "light";
      });
    }
  },
};
</script>

<style lang="scss">
.v-application {
  p {
    margin-bottom: 1rem;
  }

  ul,
  ol {
    padding-left: 20px;
  }
}

a {
  color: rgb(var(--v-theme-anchor));

  &:not(:hover) {
    text-decoration: none;
  }
}

.v-table {
  /* Outlined Tables */
  &--variant-outlined {
    border-radius: 4px !important;

    @at-root #{selector-unify(".v-theme--light", &)} {
      border: thin solid rgba(0, 0, 0, 0.42);
    }

    @at-root #{selector-unify(".v-theme--dark", &)} {
      border: thin solid rgba(255, 255, 255, 0.24);
    }
  }

  /* Striped Tables */
  &--variant-striped {
    @at-root #{selector-unify(".v-theme--light", &)} {
      tbody tr:nth-of-type(odd) {
        background-color: rgba(0, 0, 0, 0.05);
      }
    }

    @at-root #{selector-unify(".v-theme--dark", &)} {
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
