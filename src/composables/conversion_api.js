/* eslint-disable no-invalid-this */

let wordlist = [];
export const wordlistReady = ref(false);

const parseWordlist = async () => {
  console.info("Loading wordlist...");
  let timeTaken = performance.now();
  try {
    const module = await import("../data/cmudict/cmudict.dict");
    wordlist = module.default;
    wordlistReady.value = true;
  } catch (error) {
    console.error(error);
    return;
  }
  timeTaken = performance.now() - timeTaken;
  console.info(
    `Loaded ${wordlist.length.toLocaleString()} words in ${timeTaken.toLocaleString()}ms`
  );
};
parseWordlist();

const types = {
  word: [0, 1],
  number: [1, 0],
  interactive: [1, 0],
};

const guesses = [
  ["th", "1"],
  ["ch", "6"],
  ["sh", "6"],
  ["ck", "7"],
  ["ph", "8"],
  ["s", "0"],
  ["z", "0"],
  ["t", "1"],
  ["d", "1"],
  ["n", "2"],
  ["m", "3"],
  ["r", "4"],
  ["l", "5"],
  ["j", "6"],
  ["c", "7"],
  ["g", "7"],
  ["k", "7"],
  ["q", "7"],
  ["f", "8"],
  ["v", "8"],
  ["b", "9"],
  ["p", "9"],
];

export const useConversionApi = (type) => {
  const loading = computed(() => !wordlistReady.value);

  const lookupWordlist = (rawQuery) => {
    if (!rawQuery) {
      return [];
    }

    rawQuery = rawQuery.toLowerCase();
    const queries = rawQuery.split(/[\s;,]/);

    const result = [];
    const srcType = types[type][0];
    const dstType = types[type][1];

    for (let query of queries) {
      query = query.trim();
      if (!query) {
        continue;
      }

      let matches = [];

      for (const word of wordlist) {
        if (word[srcType] === query && query !== "") {
          matches.push(word[dstType]);
        }
      }

      // Filter unique
      matches = matches.filter((v, i, a) => a.indexOf(v) === i);

      if (matches.length === 0 && type === "word") {
        let guess = query;
        for (const replacement of guesses) {
          guess = guess.replaceAll(replacement[0], replacement[1]);
        }
        guess = guess.replaceAll(/[^0-9]/g, "");
        matches.push(guess);
      }

      result.push({
        query,
        result: matches,
      });
    }

    return result;
  };

  return { loading, lookupWordlist };
};
