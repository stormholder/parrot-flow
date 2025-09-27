import type { FC } from "react";
import { Outlet } from "react-router-dom";

import { Navigation } from "@/components/Navigation/Navigation";
import routes from "@/routes/root";

export const RootPage: FC = () => {
  return (
    <>
      <Navigation
        items={routes}
        style={{
          flex: 1,
        }}
      />
      <div id="detail" className="flex flex-1">
        <Outlet />
      </div>
    </>
  );
};
