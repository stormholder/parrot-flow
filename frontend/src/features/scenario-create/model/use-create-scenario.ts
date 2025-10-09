import { useState } from "react";
import { scenarioApi } from "@/entities/scenario";

export function useCreateScenario() {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<Error | null>(null);

  const createScenario = async () => {
    try {
      setLoading(true);
      setError(null);
      const result = await scenarioApi.create();
      return result;
    } catch (err) {
      const errorObj = err instanceof Error ? err : new Error("Failed to create scenario");
      setError(errorObj);
      throw errorObj;
    } finally {
      setLoading(false);
    }
  };

  return {
    createScenario,
    loading,
    error,
  };
}
