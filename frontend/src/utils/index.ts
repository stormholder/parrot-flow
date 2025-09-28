export const clsx = (...parts: string[]): string => parts.join(" ");
export const capitalize = (str: string) =>
  str.charAt(0).toUpperCase() + str.slice(1);
