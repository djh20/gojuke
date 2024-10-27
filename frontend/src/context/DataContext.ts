import { createContext } from "react";
import { Track } from "../api/track";

interface DataContextValue {
  tracks?: Track[] | null;
}

export const DataContext = createContext<DataContextValue>({});