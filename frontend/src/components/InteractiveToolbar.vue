<template>
  <v-row>
    <v-col cols="auto" align-self="center" class="pr-0">
      <v-tooltip bottom>
        <template #activator="{ on, attrs }">
          <v-btn
            icon
            @click="$emit('go-back')"
            v-bind="attrs" v-on="on"
            :disabled="!pairs.length"
          >
            <v-icon>fas fa-arrow-alt-left fa-fw</v-icon>
          </v-btn>
        </template>
        <span>Go Back</span>
      </v-tooltip>
    </v-col>
    <v-col>
      <v-tooltip bottom v-for="(pair, key) in pairs" :key="key">
        <template #activator="{ on, attrs }">
          <v-chip
            @click:close="$emit('go-back-to', key)"
            v-bind="attrs" v-on="on"
            class="mr-2"
            close
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
  name: 'InteractiveToolbar',

  props: {
    pairs: Array,
  },
};
</script>
