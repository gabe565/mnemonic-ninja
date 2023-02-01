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

    <v-row>
      <v-col>
        <v-card elevation="3" class="overflow-hidden">
          <v-tabs
            :value="currentTab"
            center-active
            color="convertTab"
            background-color="tertiary"
            class="rounded-b-0"
            show-arrows
          >
            <v-tab v-for="(tab, key) in tabs" :key="key" :to="`/convert/${key}`">
              {{ tab.name }}
            </v-tab>
          </v-tabs>

          <v-card-text>
            <v-tabs-items :value="currentTab" continuous @change="tabChange">
              <v-tab-item v-for="(tab, key) in tabs" :key="key" :value="key">
                <component
                  :is="tab.component"
                  :is-active="currentTab === key"
                  @query="tab.query = $event"
                />
              </v-tab-item>
            </v-tabs-items>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </PageLayout>
</template>

<script>
import InteractiveConverter from "../components/InteractiveConverter.vue";
import NumberConverter from "../components/NumberConverter.vue";
import WordConverter from "../components/WordConverter.vue";

export default {
  props: {
    startTab: {
      type: String,
      default: "interactive",
    },
  },

  data: () => ({
    currentTab: "interactive",
    tabs: {
      interactive: { name: "Interactive", query: {}, component: InteractiveConverter },
      number: { name: "Number to Word", query: {}, component: NumberConverter },
      word: { name: "Word to Number", query: {}, component: WordConverter },
    },
  }),

  watch: {
    startTab: {
      handler(val) {
        if (val) {
          this.currentTab = val;
        }
      },
      immediate: true,
    },
  },

  methods: {
    tabChange(key) {
      const to = this.buildLocation(this.tabs[key], key);
      this.$router.push(to);
    },
  },
};
</script>
