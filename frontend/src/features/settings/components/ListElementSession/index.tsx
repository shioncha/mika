import type { IconType } from "react-icons";
import { CiDesktop, CiMobile2 } from "react-icons/ci";

import style from "./ListElementSession.module.css";

function ListElementSession({
  userAgent,
  ipAddress,
}: {
  userAgent: string;
  ipAddress: string;
}) {
  const { name: deviceName, icon: Icon } = findDeviceByUserAgent(userAgent);

  return (
    <div className={style.item}>
      <div>
        <Icon size={32} style={{ marginRight: "10px" }} />
      </div>
      <div>
        <div>{deviceName}</div>
        <div style={{ fontSize: "0.9em", color: "#666" }}>{ipAddress}</div>
      </div>
    </div>
  );
}

function findDeviceByUserAgent(userAgent: string): {
  name: string;
  icon: IconType;
} {
  const devices = [
    { name: "iOS", icon: CiMobile2, pattern: /iPhone|iPad|iPod/ },
    { name: "Android", icon: CiMobile2, pattern: /Android/ },
    { name: "Mac", icon: CiDesktop, pattern: /Macintosh/ },
    { name: "Windows", icon: CiDesktop, pattern: /Windows/ },
    { name: "Linux", icon: CiDesktop, pattern: /Linux/ },
  ];

  for (const device of devices) {
    if (device.pattern.test(userAgent)) {
      return { name: device.name, icon: device.icon };
    }
  }

  return { name: "Unknown Device", icon: CiDesktop };
}

export default ListElementSession;
