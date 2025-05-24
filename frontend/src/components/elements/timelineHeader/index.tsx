import { CiCalendar, CiCircleCheck } from "react-icons/ci";

import style from "./style.module.css";

interface Props {
  onlyUnchecked: boolean;
  setOnlyUnchecked: (onlyChecked: boolean) => void;
}

function TimelineHeaderComponent(props: Props) {
  const { onlyUnchecked, setOnlyUnchecked } = props;

  return (
    <div className={style.header}>
      <h1>Timeline</h1>
      <div className={style.icons}>
        <button className={style.icon} aria-label="日付を選択する">
          <CiCalendar color="white" size="1.5rem" />
        </button>
        <button
          className={style.icon}
          aria-label="未チェックのみを表示する"
          onClick={() => setOnlyUnchecked(!onlyUnchecked)}
        >
          <CiCircleCheck color="white" size="1.5rem" />
        </button>
      </div>
    </div>
  );
}

export default TimelineHeaderComponent;
