import { AppShell} from '@mantine/core';
import Header from './components/structure/header';
import Body from './components/structure/body';
import Footer from './components/structure/footer';
import "./App.css"

export default function App() {
  return (
    <AppShell
      header={{ height: 100 }}
      padding="md"
      className="bg-emerald-50"
    >
      <AppShell.Header>
        <Header/>
      </AppShell.Header>

      <AppShell.Main>
        <Body/>
      </AppShell.Main>

      <AppShell.Footer>
        <Footer/>
      </AppShell.Footer>
    </AppShell>
  );
}
