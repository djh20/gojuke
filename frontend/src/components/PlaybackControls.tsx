import { BsFillSkipForwardFill, BsVolumeDownFill, BsVolumeUpFill } from "react-icons/bs";
import PlaybackControlButton from "./PlaybackControlButton";
import css from "./PlaybackControls.module.css";
import { skipTrack } from "../api/track";
import { useState } from "react";
import { clamp } from "../utils";
import { setRemoteVolume } from "../api/volume";

function PlaybackControls() {
  const [volume, setVolume] = useState(100);

  const adjustVolume = (amount: number) => {
    const newVolume = clamp(volume + amount, 0, 200)

    setVolume(newVolume);
    setRemoteVolume(newVolume);
  }

  return (
    <div className={css.controls}>
      <PlaybackControlButton 
        color="#f5c118" 
        icon={<BsFillSkipForwardFill />} 
        onClick={skipTrack}>
        Skip
      </PlaybackControlButton>

      <PlaybackControlButton 
        color="#3dcaf5" 
        icon={<BsVolumeUpFill />}
        onClick={() => adjustVolume(10)}>
        Vol +
      </PlaybackControlButton>

      <PlaybackControlButton 
        color="#3dcaf5" 
        icon={<BsVolumeDownFill />}
        onClick={() => adjustVolume(-10)}>
        Vol -
      </PlaybackControlButton>
    </div>
  );
}

export default PlaybackControls;