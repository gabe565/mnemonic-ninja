<template>
  <v-row>
    <v-col cols="auto" align-self="center" class="pr-0">
      <v-tooltip bottom>
        <template #activator="{ on, attrs }">
          <v-btn icon v-bind="attrs" :disabled="!pairs.length" @click="$emit('go-back')" v-on="on">
            <v-icon>fas fa-arrow-alt-left fa-fw</v-icon>
          </v-btn>
        </template>
        <span>Go Back</span>
      </v-tooltip>
    </v-col>
    <v-col>
      <v-tooltip v-for="(pair, key) in pairs" :key="key" bottom>
        <template #activator="{ on, attrs }">
          <v-chip
            v-bind="attrs"
            class="ma-1"
            close
            @click:close="$emit('go-back-to', key)"
            v-on="on"
          >
            {{ pair.word }}
          </v-chip>
        </template>
        <span>{{ pair.number }}</span>
      </v-tooltip>
    </v-col>
  </v-row>
</template>

<script>
export default {
  name: "InteractiveToolbar",

  props: {
    pairs: {
      type: Array,
      default: () => [],
    },
  },
};
</script>
