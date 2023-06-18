<template>
  <PageLayout>
    <template #title> Convert </template>
    <template #description>
      The mnemonic major system aids in memorizing numbers by linking numbers with specific phonetic
      sounds, allowing you to convert a number to a word or a phrase. Phrases can create a mental
      image and be easier to remember than a number. This site will do these conversions for you to
      help you remember phone numbers, birthdays, addresses, etc.
      <router-link to="/about"> Read more about how it works on the about page! </router-link>
    </template>

    <v-row v-if="error">
      <v-col>
        <v-alert type="error">{{ error }}</v-alert>
      </v-col>
    </v-row>

    <v-row>
      <v-col>
        <v-card elevation="3" class="overflow-hidden">
          <v-tabs
            :model-value="currentTab"
            center-active
            color="convertTab"
            bg-color="tertiary"
            class="rounded-b-0"
            show-arrows
          >
            <v-tab v-for="(tab, key) in tabs" :key="key" :to="`/convert/${tab.slug}`">
              {{ tab.name }}
            </v-tab>
          </v-tabs>

          <v-card-text>
            <v-window :model-value="currentTab" :touch="false">
              <v-window-item v-for="(tab, key) in tabs" :key="key">
                <component
                  :is="tab.component"
                  :is-active="currentTab === key"
                  @query="tab.query = $event"
                />
              </v-window-item>
            </v-window>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </PageLayout>
</template>

<script setup>
import InteractiveConverter from "../components/InteractiveConverter.vue";
import NumberConverter from "../components/NumberConverter.vue";
import WordConverter from "../components/WordConverter.vue";
import { error } from "../data/wordlist";

const props = defineProps({
  startTab: {
    type: String,
    default: "interactive",
  },
});

const currentTab = ref(0);

const tabs = ref([
  {
    slug: "interactive",
    name: "Interactive",
    query: {},
    component: markRaw(InteractiveConverter),
  },
  {
    slug: "number",
    name: "Number to Word",
    query: {},
    component: markRaw(NumberConverter),
  },
  {
    slug: "word",
    name: "Word to Number",
    query: {},
    component: markRaw(WordConverter),
  },
]);

watch(
  () => props.startTab,
  (val) => {
    if (val) {
      currentTab.value = Object.values(tabs.value).findIndex((e) => e.slug === val);
    }
  },
  { immediate: true },
);
</script>
