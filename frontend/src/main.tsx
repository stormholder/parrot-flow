import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import App from "@/app";
import "@/app/styles/global.css";
import { OpenAPI } from "@/shared/api-client";
import { API_BASE_URL } from "@/shared/config";
import { makeServer } from "@/shared/mocks/server";

// Configure OpenAPI client base URL
OpenAPI.BASE = API_BASE_URL;

// Start Mirage server if enabled
if (import.meta.env.VITE_ENABLE_MSW === "true") {
  console.log("[Mirage] Starting mock server...");
  makeServer({ environment: "development" });
} else {
  console.log("[Mirage] Mock server disabled, using real API");
}

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <App />
  </StrictMode>
);
