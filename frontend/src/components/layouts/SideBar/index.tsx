import "react-day-picker/style.css";

import { DayPicker } from "react-day-picker";

import style from "./SideBar.module.css";

function SideBar() {
  return (
    <div className={style.sidebar}>
      <div className={style.element}>
        <DayPicker
          animate
          mode="single"
          captionLayout="dropdown"
          className={style.dayPicker}
        />
      </div>
    </div>
  );
}

export default SideBar;
