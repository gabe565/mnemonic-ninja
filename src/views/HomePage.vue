<template>
  <page-layout>
    <template #description>
      The
      <a href="https://en.wikipedia.org/wiki/Mnemonic_major_system" target="_blank"
        >Mnemonic Major System</a
      >
      is a strategy for memorizing numbers by converting them into memorable words and phrases,
      using specific phonetic sounds.
      <br />
      Mnemonic Ninja is a free and open-source service that transforms numbers into mental images
      that are much easier to remember than abstract digits. Whether it's a phone number, birthday,
      or address, Mnemonic Ninja simplifies the conversion process, making it effortless for you to
      recall important information. Mnemonic Ninja is useful for both newcomers learning the
      Mnemonic Major System, and experienced users who might need help converting a specific number.
      <br />
      Read more about the Major System and how it can help your memory
      <router-link to="/about">on the about page</router-link>!
    </template>

    <v-row tag="section">
      <v-col cols="12">
        <span id="converters" class="anchor" />
        <h2 class="text-h4">
          <v-icon :icon="SwapIcon" class="text-h4" />
          Converters
        </h2>
      </v-col>
      <v-col v-for="converter in converters" :key="converter.name" cols="12" lg="4">
        <v-card
          height="100%"
          :to="converter.path"
          :title="converter.title"
          :text="converter.text"
          variant="tonal"
        />
      </v-col>
    </v-row>
  </page-layout>
</template>

<script setup>
import PageLayout from "../layouts/PageLayout.vue";
import SwapIcon from "~icons/material-symbols/swap-horizontal-circle-rounded";

const converters = [
  {
    title: "Interactive",
    text: "Enter a number and select words from generated lists to form a sentence that can be converted back to the original number.",
    path: "/converters/interactive",
  },
  {
    title: "Number to Word",
    text: "Enter a list of numbers to see corresponding words in a table format.",
    path: "/converters/number",
  },
  {
    title: "Word to Number",
    text: "Enter a list of words to see corresponding numbers in a table format.",
    path: "/converters/word",
  },
];
</script>
