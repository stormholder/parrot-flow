/**
 * Mock data for scenarios
 * Based on OpenAPI schema: ScenarioResponseItem
 */

import type { ScenarioResponseItem } from "@/shared/api-client";

export const mockScenarios: ScenarioResponseItem[] = [
  {
    id: "1",
    name: "Login Automation",
    description: "Automates login process for testing",
    tag: "authentication",
    icon: "🔐",
    created_at: "2024-01-15T10:00:00Z",
    updated_at: "2024-01-15T10:00:00Z",
  },
  {
    id: "2",
    name: "Form Submission",
    description: "Fills and submits contact form",
    tag: "forms",
    icon: "📝",
    created_at: "2024-01-16T11:30:00Z",
    updated_at: "2024-01-16T11:30:00Z",
  },
  {
    id: "3",
    name: "E-commerce Checkout",
    description: "Complete checkout flow including payment",
    tag: "e-commerce",
    icon: "🛒",
    created_at: "2024-01-17T09:15:00Z",
    updated_at: "2024-01-17T09:15:00Z",
  },
];

export const findScenarioById = (id: string): ScenarioResponseItem | undefined => {
  return mockScenarios.find((scenario) => scenario.id === id);
};

export const createScenario = (
  data: Partial<ScenarioResponseItem>
): ScenarioResponseItem => {
  return {
    id: Math.random().toString(36).substring(7),
    name: data.name || "New Scenario",
    description: data.description || "",
    tag: data.tag || "",
    icon: data.icon || "📄",
    created_at: new Date().toISOString(),
    updated_at: new Date().toISOString(),
  };
};
