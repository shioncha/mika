import { Link } from "react-router";

import style from "../../styles/components/elements/ListElementSetting.module.css";

interface ListElementSettingProps
  extends React.HTMLAttributes<HTMLAnchorElement> {
  to: string;
}

function ListElementSetting({ children, to }: ListElementSettingProps) {
  return (
    <Link className={style.item} to={`/settings/${to}`}>
      {children}
    </Link>
  );
}

export default ListElementSetting;
