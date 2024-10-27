import css from "./TrackList.module.css";
import { useContext, useRef } from "react";
import { DataContext } from "../context/DataContext";
import TrackListItem from "./TrackListItem";
import ScrollNavigation from "./navigation/ScrollNavigation";

function TrackList() {
  const data = useContext(DataContext);

  const listRef = useRef<HTMLDivElement | null>(null);

  const scroll = (amount: number) => {
    if (!listRef.current) return;

    listRef.current.scrollBy({top: amount})
  }

  if (!data.tracks) return <></>;

  return (
    <div className={css.list}>
      <div className={css.tracks} ref={listRef}>
        {data.tracks.map(track => <TrackListItem key={track.id} track={track} />)}
      </div>
      <ScrollNavigation scrollAmount={500} onScroll={scroll} />
    </div>
  );
}

export default TrackList;