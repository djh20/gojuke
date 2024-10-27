import { useQuery } from "@tanstack/react-query";
import { fetchTracks, Track } from "../api/track";
import { DataContext } from "../context/DataContext";
import { ReactElement } from "react";

interface DataProviderProps {
  children?: ReactElement;
}

function DataProvider(props: DataProviderProps) {
  const {
    data: tracks,
  } = useQuery<Track[]>({
    queryKey: ["tracks"],
    queryFn: fetchTracks
  });

  return (
    <DataContext.Provider value={{tracks}}>
      {props.children}
    </DataContext.Provider>
  );
}

export default DataProvider;