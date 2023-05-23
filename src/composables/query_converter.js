import isEqual from "lodash/isEqual";
import { useRoute } from "vue-router";
import { useConversionApi, wordlistReady } from "./conversion_api";
import debounce from "lodash/debounce";
import router from "../plugins/router";

export const castPair = (val) =>
  val.map((p) => {
    const split = p.split(",");
    return {
      word: split[0],
      number: split[1],
    };
  });

export const useQueryConverter = (type, props, emit, usePairs = false) => {
  const query = ref("");
  const pairs = ref([]);

  const valid = ref(true);
  const route = useRoute();

  watch(
    () => route.query,
    (val) => {
      if (props.isActive && val) {
        if (val.q) {
          query.value = val.q;
        }
        if (usePairs && val.pair) {
          pairs.value = castPair(val.pair);
        } else {
          pairs.value = [];
        }
      }
    },
    { immediate: true, flush: "post" }
  );

  const buildQueryParams = () => {
    const result = {};
    if (query.value) {
      result.q = query.value;
    }
    if (usePairs && pairs.value?.length) {
      result.pair = pairs.value.map((pair) => [pair.word, pair.number].join(","));
    }
    return result;
  };

  const updateUrl = async () => {
    const query = buildQueryParams();
    if (!isEqual(query, route.query)) {
      await router.replace({ ...route, query });
    }
    emit("query", query);
  };

  watch(
    query,
    debounce(() => {
      if (props.isActive) {
        return updateUrl();
      }
    }, 200)
  );

  const { lookupWordlist } = useConversionApi(type);
  const result = computed(() => {
    if (!query.value || !wordlistReady.value) {
      return [];
    }

    if (usePairs) {
      const results = [];
      for (let i = 0; i < query.value.length; i += 1) {
        const val = query.value.slice(0, query.value.length - i);
        results.push(lookupWordlist(val));
      }
      return results;
    } else {
      return lookupWordlist(query.value);
    }
  });

  return { query, pairs, result, valid };
};
