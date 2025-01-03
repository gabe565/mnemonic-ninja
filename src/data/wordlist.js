import { ref } from "vue";

export const loading = ref(true);
export const error = ref("");
export let wordlist = [];

const load = async () => {
  try {
    loading.value = true;
    console.info("Loading wordlist...");
    let timeTaken = performance.now();
    ({ default: wordlist } = await import("./cmudict/cmudict.dict"));
    timeTaken = performance.now() - timeTaken;
    console.info(
      `Loaded ${wordlist.length.toLocaleString()} words in ${timeTaken.toLocaleString()}ms`,
    );
  } catch (err) {
    console.error(err);
    error.value = "Failed to load word list. Please refresh or try again later.";
  } finally {
    loading.value = false;
  }
};

load();
