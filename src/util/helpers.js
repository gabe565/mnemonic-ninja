export const castArray = (...args) => (args[0] instanceof Array ? args[0] : args);

export const debounce = (fn, timeout) => {
  let id;
  return (...args) => {
    if (id) {
      clearTimeout(id);
    }
    id = setTimeout(() => {
      id = null;
      fn(...args);
    }, timeout);
  };
};
