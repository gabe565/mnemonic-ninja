<template>
  <v-table class="v-table--variant-striped rounded border">
    <template #default>
      <thead>
        <tr>
          <th v-for="(row, key) of rules.headers" :key="key" class="text-no-wrap">
            <v-icon :icon="row.icon" size="small" />
            {{ row.text }}
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="rule of rules.rules" :key="rule.number">
          <td>{{ rule.number }}</td>
          <td>
            <ul>
              <!-- eslint-disable-next-line vue/no-v-html -->
              <li v-for="sound of rule.sounds" :key="sound" v-html="sound" />
            </ul>
          </td>
          <td>
            <ul>
              <!-- eslint-disable-next-line vue/no-v-html -->
              <li v-for="letter of rule.letters" :key="letter" v-html="letter" />
            </ul>
          </td>
          <td>
            <ul>
              <!-- eslint-disable-next-line vue/no-v-html -->
              <li v-for="comment of rule.comments" :key="comment" v-html="comment" />
            </ul>
          </td>
        </tr>
      </tbody>
    </template>
  </v-table>
</template>

<script setup>
import rules from "../data/rules";
</script>

<style scoped lang="scss">
.v-table {
  ul {
    list-style-type: none;
    padding-left: 0;
  }
}
</style>
