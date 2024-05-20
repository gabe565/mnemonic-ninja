export const ready = ref(false);
export const error = ref("");
export let wordlist = [];

const load = async () => {
  console.info("Loading wordlist...");
  let timeTaken = performance.now();
  try {
    const module = await import("./cmudict/cmudict.dict");
    wordlist = module.default.split("\n").map((e) => e.split(","));
    ready.value = true;
  } catch (err) {
    console.error(err);
    error.value = "Failed to load word list. Please refresh or try again later.";
    return;
  }
  timeTaken = performance.now() - timeTaken;
  console.info(
    `Loaded ${wordlist.length.toLocaleString()} words in ${timeTaken.toLocaleString()}ms`,
  );
};

load();
