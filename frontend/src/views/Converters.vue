<template>
  <Page>
    <v-container>
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
              v-model="tab" center-active
              background-color="tertiary" class="rounded-b-0" show-arrows
            >
              <v-tab
                v-for="(tab, key) in tabs" :key="key"
                :to="{ name: 'Converters', params: { startTab: tab.slug }}"
              >{{ tab.name }}</v-tab>

              <v-tab-item
                v-for="(tab, key) in tabs" :key="key"
                :value="`/converters/${tab.slug}`"
              >
                <v-card-text>
                  <component :is="`${tab.slug}-converter`"/>
                </v-card-text>
              </v-tab-item>
            </v-tabs>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
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
    tab: 'interactive',
    tabs: [
      { slug: 'interactive', name: 'Interactive' },
      { slug: 'number', name: 'Number to Word' },
      { slug: 'word', name: 'Word to Number' },
    ],
  }),

  computed: {
    query() {
      return this.$route.query.q;
    },
  },

  created() {
    if (this.startTab) {
      this.tab = this.startTab;
    }
  },

  watch: {
    startTab(newVal) {
      if (newVal) {
        this.tab = newVal;
      }
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
]
