import Papa from "papaparse";

export const parseCsv = async (data, config) => {
  return new Promise((complete, error) => {
    Papa.parse(data, { ...config, complete, error });
  });
};
