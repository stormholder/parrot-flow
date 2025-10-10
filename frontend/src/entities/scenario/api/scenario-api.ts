import { ScenariosService } from "@/shared/api-client";
import type {
  ScenarioResponseItem,
  ListScenariosResponseBody,
  UpdateScenarioRequestBody,
  CreateScenarioResponseBody,
} from "@/shared/api-client";

export const scenarioApi = {
  async getList(filters?: {
    name?: string;
    tags?: string[];
    page?: number;
    rpp?: number;
    order?: string;
    dir?: string;
  }): Promise<ListScenariosResponseBody> {
    const response = await ScenariosService.listScenarios(
      filters?.name,
      undefined, // filters?.tag,
      filters?.page,
      filters?.rpp
      // filters?.order,
      // filters?.dir
    );
    return response as ListScenariosResponseBody;
  },

  async getById(id: string): Promise<ScenarioResponseItem> {
    const response = await ScenariosService.getScenario(id);
    return response as ScenarioResponseItem;
  },

  async create(name: string): Promise<CreateScenarioResponseBody> {
    const response = await ScenariosService.createScenario({
      name,
    });
    return response as CreateScenarioResponseBody;
  },

  async update(id: string, data: UpdateScenarioRequestBody): Promise<void> {
    await ScenariosService.updateScenario(id, data);
  },

  async delete(id: string): Promise<void> {
    await ScenariosService.deleteScenario(id);
  },
};
