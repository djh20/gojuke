import css from "./TrackList.module.css";
import { useContext } from "react";
import { DataContext } from "../context/DataContext";
import TrackListItem from "./TrackListItem";

function TrackList() {
  const data = useContext(DataContext);

  if (!data.tracks) return <></>;

  return (
    <div className={css.list}>
      {data.tracks.map(track => <TrackListItem key={track.id} track={track} />)}
    </div>
  );
}

export default TrackList;