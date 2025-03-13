import { Box } from "@mantine/core";
import "radar-sdk-js/dist/radar.css";
import { DataDisplay } from "../../components/search/display";

export default function Search() {
  return (
    <Box className={"w-full h-full"}>
      <DataDisplay />
    </Box>
  );
}
