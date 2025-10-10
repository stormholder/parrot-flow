/**
 * Scenario Entity
 *
 * Business entity representing a browser automation scenario.
 * Contains API calls, data models, and reusable UI components.
 *
 * Public API - follow FSD conventions:
 * - Export only what's needed by other layers
 * - Keep internal implementation details private
 */

export { scenarioApi } from "./api/scenario-api";
export {
  scenarioListLoader,
  scenarioItemLoader,
  type ScenarioListLoaderData,
  type ScenarioItemLoaderData,
} from "./api/loaders";
export * from "./model";
