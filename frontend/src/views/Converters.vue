<template>
  <Page>
    <v-row>
      <v-col>
        <v-card>
          <v-card-text>
            The mnemonic major system aids in memorizing numbers by linking numbers with specific
            phonetic sounds, allowing you to convert a number to a word or a phrase.
            Phrases can create a mental image and be easier to remember than a number.
            This site will do these conversions for you to help you remember phone numbers,
            birthdays, addresses, etc.
            <router-link to="/about">
              Read more about how it works on the about page!
            </router-link>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>

    <v-row>
      <v-col>
        <v-card elevation="3" class="overflow-hidden">
          <v-tabs
            :value="currentTab" center-active
            background-color="tertiary" class="rounded-b-0" show-arrows
          >
            <v-tab
              v-for="(tab, key) in tabs" :key="key"
              :to="buildLocation(tab, key)"
            >{{ tab.name }}</v-tab>
          </v-tabs>

          <v-card-text>
            <v-tabs-items :value="currentTab" @change="tabChange" continuous>
              <v-tab-item
                v-for="(tab, key) in tabs" :key="key"
                :value="key"
              >
                  <component
                    :is="`${key}-converter`"
                    :is-active="currentTab === key"
                    @query="tab.query = $event"
                  />
              </v-tab-item>
            </v-tabs-items>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </Page>
</template>

<script>
import Page from '@/layouts/Page.vue';
import InteractiveConverter from '@/components/InteractiveConverter.vue';
import NumberConverter from '@/components/NumberConverter.vue';
import WordConverter from '@/components/WordConverter.vue';

export default {
  name: 'Converters',

  props: {
    startTab: String,
  },

  data: () => ({
    currentTab: 'interactive',
    tabs: {
      interactive: { name: 'Interactive', query: {} },
      number: { name: 'Number to Word', query: {} },
      word: { name: 'Word to Number', query: {} },
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
    buildLocation(tab, key) {
      return {
        name: this.$route.name,
        params: {
          startTab: key,
        },
        query: tab.query,
      };
    },
    tabChange(key) {
      const to = this.buildLocation(this.tabs[key], key);
      this.$router.push(to);
    },
  },

  components: {
    WordConverter,
    NumberConverter,
    Page,
    InteractiveConverter,
  },
};
</script>
