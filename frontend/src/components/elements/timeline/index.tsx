import DOMpurify from "dompurify";
import { useContext, useEffect, useState } from "react";
import { useLocation, useNavigate } from "react-router";
import { Link } from "react-router";

import { AuthContext } from "../../../hooks/auth_context";
import { PostsAPI, TagsPostsAPI } from "../../../libs/api";
import { groupByDate, localTime } from "../../../libs/datetime";
import type { Post } from "../../../type/post";
import style from "./style.module.css";

interface TimelineProps {
  tag?: string;
  fadeout: boolean;
  setFadeout: (fadeout: boolean) => void;
  onlyUnchecked: boolean;
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

// Timeline component
function TimelineComponent({
  tag,
  fadeout,
  setFadeout,
  onlyUnchecked,
}: TimelineProps) {
  const { token } = useContext(AuthContext);
  const [posts, setPosts] = useState<Post[]>([]);
  const navigate = useNavigate();
  const location = useLocation();

  useEffect(() => {
    setFadeout(false);
    setPosts([]);
    if (token == null) {
      return;
    }
    if (!tag) {
      PostsAPI({ method: "GET", token, setPosts });
      return;
    }
    TagsPostsAPI({ method: "GET", tag, token, setPosts });
  }, [token, tag, navigate, setFadeout]);

  const groupedPosts = groupByDate(posts);
  const sortedDates = Object.keys(groupedPosts).sort(
    (a, b) => new Date(b).getTime() - new Date(a).getTime()
  );

  const filteredPosts = sortedDates.map((date) =>
    groupedPosts[date].filter(
      (post) => !onlyUnchecked || (!post.is_checked && post.has_checkbox)
    )
  );

  return (
    <div className={fadeout ? style.fadeout : undefined}>
      {onlyUnchecked && (
        <div className={style.unchecked}>
          {filteredPosts.length}件の未完了のタスクがあります
          {filteredPosts.map((posts, index) => (
            <div key={index} className={style.uncheckedPost}>
              {posts.map((post) => (
                <Link
                  key={post.ID}
                  className={style.post}
                  to={`/posts/${post.ID}`}
                  state={{ background: location }}
                >
                  <p
                    className={style.content}
                    dangerouslySetInnerHTML={{
                      __html: DOMpurify.sanitize(
                        richText(post.Content),
                        sanitizeOptions
                      ),
                    }}
                  />
                </Link>
              ))}
            </div>
          ))}
        </div>
      )}
      {sortedDates.map((date) => (
        <div key={date} className={style.animation}>
          <span className={style.date}>{date}</span>
          <div className={style.posts}>
            {groupedPosts[date].map((post) => (
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
                      __html: DOMpurify.sanitize(
                        richText(post.Content),
                        sanitizeOptions
                      ),
                    }}
                    onClick={(e) => {
                      if ((e.target as HTMLElement).tagName === "A") {
                        e.stopPropagation();
                      }
                    }}
                  />
                </div>
                <span
                  className={style.time}
                  onClick={(e) => e.stopPropagation()}
                >
                  {localTime(post.CreatedAt)}
                </span>
              </Link>
            ))}
          </div>
        </div>
      ))}
    </div>
  );
}

export default TimelineComponent;
