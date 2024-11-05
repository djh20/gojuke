import css from "./ScrollNavigation.module.css";
import ScrollNavigationButton from "./ScrollNavigationButton";

interface ScrollNavigationProps {
  onScroll: (direction: number) => void;
}

function ScrollNavigation(props: ScrollNavigationProps) {
  return (
    <div className={css.navigation}>
      <ScrollNavigationButton direction="up" onClick={() => props.onScroll(-1)} />
      <ScrollNavigationButton direction="down" onClick={() => props.onScroll(1)} />
    </div>
  );
}

export default ScrollNavigation;