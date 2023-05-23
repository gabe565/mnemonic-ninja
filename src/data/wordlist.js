export const ready = ref(false);
export let wordlist = [];

const load = async () => {
  console.info("Loading wordlist...");
  let timeTaken = performance.now();
  try {
    const module = await import("./cmudict/cmudict.dict");
    wordlist = module.default;
    ready.value = true;
  } catch (error) {
    console.error(error);
    return;
  }
  timeTaken = performance.now() - timeTaken;
  console.info(
    `Loaded ${wordlist.length.toLocaleString()} words in ${timeTaken.toLocaleString()}ms`
  );
};

load();
