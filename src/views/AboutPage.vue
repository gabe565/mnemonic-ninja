<template>
  <Page>
    <template #description>
      The
      <a href="https://en.wikipedia.org/wiki/Mnemonic_major_system" target="_blank"
        >Mnemonic Major System</a
      >
      is a great way to easily remember numbers, but learning to remember which sounds and numbers
      correspond to each other takes time.<br />
      There are other online converters, but they all performed conversions by directly changing
      specific letters to a number. This is not how the system is meant to work, though. The system
      is meant to work with phonetic sounds. That is how this site does conversions. If a word or
      number is entered that is in its database (Which contains over 135,000 words!), then it will
      perform the conversion according to how the word is pronounced.
    </template>
    <v-row class="pb-4">
      <v-col>
        <v-card>
          <v-card-title class="text-h4"> Conversion Rules </v-card-title>
          <v-card-text>
            <p>
              Based on the
              <a
                href="https://en.wikipedia.org/wiki/Mnemonic_major_system#The_system"
                target="_blank"
                >Wikipedia page</a
              >, here are the numbers and their corresponding sounds:
            </p>
            <RulesTable />
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <v-card>
          <v-card-title class="text-h4"> Development </v-card-title>
          <v-card-text>
            <p>
              To create this application, I used the wonderful
              <a href="http://www.speech.cs.cmu.edu/cgi-bin/cmudict" target="_blank"
                >CMU Dictionary</a
              >
              (I actually used
              <a href="https://github.com/cmusphinx/cmudict" target="_blank"
                >this updated version on GitHub</a
              >).<br />
              Then I iterate over every Arpabet entry and convert it to its corresponding numeric
              form.
            </p>
            <p>
              The results are stored in the database instead of converting live so that queries take
              less processing and respond more quickly.
            </p>
            <p>
              When performing conversions, the application will attempt to perform the conversion
              from the pronounciation in the database first.<br />
              If the word cannot be found, then it will fallback to estimating the conversion by
              comparing letters instead of sounds.<br />
            </p>
            <div>
              To see this in action,
              <router-link :to="{ path: '/convert/word', query: { q: 'garage garages garagez' } }">
                look at the conversions of <code>garage</code>, <code>garages</code>, and
                <code>garagez</code> </router-link
              >:
              <ol>
                <li>
                  <code>garage</code> becomes "746" because the last /ge/ is not a hard G, but a
                  soft G which corresponds to the number 6.
                </li>
                <li>
                  <code>garages</code> will successfully be converted to "7460" because the
                  pronounciation is the same, but with an /s/ at the end of the word.
                </li>
                <li>
                  Now to see the wordlist in action. <code>garagez</code> will convert to "7470"
                  because the word is not in the wordlist, so the system has to guess on the
                  conversion.
                </li>
              </ol>
            </div>
            <p>
              If you find any issues in the site, feel free to
              <a href="https://gabecook.com/connect" target="_blank"
                >email me through my contact form</a
              >
              or
              <a href="https://github.com/gabe565/mnemonic-major-converter" target="_blank"
                >open an issue on this project's GitHub</a
              >.
            </p>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </Page>
</template>

<script>
import Page from "../layouts/PageLayout.vue";
import RulesTable from "../components/RulesTable.vue";

export default {
  name: "AboutPage",

  components: {
    RulesTable,
    Page,
  },
};
</script>
