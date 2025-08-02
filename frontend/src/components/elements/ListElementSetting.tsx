import { FaChevronRight } from "react-icons/fa6";
import { Link } from "react-router";

import style from "../../styles/components/elements/ListElementSetting.module.css";

interface ListElementSettingProps
  extends React.HTMLAttributes<HTMLAnchorElement> {
  to: string;
  name: string;
}

function ListElementSetting({ children, to, name }: ListElementSettingProps) {
  return (
    <div className={style.item}>
      <Link className={style.link} to={`/settings/${to}`}></Link>
      <div className={style.content}>
        <span>{name}</span>
        <span className={style.value}>{children}</span>
      </div>
      <FaChevronRight className={style.arrow} />
    </div>
  );
}

export default ListElementSetting;
