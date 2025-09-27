import type { IconType } from "react-icons/lib"

export type NavMenuItemProps = {
  id: number
  name: string
  link: string
  icon: IconType
}

export interface NavigationProps {
  items: NavMenuItemProps[]
}