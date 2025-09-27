import { useState } from "react";

const DashboardPage = () => {
  const [count, setCount] = useState(0);
  return (
    <>
      {/* <Title>Vite + React + Radix</Title>
      <Card>
        <Flex gap={"3"} direction={"column"} align={"center"}>
          <Button
            type={"primary"}
            onClick={() => setCount((count) => count + 1)}
          >
            Count is {count}
          </Button>
          <Text>
            Edit <Code>src/App.tsx</Code> and save to test HMR
          </Text>
        </Flex>
      </Card> */}
    </>
  );
};

export default DashboardPage;
