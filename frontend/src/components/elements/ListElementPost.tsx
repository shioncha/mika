import DOMpurify from "dompurify";
import { Link, useLocation } from "react-router";

import { postService } from "../../libs/ContentService";
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

  const handleCheckboxChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    e.stopPropagation();
    postService
      .updateCheckbox(post.ID, e.currentTarget.checked)
      .then(() => {
        // Optionally handle success, e.g., show a toast or update state
      })
      .catch(() => {
        console.error("Failed to update checkbox state");
      });
  };

  return (
    <Link
      key={post.ID}
      className={style.post}
      to={`/posts/${post.ID}`}
      state={{ background: location }}
    >
      <div className={style.contentContainer}>
        {post.HasCheckbox ? (
          <input
            type="checkbox"
            defaultChecked={post.IsChecked}
            className={style.checkbox}
            onChange={handleCheckboxChange}
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
