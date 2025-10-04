import { customAlphabet } from "nanoid/non-secure";

export const clsx = (...parts: string[]): string => parts.join(" ");

export const isAppleOS = () =>
  window.navigator.platform.startsWith("Mac") ||
  window.navigator.platform.startsWith("iPhone") ||
  window.navigator.platform.startsWith("iPad") ||
  window.navigator.platform.startsWith("iPod");

export const capitalize = (str: string) =>
  str.charAt(0).toUpperCase() + str.slice(1);

export const randomBetween = (min: number, max: number): number => {
  return Math.floor(Math.random() * (max - min + 1) + min);
};

export const nanoid = () =>
  customAlphabet(
    "X36pBZhmqYIVd217_tDEgLxRUosk9AzSjyWOCfJePrwMlTNuKHbv8c05aiQ4GFn"
  )();
