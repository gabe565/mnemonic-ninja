import { defaultInclude, numberRegex, replacers } from "./config";

const truncateAt = (src, delimiter) => {
  const i = src.indexOf(delimiter);
  if (i === -1) {
    return src;
  }
  return src.slice(0, i);
};

const parseCmudictLine = (line) => {
  let [word, ...arpabets] = truncateAt(line, "#").trim().split(" ");
  word = truncateAt(word, "(");

  let number = "";
  for (const arpabet of arpabets) {
    const part = replacers[arpabet.trim()];
    if (part === undefined) {
      throw new Error(`Invalid arpabet in "${word}": "${arpabet}"`);
    }
    number += part;
  }

  if (numberRegex.test(number)) {
    throw new Error(`Letters detected in output: ${line}`);
  }

  return [word, number];
};

const compileFileToJS = (src) => {
  const data = src
    .trim()
    .split("\n")
    .map((e) => parseCmudictLine(e))
    .filter(Boolean);

  return `const data = ${JSON.stringify(data)};\nexport default data;`;
};

export const CompileCmudict = () => ({
  name: "mnemonic-ninja:compile-cmudict",

  transform(src, id) {
    for (const include of defaultInclude) {
      if (include.test(id)) {
        return {
          code: compileFileToJS(src),
        };
      }
    }
  },
});
