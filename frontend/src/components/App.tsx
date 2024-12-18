import '@fontsource/open-sans/400.css';
import '@fontsource/open-sans/600.css';
import '@fontsource/open-sans/700.css';
import TrackList from './TrackList';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import DataProvider from './DataProvider';
import PlaybackControls from './playback/PlaybackControls';

const queryClient = new QueryClient();

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <DataProvider>
        <h1>Jukebox</h1>
        <PlaybackControls />
        <TrackList />
      </DataProvider>
    </QueryClientProvider>
  );
}

export default App;
