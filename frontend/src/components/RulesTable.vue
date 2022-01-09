<template>
  <v-simple-table class="v-data-table--outlined v-data-table--striped">
    <template v-slot:default>
      <thead>
      <tr>
        <th class="text-no-wrap" v-for="(row, key) of rules.headers" :key="key">
          <v-icon x-small>fas {{ row.icon }} fa-fw</v-icon>
          {{ row.text }}
        </th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="rule of rules.rules" :key="rule.number">
        <td>{{ rule.number }}</td>
        <td>
          <ul>
            <li v-for="sound of rule.sounds" :key="sound" v-html="sound"/>
          </ul>
        </td>
        <td>
          <ul>
            <li v-for="letter of rule.letters" :key="letter" v-html="letter"/>
          </ul>
        </td>
        <td>
          <ul>
            <li v-for="comment of rule.comments" :key="comment" v-html="comment"/>
          </ul>
        </td>
      </tr>
      </tbody>
    </template>
  </v-simple-table>
</template>

<script>
import rules from '@/data/rules';

export default {
  name: 'RulesTable',

  data: () => ({
    rules,
  }),
};
</script>

<style scoped lang="scss">
.v-data-table {
  ul {
    list-style-type: none;
    padding-left: 0;
  }
}
</style>
