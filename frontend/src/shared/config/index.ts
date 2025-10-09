/**
 * Application configuration
 * Centralizes environment variables and app constants
 */

/**
 * API base URL - defaults to localhost for development
 */
export const API_BASE_URL =
  import.meta.env.VITE_API_BASE_URL || "http://localhost:8080";

/**
 * Application name
 */
export const APP_NAME = "Parrot Flow";

/**
 * API endpoints configuration
 */
export const API_ENDPOINTS = {
  SCENARIOS: "/api/v1/scenarios",
  RUNS: "/api/v1/runs",
} as const;
