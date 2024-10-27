import { ReactElement, ReactNode } from "react";
import css from "./PlaybackControlButton.module.css";

interface PlaybackControlButtonProps {
  children?: ReactNode;
  icon: ReactElement;
  color: string;
  onClick?: () => void;
}

function PlaybackControlButton(props: PlaybackControlButtonProps) {
  return (
    <div className={css.button} style={{borderColor: props.color}} onClick={props.onClick}>
      {props.icon}
      <div className={css.label}>{props.children}</div>
    </div>
  );
}

export default PlaybackControlButton;