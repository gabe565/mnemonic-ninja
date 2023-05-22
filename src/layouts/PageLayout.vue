<template>
  <v-container>
    <v-row class="pa-5" justify="center">
      <h1 class="text-h4">
        <slot name="title">
          {{ $route.name }}
        </slot>
      </h1>
    </v-row>

    <section>
      <h2 class="d-sr-only">Description</h2>
      <v-row v-if="$slots.description">
        <v-col>
          <v-card>
            <v-card-text>
              <slot name="description" />
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </section>

    <slot />
  </v-container>
</template>
<script setup></script>
