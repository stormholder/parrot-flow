import type { FC } from "react";
import { Container, Flex } from "@radix-ui/themes";

import type { NavigationProps } from "./types";
import { NavigationItem } from "./NavigationItem";

export const Navigation: FC<NavigationProps> = ({ items }) => {
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
    <Container size={"1"}>
      <Flex direction={"column"} gap={"2"}>
        <ul>{menuItems}</ul>
      </Flex>
    </Container>
  );
};
