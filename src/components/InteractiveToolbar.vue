<template>
  <v-row align="center">
    <v-col cols="auto" align-self="center" class="pr-0">
      <v-tooltip v-model="showBackTooltip" location="top">
        <template #activator="{ props }">
          <v-btn icon v-bind="props" :disabled="!pairs.length" @click="goBack">
            <v-icon :icon="mdiArrowLeftBold" />
          </v-btn>
        </template>
        <span>Go Back</span>
      </v-tooltip>
    </v-col>
    <v-col>
      <v-tooltip v-for="(pair, key) in pairs" :key="key" location="top">
        <template #activator="{ props }">
          <v-chip v-bind="props" class="ma-1" closable close @click:close="emit('go-back-to', key)">
            {{ pair.word }}
          </v-chip>
        </template>
        <span>{{ pair.number }}</span>
      </v-tooltip>
    </v-col>
  </v-row>
</template>

<script setup>
import { mdiArrowLeftBold } from "@mdi/js";

const props = defineProps({
  pairs: {
    type: Array,
    default: () => [],
  },
});
const emit = defineEmits(["go-back", "go-back-to"]);

const showBackTooltip = ref(false);

const goBack = () => {
  if (props.pairs.length === 1) {
    showBackTooltip.value = false;
  }
  emit("go-back");
};
</script>
