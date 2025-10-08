import type { FC, HTMLAttributes } from "react";

import type { NavigationProps } from "./types";
import { NavigationItem } from "./NavigationItem";

export const Navigation: FC<
  NavigationProps & HTMLAttributes<HTMLDivElement>
> = ({ items }) => {
  const menuItems = items.map((item) => (
    <li key={item.id}>
      <NavigationItem
        id={item.id}
        link={item.link}
        name={item.name}
        icon={item.icon}
      />
    </li>
  ));

  return (
    <div className="h-screen flex-shrink-0 static left-0 top-0 z-50 bg-white dark:bg-neutral-800 border-r border-gray-700 w-[5rem] lg:w-[10rem] select-none">
      <nav className="flex flex-col h-[calc(100vh - 50px)]">
        <div className="overflow-y-auto pt-2 w-full">
          <ul className="list-none p-0 m-0">{menuItems}</ul>
        </div>
      </nav>
    </div>
  );
};
