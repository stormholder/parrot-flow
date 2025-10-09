import { RouterProvider } from "react-router-dom";
import { HeroUIProvider } from "@heroui/react";
import { router } from "./router/router";
import { AppLayout } from "@/shared";

/**
 * Root application component
 * Provides UI framework, routing, and layout for the entire application
 */
const App = () => {
  return (
    <HeroUIProvider>
      <AppLayout>
        <RouterProvider router={router} />
      </AppLayout>
    </HeroUIProvider>
  );
};

export default App;
