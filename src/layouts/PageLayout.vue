<template>
  <div class="d-flex align-center" style="min-height: 100%">
    <v-container class="pt-6 pt-lg-12">
      <v-row v-if="$slots.description" tag="section">
        <h2 class="d-sr-only">Description</h2>
        <v-col>
          <slot name="description" />
        </v-col>
      </v-row>

      <slot />
    </v-container>
  </div>
</template>
