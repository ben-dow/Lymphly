import { AppShell } from "@mantine/core";
import "./App.css";
import Body from "./components/structure/body";
import Footer from "./components/structure/footer";
import Header from "./components/structure/header";

export default function App() {
  return (
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
        <Header />
      </AppShell.Header>

      <AppShell.Main className="h-full w-full">
        <Body />
      </AppShell.Main>

      <AppShell.Footer>
        <Footer />
      </AppShell.Footer>
    </AppShell>
  );
}
