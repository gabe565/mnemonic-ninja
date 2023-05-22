<template>
  <PageLayout>
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
    <section>
      <v-row>
        <v-col>
          <v-card>
            <v-card-title>
              <span id="conversion-rules" class="anchor" />
              <h2 class="text-h4">Conversion Rules</h2>
            </v-card-title>
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
    </section>
    <section>
      <v-row>
        <v-col>
          <v-card>
            <v-card-title>
              <span id="security" class="anchor" />
              <h2 class="text-h4">Security</h2>
            </v-card-title>
            <v-card-text>
              <p>
                All conversions are performed directly in your browser. No input data is logged or
                sent to any remote servers.
              </p>
              <p>
                When you load the site for the first time (or after an update), a word list is
                downloaded to allow conversions to occur locally on your machine. You can verify
                this by
                <a href="https://developer.chrome.com/docs/devtools/network/#open" target="_blank"
                  >inspecting this site's network activity</a
                >
                or by taking a look at the
                <a href="https://github.com/gabe565/mnemonic-ninja" target="_blank"
                  >codebase on GitHub</a
                >.
              </p>
              <p>
                Mnemonic Ninja even works without an internet connection! A service worker is
                installed which caches all necessary assets. If you're using a mobile phone, you can
                add the site to your home screen, open the app without an internet connection, and
                everything should behave normally!
              </p>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </section>
    <section>
      <v-row>
        <v-col>
          <v-card>
            <v-card-title>
              <span id="development" class="anchor" />
              <h4 class="text-h4">Development</h4>
            </v-card-title>
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
                The results are stored in the database instead of converting live so that queries
                take less processing and respond more quickly.
              </p>
              <p>
                When performing conversions, the application will attempt to perform the conversion
                from the pronounciation in the database first.<br />
                If the word cannot be found, then it will fallback to estimating the conversion by
                comparing letters instead of sounds.<br />
              </p>
              <div>
                To see this in action,
                <router-link
                  :to="{ path: '/convert/word', query: { q: 'garage garages garagez' } }"
                >
                  look at the conversions of <v-code tag="code">garage</v-code>,
                  <v-code tag="code">garages</v-code>, and
                  <v-code tag="code">garagez</v-code> </router-link
                >:
                <ol>
                  <li>
                    <v-code tag="code">garage</v-code> becomes "746" because the last /ge/ is not a
                    hard G, but a soft G which corresponds to the number 6.
                  </li>
                  <li>
                    <v-code tag="code">garages</v-code> will successfully be converted to "7460"
                    because the pronounciation is the same, but with an /s/ at the end of the word.
                  </li>
                  <li>
                    Now to see the wordlist in action. <v-code tag="code">garagez</v-code> will
                    convert to "7470" because the word is not in the wordlist, so the system has to
                    guess on the conversion.
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
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </section>
  </PageLayout>
</template>

<script setup></script>
