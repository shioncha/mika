import { useContext, useEffect, useState } from "react";
import { Link, useNavigate } from "react-router";

import { AuthContext } from "../../../hooks/auth_context";
import { TagsAPI } from "../../../libs/api";
import type { Tag } from "../../../type/tag";
import style from "./style.module.css";

function TagsComponent({
  tag,
  setFadeout,
}: {
  tag?: string;
  setFadeout: (fadeout: boolean) => void;
}) {
  const { token } = useContext(AuthContext);
  const [tags, setTags] = useState<Tag[]>([]);
  const navigate = useNavigate();

  useEffect(() => {
    if (!token) {
      return;
    }
    TagsAPI({ method: "GET", token, setTags });
  }, [token]);

  function handleLinkClick(e: React.MouseEvent<HTMLAnchorElement>) {
    e.preventDefault();
    const href = e.currentTarget.getAttribute("href")!;
    if (tag ? href == `/tags/${tag}` : href == "/") return;
    setFadeout(true);
    setTimeout(() => {
      navigate(href);
    }, 200);
  }

  return (
    <div className={style.tags}>
      <Link
        to="/"
        onClick={handleLinkClick}
        className={`${style.tag} ${tag ?? style.active}`}
      >
        All
      </Link>
      {tags != null &&
        tags.map((t) => (
          <Link
            key={t.ID}
            to={`/tags/${t.Name}`}
            onClick={handleLinkClick}
            className={`${style.tag} ${t.Name == tag ? style.active : ""}`}
          >
            #{t.Name}
          </Link>
        ))}
    </div>
  );
}

export default TagsComponent;
