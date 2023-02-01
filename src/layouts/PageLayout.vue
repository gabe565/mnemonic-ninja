<template>
  <v-container>
    <v-row class="pa-5" justify="center">
      <h1 class="text-h4">
        <slot name="title">
          {{ $route.name }}
        </slot>
      </h1>
    </v-row>

    <v-row v-if="$slots.description">
      <v-col>
        <v-card>
          <v-card-text>
            <slot name="description" />
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>

    <slot />
  </v-container>
</template>
