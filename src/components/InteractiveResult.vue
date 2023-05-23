<template>
  <v-col class="px-0">
    <v-row>
      <v-col>
        <h3 class="text-h4">Chosen Phrase</h3>
      </v-col>
    </v-row>
    <v-row>
      <v-col class="v-card__text text--secondary">
        Here's the phrase you chose. Once you memorize this phrase, you can go to the
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
            <v-tooltip location="bottom" aria-label="Copy to Clipboard">
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

const props = defineProps({
  modelValue: {
    type: String,
    default: "",
  },
});

const showCopiedSnackbar = ref(false);

const copyPhrase = async () => {
  if (showCopiedSnackbar.value) {
    showCopiedSnackbar.value = false;
    await nextTick();
  }
  await navigator.clipboard.writeText(props.modelValue);
  showCopiedSnackbar.value = true;
};
</script>
