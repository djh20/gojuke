import css from "./TrackListItem.module.css"
import { playTrack, Track } from "../api/track";
import { BsPlayCircleFill } from "react-icons/bs";
import { useState } from "react";

interface TrackListItemProps {
  track: Track;
}

function TrackListItem(props: TrackListItemProps) {
  const [ isBlocked, setBlocked ] = useState(false);

  const play = () => {
    if (isBlocked) return;

    setBlocked(true);
    playTrack(props.track);

    setTimeout(() => setBlocked(false), 500);
  }

  return (
    <div className={css.item} onClick={play}>
      <span className={css.title}>{props.track.name}</span>
      <div className={[css.playButton, isBlocked ? css.disabled : null].join(' ')}>
        <BsPlayCircleFill className={css.playIcon} />
        <span className={css.playLabel}>PLAY</span>
      </div>
    </div>
  )
}

export default TrackListItem;