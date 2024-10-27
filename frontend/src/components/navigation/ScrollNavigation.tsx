import css from "./ScrollNavigation.module.css";
import ScrollNavigationButton from "./ScrollNavigationButton";

interface ScrollNavigationProps {
  scrollAmount: number;
  onScroll: (pos: number) => void;
}

function ScrollNavigation(props: ScrollNavigationProps) {
  return (
    <div className={css.navigation}>
      <ScrollNavigationButton direction="up" onClick={() => props.onScroll(-props.scrollAmount)} />
      <ScrollNavigationButton direction="down" onClick={() => props.onScroll(props.scrollAmount)} />
    </div>
  );
}

export default ScrollNavigation;