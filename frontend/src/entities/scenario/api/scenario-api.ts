import { ScenarioService } from "@/shared/api-client";
import type {
  ScenarioResponseBody,
  ScenarioListResponseBody,
  ScenarioPatchRequestBody,
  ScenarioCreateResponseBody,
} from "@/shared/api-client";

export const scenarioApi = {
  async getList(filters?: {
    name?: string;
    tags?: string[];
    page?: number;
    rpp?: number;
    order?: string;
    dir?: string;
  }): Promise<ScenarioListResponseBody> {
    const response = await ScenarioService.getScenarios(
      filters?.name,
      filters?.tags,
      filters?.page,
      filters?.rpp,
      filters?.order,
      filters?.dir
    );
    return response as ScenarioListResponseBody;
  },

  async getById(id: number): Promise<ScenarioResponseBody> {
    const response = await ScenarioService.getScenario(id);
    return response as ScenarioResponseBody;
  },

  async create(): Promise<ScenarioCreateResponseBody> {
    const response = await ScenarioService.createScenario();
    return response as ScenarioCreateResponseBody;
  },

  async update(id: number, data: ScenarioPatchRequestBody): Promise<void> {
    await ScenarioService.updateScenario(id, data);
  },

  async delete(id: number): Promise<void> {
    await ScenarioService.deleteScenario(id);
  },
};
