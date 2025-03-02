import { Box } from "@mantine/core";
import "radar-sdk-js/dist/radar.css";
import { DataDisplay } from "../../components/search/display";

export default function Search() {
  return (
    <Box className={"p-5 w-full h-full flex flex-col gap-5"}>
      <Box className="flex justify-center">
        <DataDisplay />
      </Box>
    </Box>
  );
}
