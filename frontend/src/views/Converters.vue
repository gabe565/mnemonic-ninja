<template>
  <Page>
    <v-container>
      <v-row>
        <v-col>
          <p>
            The mnemonic major system aids in numeric memorization by linking numbers with specific phonetic sounds, allowing you to convert a number to a word.
            This site will help do these conversions for you.
            <router-link to="/about">Read more on the about page</router-link>!
          </p>
          <p>
            Fill in a single value or a comma-separated list on the left and view the resulting conversions on the right.
          </p>
        </v-col>
      </v-row>

      <v-row>
        <v-col>
          <v-card elevation="3">
            <v-tabs
              v-model="tab" :grow="$vuetify.breakpoint.smAndDown"
              background-color="tertiary"
            >
              <v-tab>Number to Word</v-tab>
              <v-tab>Word to Number</v-tab>

              <v-tab-item>
                <converter
                  title="Number to Word"
                  from-label="Number"
                  from-placeholder="70395"
                  :from-value="num"
                  :from-regex="/[^0-9\s,;]/"
                  to-label="Word"
                  url="/api/number"
                  class="pb-5"
                >
                  <h2 class="text-h4 pb-5 text-center">Number to Word</h2>
                  Enter a number or a list of numbers to get a converted list of words.
                  <br>
                  Many words can show up for a single number. If this happens, the result box will be scrollable.
                </converter>
              </v-tab-item>

              <v-tab-item>
                <converter
                  title="Word to Number"
                  from-label="Word"
                  from-placeholder="example"
                  :from-value="word"
                  :from-regex="/[^A-Za-z-'\s,;.]/"
                  to-label="Number"
                  url="/api/word"
                  class="pb-5"
                >
                  <h2 class="text-h4 pb-5 text-center">Word to Number</h2>
                  Enter a word or a list of words to get a converted list of numbers.
                  <br>
                  More than one number may show up for a single word. This means there is more than one pronunciation!
                </converter>
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
import Converter from '@/components/Converter.vue';

export const Tabs = {
  NumberToWord: 0,
  WordToNumber: 1,
};

export default {
  name: 'Converters',

  props: {
    startTab: Number,
    word: String,
    num: String,
  },

  data: () => ({
    tab: 0,
  }),

  created() {
    if (this.startTab) {
      this.tab = this.startTab;
    }
  },

  components: {
    Page,
    Converter,
  },
};
</script>
]
