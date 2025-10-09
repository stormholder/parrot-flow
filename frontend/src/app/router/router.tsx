import { RootPage } from "@/pages/root";
import { createBrowserRouter } from "react-router-dom";

import { ErrorPage } from "@/pages/error";
import { DashboardPage } from "@/pages/dashboard";
import { scenarioItemLoader, scenarioListLoader } from "@/entities/scenario";
import { ScenarioListPage } from "@/pages/scenario-list";
import { ScenarioItemPage } from "@/pages/scenario-item";
import { ScenarioFlowPage } from "@/pages/scenario-flow";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <RootPage />,
    errorElement: <ErrorPage />,
    children: [
      {
        index: true,
        element: <DashboardPage />,
      },
      {
        path: "scenarios",
        loader: scenarioListLoader,
        element: <ScenarioListPage />,
      },
      {
        path: "scenarios/:scenarioId",
        loader: scenarioItemLoader,
        element: <ScenarioItemPage />,
      },
      {
        path: "scenarios/:scenarioId/flow",
        loader: scenarioItemLoader,
        element: <ScenarioFlowPage />,
      },
    ],
  },
]);
