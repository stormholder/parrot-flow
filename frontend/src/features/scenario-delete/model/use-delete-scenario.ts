import { useState } from "react";
import { scenarioApi } from "@/entities/scenario";

export function useDeleteScenario() {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<Error | null>(null);

  const deleteScenario = async (id: number) => {
    try {
      setLoading(true);
      setError(null);
      await scenarioApi.delete(id);
    } catch (err) {
      const errorObj = err instanceof Error ? err : new Error("Failed to delete scenario");
      setError(errorObj);
      throw errorObj;
    } finally {
      setLoading(false);
    }
  };

  return {
    deleteScenario,
    loading,
    error,
  };
}
