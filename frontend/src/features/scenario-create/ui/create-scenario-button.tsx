import { Button } from "@heroui/react";
import { useCreateScenario } from "../model";

interface CreateScenarioButtonProps {
  onSuccess?: (id: number) => void;
  children?: React.ReactNode;
}

export const CreateScenarioButton = ({ onSuccess, children }: CreateScenarioButtonProps) => {
  const { createScenario, loading } = useCreateScenario();

  const handleCreate = async () => {
    try {
      const result = await createScenario();
      if (result?.id && onSuccess) {
        onSuccess(result.id);
      }
    } catch (error) {
      // Error is handled by the hook
      console.error("Failed to create scenario:", error);
    }
  };

  return (
    <Button onClick={handleCreate} isLoading={loading}>
      {children || "Create Scenario"}
    </Button>
  );
};
