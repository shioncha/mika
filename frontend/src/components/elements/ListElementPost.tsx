import DOMpurify from "dompurify";
import { Link, useLocation } from "react-router";

import { formatDate, localTime } from "../../libs/datetime";
import style from "../../styles/components/elements/ListElementPost.module.css";
import type { Post } from "../../type/post";

interface ListElementPostProps {
  post: Post;
}

// Utility functions
function addLink(str: string): string {
  const url = /((h?)(ttps?:\/\/[a-zA-Z0-9.\-_@:/~?%&;=+#',()*!]+))/g;
  return str.replace(url, (_match, url, _h, href) => {
    return `<a href="h${href}" target="_blank" rel="noopener noreferrer">${url}</a>`;
  });
}

function highlightTag(str: string): string {
  const tagPattern = /#[^\s#]+/g;
  return str.replace(tagPattern, (match) => {
    return `<span class="${style.tag}">${match}</span>`;
  });
}

function richText(str: string): string {
  return addLink(highlightTag(str));
}

// DOMpurify options
const sanitizeOptions = {
  ADD_TAGS: ["a"],
  ADD_ATTR: ["target"],
};

function ListElementPost({ post }: ListElementPostProps) {
  const location = useLocation();

  return (
    <Link
      key={post.ID}
      className={style.post}
      to={`/posts/${post.ID}`}
      state={{ background: location }}
    >
      <div className={style.contentContainer}>
        {post.has_checkbox ? (
          <input
            type="checkbox"
            defaultChecked={post.is_checked}
            className={style.checkbox}
            onClick={(e) => e.stopPropagation()}
          />
        ) : null}
        <p
          className={style.content}
          dangerouslySetInnerHTML={{
            __html: DOMpurify.sanitize(richText(post.Content), sanitizeOptions),
          }}
          onClick={(e) => {
            if ((e.target as HTMLElement).tagName === "A") {
              e.stopPropagation();
            }
          }}
        />
      </div>
      <span className={style.time} onClick={(e) => e.stopPropagation()}>
        {localTime(formatDate(post.CreatedAt))}
      </span>
    </Link>
  );
}

export default ListElementPost;
