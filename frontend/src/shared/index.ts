/**
 * Shared Layer
 * Business-agnostic infrastructure code
 *
 * Segments:
 * - api-client: Auto-generated OpenAPI client (DO NOT EDIT)
 * - config: Application configuration and environment variables
 * - lib: Utility functions, custom hooks, and helpers
 * - ui: Reusable UI components
 */

// Re-export commonly used modules
export * from "./config";
export * from "./lib";
export * from "./ui";

// API client is intentionally not re-exported here
// Import directly: import { ScenarioService } from "@/shared/api-client"
