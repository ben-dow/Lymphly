import { Box, MantineProvider } from "@mantine/core";
import "@mantine/core/styles.css";
import React from "react";
import "./tailwind.css";


export default function LayoutDefault({ children }: { children: React.ReactNode }) {
  return (
    <MantineProvider>
      <Box className="h-screen md:overflow-hidden">
        {children}
      </Box>
    </MantineProvider>
  );
}
