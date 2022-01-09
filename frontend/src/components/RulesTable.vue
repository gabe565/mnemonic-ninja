<template>
  <v-simple-table class="v-data-table--outlined v-data-table--striped">
    <template v-slot:default>
      <thead>
      <tr>
        <th class="all" v-for="(row, key) of Rules.Headers" :key="key">
          <v-icon x-small>far {{ row.icon }} fa-fw</v-icon>
          &nbsp;{{ row.text }}
        </th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="rule of Rules.Rules" :key="rule.number">
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
import Rules from '@/data/rules';

export default {
  name: 'RulesTable',

  created() {
    this.Rules = Rules;
  },
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
