import { wordlist, loading } from "../data/wordlist";

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
  const lookupWordlist = (rawQuery, allowEmpty = false) => {
    if (!rawQuery && !allowEmpty) {
      return [];
    }

    rawQuery = rawQuery.toLowerCase();
    const queries = rawQuery.split(/[\s;,]/);

    const result = [];
    const srcType = types[type][0];
    const dstType = types[type][1];

    for (let query of queries) {
      query = query.trim();
      if (!query && !allowEmpty) {
        continue;
      }

      let matches = [];

      for (const word of wordlist) {
        if (word[srcType] === query) {
          matches.push(word[dstType]);
        }
      }

      // Filter unique
      matches = matches.filter((v, i, a) => a.indexOf(v) === i);

      let error = false;
      let guess = false;
      if (matches.length === 0) {
        if (type === "word") {
          let val = query.replaceAll(/[^A-Za-z]/g, "");
          error = val !== query;
          if (!error) {
            guess = true;
            for (const replacement of guesses) {
              val = val.replaceAll(replacement[0], replacement[1]);
            }
            val = val.replaceAll(/[^0-9]/g, "");
            matches.push(val);
          }
        } else if (type === "number") {
          if (query.match(/[^0-9]/)) {
            error = true;
          }
        }
      }

      result.push({
        query,
        result: matches,
        class: error ? "text-error" : guess ? "text-warning" : "",
        title: error ? "Invalid input" : guess ? "Unknown word. Output is approximated." : "",
      });
    }

    return result;
  };

  return { loading, lookupWordlist };
};
