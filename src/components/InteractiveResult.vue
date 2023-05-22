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
        <router-link :to="{ path: '/convert/word', query: { q: modelValue } }">
          word to number converter
        </router-link>
        to change it back to your number.
        <br />
        You can also start over by typing in a new query, or change the chosen words by deleting a
        word up above.
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <v-text-field
          :model-value="modelValue"
          variant="outlined"
          rounded
          readonly
          label="Phrase"
          @focus="$event.target.select()"
        >
          <template #prepend-inner>
            <v-tooltip location="bottom">
              <template #activator="{ props }">
                <v-btn icon v-bind="props" @click="copyPhrase">
                  <v-icon :icon="mdiContentCopy" />
                </v-btn>
              </template>
              <span>Copy to Clipboard</span>
            </v-tooltip>
          </template>
        </v-text-field>
      </v-col>
    </v-row>
    <v-snackbar v-model="showCopiedSnackbar" timeout="5000" bottom class="pb-14 pb-md-0">
      Copied to clipboard.
    </v-snackbar>
  </v-col>
</template>

<script setup>
import { mdiContentCopy } from "@mdi/js";
</script>

<script>
export default {
  props: {
    modelValue: {
      type: String,
      default: "",
    },
  },

  data: () => ({
    showCopiedSnackbar: false,
  }),

  methods: {
    async copyPhrase() {
      await navigator.clipboard.writeText(this.modelValue);
      if (this.showCopiedSnackbar) {
        this.showCopiedSnackbar = false;
        await this.$nextTick();
      }
      this.showCopiedSnackbar = true;
    },
  },
};
</script>
