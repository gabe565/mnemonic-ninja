<template>
  <v-app>
    <v-btn class="position-absolute d-sr-only-focusable" href="#content" style="z-index: 1007"
      >Skip to main content</v-btn
    >

    <v-app-bar app theme="dark" color="primary">
      <v-btn to="/" class="text-body-2 text-none px-2" size="x-large">
        <template #prepend>
          <v-icon :icon="LogoIcon" style="--v-icon-size-multiplier: 1.25" />
        </template>

        <v-app-bar-title>Mnemonic Ninja</v-app-bar-title>
      </v-btn>

      <v-spacer />

      <GitHubButton />

      <template v-if="mdAndUp" #extension>
        <v-tabs :model-value="currentRoute" align-tabs="center" color="white" style="width: 100%">
          <v-tab v-for="route in routes" :key="route.path" :to="route.to">
            <v-icon :icon="route.icon" class="pr-2" size="x-large" />
            {{ route.name }}
          </v-tab>
        </v-tabs>
      </template>
    </v-app-bar>

    <v-bottom-navigation v-if="smAndDown" position="fixed" bg-color="primary" theme="dark" grow>
      <v-btn
        v-for="(route, key) in routes"
        :key="route.path"
        :active="currentRoute === key"
        :to="route.to"
      >
        <v-icon :icon="route.icon" />
        <span>{{ route.name }}</span>
      </v-btn>
    </v-bottom-navigation>

    <UpdateSnackbar />

    <v-main>
      <span id="content" class="anchor" />
      <router-view class="pb-16" />
    </v-main>
  </v-app>
</template>

<script setup>
import { mdiInformation, mdiSwapHorizontalCircle } from "@mdi/js";
import LogoIcon from "./assets/logo.svg";
import { useRoute } from "vue-router";
import { useDisplay, useTheme } from "vuetify";

const routes = ref([
  { name: "Convert", icon: mdiSwapHorizontalCircle, to: "/convert" },
  { name: "About", icon: mdiInformation, to: "/about" },
]);

const currentRoute = computed(() => {
  const route = useRoute();
  return routes.value.findIndex((e) => route.path.startsWith(e.to));
});

const { mdAndUp, smAndDown } = useDisplay();

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
.v-application {
  @media (min-width: 1920px) {
    .v-container {
      max-width: 1400px;
    }
  }

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
  top: -150px;
  visibility: hidden;
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
