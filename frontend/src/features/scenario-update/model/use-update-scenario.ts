import { useState } from "react";
import { scenarioApi } from "@/entities/scenario";
import type { ScenarioPatchRequestBody } from "@/shared/api-client";

export function useUpdateScenario() {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<Error | null>(null);

  const updateScenario = async (id: number, data: ScenarioPatchRequestBody) => {
    try {
      setLoading(true);
      setError(null);
      await scenarioApi.update(id, data);
    } catch (err) {
      const errorObj = err instanceof Error ? err : new Error("Failed to update scenario");
      setError(errorObj);
      throw errorObj;
    } finally {
      setLoading(false);
    }
  };

  return {
    updateScenario,
    loading,
    error,
  };
}
