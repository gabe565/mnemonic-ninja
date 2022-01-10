<template>
  <v-col class="pa-0">
    <v-row>
      <v-col>
        <h2 class="text-h4">Chosen Phrase</h2>
      </v-col>
    </v-row>
    <v-row>
      <v-col class="v-card__text text--secondary">
        Here your phrase. Once you memorize this phrase, you can go to the
        <router-link :to="`/converters/word/${value}`">word to number converter</router-link> to change it back to your number.
        <br>
        You can also start over by typing in a new query, or change the chosen words by deleting
        a word up above.
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <v-text-field
          filled rounded readonly label="Phrase"
          v-model="value"
          @focus="$event.target.select()"
        >
          <template #prepend-inner>
            <v-tooltip bottom>
              <template #activator="{ on, attrs }">
                <v-btn
                  icon @click="copyPhrase"
                  v-bind="attrs" v-on="on"
                >
                  <v-icon>fas fa-copy</v-icon>
                </v-btn>
              </template>
              <span>Copy to Clipboard</span>
            </v-tooltip>
          </template>
        </v-text-field>
      </v-col>
    </v-row>
    <v-snackbar
      v-model="showCopiedSnackbar" timeout="5000"
      bottom class="pb-14 pb-md-0"
    >
      Copied to clipboard.
    </v-snackbar>
  </v-col>
</template>

<script>
export default {
  name: 'InteractiveResult',

  props: {
    value: String,
  },

  data: () => ({
    showCopiedSnackbar: false,
  }),

  methods: {
    async copyPhrase() {
      await navigator.clipboard.writeText(this.value);
      if (this.showCopiedSnackbar) {
        this.showCopiedSnackbar = false;
        await this.$nextTick();
      }
      this.showCopiedSnackbar = true;
    },
  },
};
</script>
