/**
 * Inserts spaces into camelCased string
 *
 * @param {string} key
 * @return {string}
 */
export function camelToSentence(key) {
  return key.replace(/([A-Z])/g, ' $1');
}

/**
 * Check if a string contains/includes a substring
 *
 * @param {string} string
 * @param {string} substring
 * @return {boolean}
 */
export function includesSubstring(string, substring) {
  return string.toLowerCase().includes(substring.toLowerCase());
}


/**
 * Check if a string matches a substring
 *
 * @param {string} string
 * @param {string} substring
 * @return {boolean}
 */
export function stringsMatch(string, substring) {
  return string.toLowerCase() === substring.toLowerCase();
}
