export const wait = (timeout) => new Promise((resolve) => setTimeout(resolve, timeout));

export const castArray = (...args) => (args[0] instanceof Array ? args[0] : args);
