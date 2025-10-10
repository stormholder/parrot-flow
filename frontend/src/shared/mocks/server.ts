/**
 * MirageJS Server Configuration
 * In-memory API mock server without Service Workers
 */

import { createServer, Model, Response } from "miragejs";
import { mockScenarios } from "./data/scenarios";

export function makeServer({ environment = "development" } = {}) {
  console.log("[Mirage] Creating server in", environment, "mode");

  const server = createServer({
    environment,

    models: {
      scenario: Model,
    },

    seeds(server) {
      console.log("[Mirage] Seeding database with", mockScenarios.length, "scenarios");
      mockScenarios.forEach((scenario) => {
        server.create("scenario", scenario);
      });
    },

    routes() {
      // Set namespace to match your API
      this.namespace = "api";

      // Logging
      this.logging = true;

      // Timing - simulate network delay (optional, set to 0 for instant)
      this.timing = 400;

      // GET /api/scenarios/ - List scenarios
      this.get("/scenarios/", (schema, request) => {
        console.log("[Mirage] GET /api/scenarios/", request.queryParams);
        const queryParams = request.queryParams;
        const name = typeof queryParams.name === 'string' ? queryParams.name : undefined;
        const tag = typeof queryParams.tag === 'string' ? queryParams.tag : undefined;
        const page = typeof queryParams.page === 'string' ? queryParams.page : "1";
        const rpp = typeof queryParams.rpp === 'string' ? queryParams.rpp : "10";

        let scenarios = schema.all("scenario").models as any[];

        // Apply filters
        if (name) {
          scenarios = scenarios.filter((s: any) =>
            s.attrs.name.toLowerCase().includes(name.toLowerCase())
          );
        }
        if (tag) {
          scenarios = scenarios.filter((s: any) => s.attrs.tag === tag);
        }

        // Pagination
        const pageNum = parseInt(page, 10);
        const rppNum = parseInt(rpp, 10);
        const start = (pageNum - 1) * rppNum;
        const end = start + rppNum;
        const paginatedData = scenarios.slice(start, end).map((s: any) => s.attrs);

        return {
          data: paginatedData,
          total: scenarios.length,
          page: pageNum,
          rpp: rppNum,
        };
      });

      // POST /api/scenarios/ - Create scenario
      this.post("/scenarios/", (schema, request) => {
        console.log("[Mirage] POST /api/scenarios/");
        const attrs = JSON.parse(request.requestBody);
        const now = new Date().toISOString();

        const newScenario = schema.create("scenario", {
          ...attrs,
          id: Math.random().toString(36).substring(7),
          description: attrs.description || "",
          tag: attrs.tag || "",
          icon: attrs.icon || "📄",
          created_at: now,
          updated_at: now,
        });

        return newScenario;
      });

      // GET /api/scenarios/:id - Get scenario by ID
      this.get("/scenarios/:id", (schema, request) => {
        const { id } = request.params;
        console.log("[Mirage] GET /api/scenarios/" + id);

        const scenario = schema.find("scenario", id);

        if (!scenario) {
          return new Response(
            404,
            {},
            {
              status: 404,
              title: "Not Found",
              detail: `Scenario with id ${id} not found`,
            }
          );
        }

        return scenario;
      });

      // PATCH /api/scenarios/:id - Update scenario
      this.patch("/scenarios/:id", (schema, request) => {
        const { id } = request.params;
        console.log("[Mirage] PATCH /api/scenarios/" + id);

        const attrs = JSON.parse(request.requestBody);
        const scenario = schema.find("scenario", id);

        if (!scenario) {
          return new Response(
            404,
            {},
            {
              status: 404,
              title: "Not Found",
              detail: `Scenario with id ${id} not found`,
            }
          );
        }

        scenario.update({
          ...attrs,
          updated_at: new Date().toISOString(),
        });

        const updated = scenario.attrs as any;

        return {
          id: updated.id,
          name: updated.name,
          description: updated.description,
          tag: updated.tag,
          icon: updated.icon,
          updated_at: updated.updated_at,
        };
      });

      // DELETE /api/scenarios/:id - Delete scenario
      this.delete("/scenarios/:id", (schema, request) => {
        const { id } = request.params;
        console.log("[Mirage] DELETE /api/scenarios/" + id);

        const scenario = schema.find("scenario", id);

        if (!scenario) {
          return new Response(
            404,
            {},
            {
              status: 404,
              title: "Not Found",
              detail: `Scenario with id ${id} not found`,
            }
          );
        }

        scenario.destroy();

        return {
          success: true,
        };
      });

      // Runs endpoints (placeholder for future)
      this.get("/runs", () => ({ data: [], total: 0, page: 1, rpp: 10 }));
      this.post("/runs", () => ({}));
      this.get("/runs/:id", () => ({}));
      this.post("/runs/:id/start", () => ({}));

      // Pass through for any unhandled requests (like assets)
      this.passthrough((request) => {
        // Don't pass through API calls
        if (request.url.includes("/api/")) {
          return false;
        }
        // Pass through everything else (Vite HMR, assets, etc.)
        return true;
      });
    },
  });

  console.log("[Mirage] Server created and ready!");
  return server;
}
