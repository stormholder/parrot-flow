import type { IconType } from "react-icons/lib";
import {
  TbHome,
  TbSettingsAutomation,
  TbSubtask,
  TbHistory,
} from "react-icons/tb";

type ApplicationRoute = {
  id: number;
  name: string;
  link: string;
  icon: IconType;
};

const routes: ApplicationRoute[] = [
  {
    id: 0,
    name: "Home",
    link: "/",
    icon: TbHome,
  },
  {
    id: 1,
    name: "Scenarios",
    link: "/scenarios",
    icon: TbSubtask,
  },
  {
    id: 2,
    name: "Agents",
    link: "/agents",
    icon: TbSettingsAutomation,
  },
  {
    id: 3,
    name: "Run history",
    link: "/history",
    icon: TbHistory,
  },
];

export default routes;
