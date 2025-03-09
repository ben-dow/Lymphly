import { AppShell, MantineProvider } from "@mantine/core";
import "@mantine/core/styles.css";
import { useDisclosure } from "@mantine/hooks";
import React from "react";
import "./tailwind.css";

import HeadDefault from "../pages/+Head";

export default function LayoutDefault({ children }: { children: React.ReactNode }) {
  const [opened, { toggle }] = useDisclosure();
  return (
    <MantineProvider>
      <AppShell
        withBorder={false}
        header={{ height: 100 }}
        padding="md"
        className="bg-sky-200 sm:h-screen  w-full"
        footer={{
          height: 100,
        }}
      >
        <AppShell.Header>
          <HeadDefault />
        </AppShell.Header>

        <AppShell.Main>
          {children}
        </AppShell.Main>

        <AppShell.Footer>
          <div className="bg-sky-900 h-full w-full"></div>
        </AppShell.Footer>
      </AppShell>
    </MantineProvider>
  );
}
