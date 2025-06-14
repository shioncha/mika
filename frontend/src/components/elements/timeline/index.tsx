import { useContext, useEffect, useState } from "react";

import { AuthContext } from "../../../hooks/auth_context";
import { PostsAPI, TagsPostsAPI } from "../../../libs/api";
import { groupByDate } from "../../../libs/datetime";
import type { Post } from "../../../type/post";
import List from "../List";
import ListElementPost from "../ListElementPost";
import style from "./style.module.css";

interface TimelineProps {
  tag?: string;
  fadeout: boolean;
  setFadeout: (fadeout: boolean) => void;
  onlyUnchecked: boolean;
}

function TimelineComponent({
  tag,
  fadeout,
  setFadeout,
  onlyUnchecked,
}: TimelineProps) {
  const { token } = useContext(AuthContext);
  const [posts, setPosts] = useState<Post[]>([]);

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
  }, [token, tag, setFadeout]);

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
            <List key={index} className={style.uncheckedPost}>
              {posts.map((post) => (
                <ListElementPost key={post.ID} post={post} />
              ))}
            </List>
          ))}
        </div>
      )}
      {sortedDates.map((date) => (
        <div key={date} className={style.animation}>
          <span className={style.date}>{date}</span>
          <List className={style.posts}>
            {groupedPosts[date].length === 0 ? (
              <div className={style.noPosts}>投稿はありません</div>
            ) : (
              groupedPosts[date].map((post) => (
                <ListElementPost key={post.ID} post={post} />
              ))
            )}
          </List>
        </div>
      ))}
    </div>
  );
}

export default TimelineComponent;
