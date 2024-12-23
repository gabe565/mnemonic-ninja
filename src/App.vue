<template>
  <v-app>
    <v-btn class="position-absolute d-sr-only-focusable" href="#content" style="z-index: 1007">
      Skip to main content
    </v-btn>

    <v-navigation-drawer v-model="drawer" color="primary" width="220" mobile-breakpoint="md">
      <template #prepend>
        <v-list>
          <v-list-item to="/" title="Mnemonic Ninja" :prepend-icon="LogoIcon" :active="false" />
        </v-list>
        <v-divider />
      </template>

      <v-list nav>
        <template v-for="(group, title) in routes" :key="title">
          <v-list-item
            v-for="route in group"
            :key="route.path"
            :to="route.path"
            link
            :title="route.name"
            :prepend-icon="route.meta.icon"
          />
          <v-divider v-if="title !== routeKeys[routeKeys.length - 1]" />
        </template>
      </v-list>

      <template #append>
        <v-divider />
        <v-list nav density="compact">
          <v-list-item
            href="https://gabecook.com"
            target="_blank"
            rel="noopener"
            title="Built by Gabe Cook"
            link
            :prepend-icon="PersonIcon"
          />
          <v-list-item
            href="https://github.com/gabe565/mnemonic-ninja"
            target="_blank"
            rel="noopener"
            title="Source on GitHub"
            :prepend-icon="GitHubIcon"
            link
          />
        </v-list>
      </template>
    </v-navigation-drawer>

    <v-bottom-navigation v-if="isMobile" position="fixed" bg-color="primary" theme="dark" grow>
      <template v-for="(group, title) in routes" :key="title">
        <template v-if="title === 'Converters'">
          <v-btn
            v-for="route in group"
            :key="route.path"
            :active="route.path === currentRoute.path"
            :to="route.path"
            :prepend-icon="route.meta.icon"
          >
            <span>{{ route.meta?.short || route.name }}</span>
          </v-btn>
        </template>
      </template>
    </v-bottom-navigation>

    <update-snackbar />

    <v-main>
      <v-app-bar color="primary" flat>
        <template #title>
          <template v-if="isMobile && $route.name === 'Home'"> Mnemonic Ninja</template>
          <template v-else>{{ $route.name }}</template>
        </template>
        <template #prepend>
          <v-btn v-if="isMobile" :icon="MenuIcon" @click.stop="drawer = !drawer" />
        </template>
      </v-app-bar>

      <span id="content" class="anchor" />
      <router-view v-slot="{ Component }">
        <keep-alive>
          <component :is="Component" />
        </keep-alive>
      </router-view>
    </v-main>
  </v-app>
</template>

<script setup>
import { computed, onBeforeMount, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import { useDisplay, useTheme } from "vuetify";
import UpdateSnackbar from "./components/UpdateSnackbar.vue";
import MenuIcon from "~icons/material-symbols/menu-rounded";
import PersonIcon from "~icons/material-symbols/person-rounded";
import LogoIcon from "~icons/mnemonic-ninja/logo";
import GitHubIcon from "~icons/simple-icons/github";

const { smAndDown: isMobile } = useDisplay();

const drawer = ref(!isMobile.value);

const currentRoute = useRoute();

const routes = computed(() => {
  return useRouter()
    .options.routes.filter((route) => route.meta?.showInNav)
    .reduce((acc, route) => {
      const group = route.meta?.group || "Other";
      if (!acc[group]) {
        acc[group] = [];
      }
      acc[group].push(route);
      return acc;
    }, {});
});

const routeKeys = computed(() => Object.keys(routes.value));

onBeforeMount(() => {
  // check for browser support
  if (window.matchMedia && window.matchMedia("(prefers-color-scheme)").media !== "not all") {
    // set to preferred scheme
    const mediaQuery = window.matchMedia("(prefers-color-scheme: dark)");
    const theme = useTheme();
    theme.name.value = mediaQuery.matches ? "dark" : "light";
    // react to changes
    mediaQuery.addEventListener("change", (e) => {
      theme.name.value = e.matches ? "dark" : "light";
    });
  }
});
</script>

<style lang="scss">
@use "sass:selector";

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

span.anchor {
  display: block;
  position: relative;
  top: -80px;
  visibility: hidden;
}

.v-table {
  /* Outlined Tables */
  &--variant-outlined {
    border-radius: 4px !important;

    @at-root #{selector.unify(".v-theme--light", &)} {
      border: thin solid rgba(0, 0, 0, 0.42);
    }

    @at-root #{selector.unify(".v-theme--dark", &)} {
      border: thin solid rgba(255, 255, 255, 0.24);
    }
  }

  /* Striped Tables */
  &--variant-striped {
    @at-root #{selector.unify(".v-theme--light", &)} {
      tbody tr:nth-of-type(odd) {
        background-color: rgba(0, 0, 0, 0.05);
      }
    }

    @at-root #{selector.unify(".v-theme--dark", &)} {
      tbody tr:nth-of-type(odd) {
        background-color: rgba(255, 255, 255, 0.04);
      }
    }
  }
}

.v-tab.v-btn .v-icon {
  --v-icon-size-multiplier: 1;
}

.icon-spin {
  animation-name: icon-spin;
  animation-duration: 1s;
  animation-iteration-count: infinite;
  animation-timing-function: linear;
}

@keyframes icon-spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}

code {
  font-family: SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", "Courier New", monospace;
}

.text-decoration-dotted {
  text-decoration: underline dotted;
}
</style>
