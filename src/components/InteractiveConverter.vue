<template>
  <v-form v-model="valid" @submit.prevent>
    <v-container>
      <v-row>
        <v-col>
          <h2 class="text-h4 text-center">Interactive</h2>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <p class="v-card__text text--secondary mb-0 pa-0">
            Enter a number to get a word cloud of available words. Select one of these words to
            start your phrase. The corresponding numbers will be filtered out of your query,
            updating the word cloud to show words for the unselected part. Keep selecting words to
            build a phrase!
          </p>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <v-text-field
            v-model="query"
            rounded
            variant="outlined"
            clearable
            label="Query"
            placeholder="70395"
            :rules="rules"
            @keydown.backspace.passive.capture="keyEvent"
            @keydown.left.passive.capture="keyEvent"
            @select.passive.capture="selectEvent"
            @click:clear="clear"
          >
            <template #prepend-inner>
              <span v-for="(pair, key) in pairs" :key="key" class="text-grey">{{
                pair.number
              }}</span>
            </template>
          </v-text-field>
        </v-col>
      </v-row>
      <template v-if="query || pairs.length">
        <v-row>
          <v-divider />
        </v-row>
        <interactive-toolbar :pairs="pairs" @go-back="goBack" @go-back-to="goBackTo" />
        <v-row>
          <v-divider />
        </v-row>
        <interactive-words v-if="result && !allUsed" :result="result" @select="select" />
      </template>
      <v-row v-if="allUsed">
        <v-col>
          <interactive-result v-model="phrase" />
        </v-col>
      </v-row>
    </v-container>
  </v-form>
</template>

<script setup>
import { useQueryConverter } from "../composables/query_converter";
import { useQueryRules } from "../composables/query_rules";

const props = defineProps({
  isActive: Boolean,
});
const emit = defineEmits(["query"]);

const rules = useQueryRules(/[^0-9]/);

const { query, pairs, result, valid } = useQueryConverter("number", props, emit, true);

const phrase = computed(() => {
  return pairs.value.map((e) => e.word).join(" ");
});

const allUsed = computed(() => {
  return !query.value && pairs.value.length;
});

let moved = false;

const keyEvent = async (event) => {
  moved = false;
  const { target } = event;
  if (!pairs.value.length) {
    return;
  }
  if (target.selectionStart === 0) {
    const addedLength = pairs.value[pairs.value.length - 1].number.length;
    const prevSelectionStart = target.selectionStart;
    const prevSelectionEnd = target.selectionEnd;
    goBack();
    await nextTick();
    target.selectionStart = prevSelectionStart + addedLength;
    target.selectionEnd = prevSelectionEnd + addedLength;
  } else if (target.selectionStart === 1) {
    if (target.selectionStart !== target.selectionEnd) {
      moved = true;
    }
  }
};

const selectEvent = async (event) => {
  const { target } = event;
  if (
    !pairs.value.length ||
    target.selectionStart !== 0 ||
    target.selectionEnd !== query.value.length ||
    moved
  ) {
    moved = false;
    return;
  }
  goBackTo(0);
  await nextTick();
  target.selectionStart = 0;
  target.selectionEnd = query.value.length;
};

const clear = () => {
  query.value = "";
  pairs.value = [];
};

const select = (pair) => {
  pairs.value.push(pair);
  query.value = query.value.slice(pair.number.length);
};

const goBack = (n = 1) => {
  for (let i = 0; i < n; i += 1) {
    query.value = pairs.value.pop().number + query.value;
  }
};

const goBackTo = (key) => {
  if (pairs.value.length) {
    goBack(pairs.value.length - key);
  }
};
</script>

<style scoped lang="scss">
.v-application.v-theme--dark .text-grey {
  color: #aaa;
}
.v-application.v-theme--light .text-grey {
  color: #777;
}
</style>
