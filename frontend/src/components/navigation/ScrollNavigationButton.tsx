import { BsArrowDown, BsArrowUp } from "react-icons/bs";
import css from "./ScrollNavigationButton.module.css";

interface ScrollNavigationButtonProps {
  direction: "up" | "down";
  onClick?: () => void;
}

function ScrollNavigationButton(props: ScrollNavigationButtonProps) {
  return (
    <div className={css.button} onClick={props.onClick}>
      {props.direction == "up" ? <BsArrowUp/> : <BsArrowDown/> }
    </div>
  );
}

export default ScrollNavigationButton;