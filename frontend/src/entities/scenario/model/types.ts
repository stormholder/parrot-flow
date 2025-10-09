export type {
  ScenarioResponseBody as Scenario,
  ScenarioListResponseBody as ScenarioList,
  ScenarioContext,
  ScenarioPatchRequestBody as ScenarioUpdateInput,
} from "@/shared/api-client";

export type { Node, Edge } from "@/shared/api-client";

export interface ScenarioFilters {
  name?: string;
  tags?: string[];
  page?: number;
  rpp?: number;
  order?: string;
  dir?: "asc" | "desc";
}
