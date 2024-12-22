<template>
  <page-layout>
    <template #description>
      The
      <a href="https://en.wikipedia.org/wiki/Mnemonic_major_system" target="_blank"
        >Mnemonic Major System</a
      >
      is an effective technique for remembering numbers, though learning the associations between
      sounds and numbers can take time.
      <br />
      Unlike other online converters that simply match letters to numbers, Mnemonic Ninja performs
      conversions based on phonetic sounds, which is the correct method for using the system.
      <br />
      With a database of over 135,000 words, you can rely on Mnemonic Ninja for precise and helpful
      conversions.
    </template>
    <v-row tag="section">
      <v-col cols="12">
        <span id="conversion-rules" class="anchor" />
        <h2 class="text-h4">
          <v-icon :icon="RulesIcon" class="text-h4" />
          Conversion Rules
        </h2>
      </v-col>
      <v-col>
        <p>
          Based on the
          <a href="https://en.wikipedia.org/wiki/Mnemonic_major_system#The_system" target="_blank"
            >Wikipedia page</a
          >, here are the numbers and their corresponding sounds:
        </p>
        <rules-table />
      </v-col>
    </v-row>
    <v-row tag="section">
      <v-col cols="12">
        <span id="security" class="anchor" />
        <h2 class="text-h4">
          <v-icon :icon="SecurityIcon" class="text-h4" />
          Security
        </h2>
      </v-col>
      <v-col>
        <p>
          All conversions are performed directly in your browser. No input data is logged or sent to
          any remote servers.
        </p>
        <p>
          When you load the site for the first time (or after an update), a word list is downloaded
          to allow conversions to occur locally on your machine. You can verify this by
          <a href="https://developer.chrome.com/docs/devtools/network/#open" target="_blank"
            >inspecting this site's network activity</a
          >
          or by taking a look at the
          <a href="https://github.com/gabe565/mnemonic-ninja" target="_blank">codebase on GitHub</a
          >.
        </p>
        <p>
          Mnemonic Ninja even works without an internet connection! A service worker is installed
          which caches all necessary assets. If you're using a mobile phone, you can add the site to
          your home screen, open the app without an internet connection, and everything should
          behave normally!
        </p>
      </v-col>
    </v-row>
    <v-row tag="section">
      <v-col cols="12">
        <span id="development" class="anchor" />
        <h2 class="text-h4">
          <v-icon class="text-h4" :icon="DevelopmentIcon" />
          Development
        </h2>
      </v-col>
      <v-col>
        <p>
          To create this application, I used the wonderful
          <a href="http://www.speech.cs.cmu.edu/cgi-bin/cmudict" target="_blank">CMU Dictionary</a>
          (I actually used
          <a href="https://github.com/cmusphinx/cmudict" target="_blank"
            >this updated version on GitHub</a
          >).<br />
          Then I iterate over every Arpabet entry and convert it to its corresponding numeric form.
        </p>
        <p>
          The results are stored in the database instead of converting live so that queries take
          less processing and respond more quickly.
        </p>
        <p>
          When performing conversions, the application will attempt to perform the conversion from
          the pronunciation in the database first.<br />
          If the word cannot be found, then it will approximate the output by comparing letters
          instead of sounds.<br />
        </p>
        <div>
          To see this in action,
          <router-link :to="{ path: '/convert/word', query: { q: 'garage garages garagez' } }">
            look at the conversions of <v-code tag="code">garage</v-code>,
            <v-code tag="code">garages</v-code>, and
            <v-code tag="code">garagez</v-code> </router-link
          >:
          <ol>
            <li>
              <v-code tag="code">garage</v-code> becomes "746" because the last /ge/ is not a hard
              G, but a soft G which corresponds to the number 6.
            </li>
            <li>
              <v-code tag="code">garages</v-code> will successfully be converted to "7460" because
              the pronunciation is the same, but with an /s/ at the end of the word.
            </li>
            <li>
              Now to see the wordlist in action. <v-code tag="code">garagez</v-code> will convert to
              "7470" because the word is not in the wordlist, so the system has to approximate the
              output.
            </li>
          </ol>
        </div>
        <p>
          If you find any issues in the site, feel free to
          <a href="https://gabecook.com/connect" target="_blank"
            >email me through my contact form</a
          >
          or
          <a href="https://github.com/gabe565/mnemonic-ninja/issues/new" target="_blank"
            >open an issue on this project's GitHub</a
          >.
        </p>
      </v-col>
    </v-row>
  </page-layout>
</template>

<script setup>
import RulesIcon from "~icons/material-symbols/checklist-rounded";
import SecurityIcon from "~icons/material-symbols/shield-lock-rounded";
import DevelopmentIcon from "~icons/material-symbols/code-rounded";
import PageLayout from "../layouts/PageLayout.vue";
import RulesTable from "../components/RulesTable.vue";
</script>
