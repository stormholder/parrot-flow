import { NavLink } from "react-router-dom";
import type { NavMenuItemProps } from "./types";

export const NavigationItem = ({
  id,
  name,
  link,
  icon: Icon,
}: NavMenuItemProps): React.ReactNode => {
  const baseClasses =
    "flex items-center cursor-pointer no-underline p-3 justify-center lg:justify-start text-black dark:text-white";
  return (
    <NavLink
      to={link}
      key={id}
      className={({ isActive }) => {
        return isActive
          ? `${baseClasses} font-semibold bg-black text-white dark:bg-white dark:text-black`
          : `${baseClasses} font-light`;
      }}
    >
      <i className="mr-2 text-xl">
        <Icon size={20} />
      </i>
      <span className="hidden lg:block">{name}</span>
    </NavLink>
  );
};
