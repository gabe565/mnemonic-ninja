import { useRoute } from "vue-router";
import { useConversionApi } from "./conversion_api";
import router from "../plugins/router";
import { castArray, debounce } from "../util/helpers";
import { computed, ref, watch } from "vue";

export const castPair = (val) =>
  val.map((p) => {
    const split = p.split(",");
    return {
      word: split[0],
      number: split[1],
    };
  });

export const useQueryConverter = (type, props, usePairs = false) => {
  const query = ref("");
  const pairs = ref([]);

  const valid = ref(true);
  const route = useRoute();

  watch(
    () => props,
    (val) => {
      if (val.q && (!usePairs || val.pair)) {
        if (val.q || val.pair) {
          query.value = val.q || "";
        }
        if (usePairs && val.pair) {
          pairs.value = castPair(castArray(val.pair));
        } else {
          pairs.value = [];
        }
      } else {
        if (query.value) {
          updateUrl();
        }
      }
    },
    { immediate: true, deep: true },
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

  const updateUrl = async (push = false) => {
    const query = buildQueryParams();
    if (query.q !== props.q || query.pair !== props.pair) {
      if (push) {
        await router.push({ ...route, query });
      } else {
        await router.replace({ ...route, query });
      }
    }
  };

  watch(
    query,
    debounce(() => {
      return updateUrl();
    }, 200),
  );

  const { loading, lookupWordlist } = useConversionApi(type);
  const result = computed(() => {
    if (!query.value || loading.value) {
      return [];
    }

    if (usePairs) {
      const results = [];
      for (let i = 0; i <= query.value.length; i += 1) {
        const val = query.value.slice(0, query.value.length - i);
        results.push(lookupWordlist(val, i === query.value.length));
      }
      return results;
    } else {
      return lookupWordlist(query.value);
    }
  });

  return { query, pairs, result, valid, loading };
};
