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
              v-model="tab" :grow="$vuetify.breakpoint.smAndDown"
              background-color="tertiary" class="rounded-b-0"
            >
              <v-tab :to="{ name: 'Converters', params: { startTab: 'interactive' }}">Interactive</v-tab>
              <v-tab :to="{ name: 'Converters', params: { startTab: 'num' }}">Batch Number to Word</v-tab>
              <v-tab :to="{ name: 'Converters', params: { startTab: 'word' }}">Batch Word to Number</v-tab>

              <v-tab-item value="/converters/interactive">
                <v-card-text>
                  <interactive-converter
                    title="Interactive"
                    query-placeholder="70395"
                    :query-regex="/[^0-9\s,;]/"
                    :query-value="startTab === 'interactive' ? query : undefined"
                    url="/api/interactive"
                  >
                    Enter a number to get a word cloud of available words.
                    Select one of these words to start your phrase.
                    The corresponding numbers will be filtered out of your query,
                    updating the word cloud to show words for the unselected part.
                    Keep selecting words to build a phrase!
                  </interactive-converter>
                </v-card-text>
              </v-tab-item>

              <v-tab-item value="/converters/num">
                <v-card-text>
                  <batch-converter
                    title="Number to Word"
                    query-label="Number"
                    query-placeholder="70395"
                    :query="startTab === 'num' ? query : undefined"
                    :query-regex="/[^0-9\s,;]/"
                    result-label="Word"
                    url="/api/number"
                  >
                    Enter a number or a list of numbers to get a converted list of words.
                    <br>
                    Many words can show up for a single number.
                    If this happens, the result box will be scrollable.
                  </batch-converter>
                </v-card-text>
              </v-tab-item>

              <v-tab-item value="/converters/word">
                <v-card-text>
                  <batch-converter
                    title="Word to Number"
                    query-label="Word"
                    query-placeholder="example"
                    :query-value="startTab === 'word' ? query : undefined"
                    :query-regex="/[^A-Za-z-'\s,;.]/"
                    result-label="Number"
                    url="/api/word"
                  >
                    Enter a word or a list of words to get a converted list of numbers.
                    <br>
                    More than one number may show up for a single word.
                    This means there is more than one pronunciation!
                  </batch-converter>
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
import BatchConverter from '@/components/BatchConverter.vue';
import InteractiveConverter from '@/components/InteractiveConverter.vue';

export default {
  name: 'Converters',

  props: {
    startTab: String,
    query: String,
  },

  data: () => ({
    tab: 'interactive',
  }),

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
    BatchConverter,
    Page,
    InteractiveConverter,
  },
};
</script>
]
