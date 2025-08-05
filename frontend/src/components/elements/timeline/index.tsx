import { useEffect, useState, useCallback } from "react";
import { useInView } from "react-intersection-observer";

import { postService, tagService } from "../../../libs/ContentService";
import { formatDate, groupByDate, localDate } from "../../../libs/datetime";
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
  const [posts, setPosts] = useState<Post[]>([]);
  const [nextCursor, setNextCursor] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [hasMore, setHasMore] = useState(true);

  useEffect(() => {
    let ignore = false;

    const loadInitialPosts = async () => {
      setIsLoading(true);
      try {
        const response = tag
          ? await tagService.fetchPostsByTag(tag, 20)
          : await postService.fetchPosts(20);

        if (!ignore) {
          setPosts(response.posts);
          if (response.next_cursor) {
            setNextCursor(response.next_cursor);
            setHasMore(true);
          } else {
            setNextCursor(null);
            setHasMore(false);
          }
        }
      } catch {
        console.error("Failed to load initial posts.");
      } finally {
        if (!ignore) {
          setIsLoading(false);
        }
      }
    };

    setFadeout(false);
    setPosts([]);
    setNextCursor(null);
    setHasMore(true);
    loadInitialPosts();

    return () => {
      ignore = true;
    };
  }, [tag, setFadeout]);

  const loadMorePosts = useCallback(async () => {
    if (isLoading || !hasMore || !nextCursor) return;

    setIsLoading(true);
    try {
      const response = tag
        ? await tagService.fetchPostsByTag(tag, 20, nextCursor)
        : await postService.fetchPosts(20, nextCursor);

      setPosts((prevPosts: Post[]) => [...prevPosts, ...response.posts]);

      if (response.next_cursor) {
        setNextCursor(response.next_cursor);
      } else {
        setHasMore(false);
        setNextCursor(null);
      }
    } catch (error) {
      console.error("Failed to load more posts:", error);
    } finally {
      setIsLoading(false);
    }
  }, [isLoading, hasMore, nextCursor, tag]);

  const { ref, inView } = useInView({
    threshold: 0,
  });

  useEffect(() => {
    if (inView && hasMore && !isLoading && posts.length > 0) {
      loadMorePosts();
    }
  }, [inView, hasMore, isLoading, posts.length, loadMorePosts]);

  const groupedPosts = groupByDate(posts);
  const sortedDates = Object.keys(groupedPosts).sort(
    (a, b) => new Date(b).getTime() - new Date(a).getTime()
  );

  const finalGroupedPosts: { [date: string]: Post[] } = {};
  sortedDates.forEach((date) => {
    const filtered = groupedPosts[date].filter(
      (post) => !onlyUnchecked || (!post.is_checked && post.has_checkbox)
    );
    if (filtered.length > 0) {
      finalGroupedPosts[date] = filtered;
    }
  });
  const finalSortedDates = Object.keys(finalGroupedPosts);

  return (
    <div className={fadeout ? style.fadeout : undefined}>
      {finalSortedDates.map((date) => (
        <div key={date} className={style.animation}>
          <span className={style.date}>{localDate(formatDate(date))}</span>
          <List className={style.posts}>
            {finalGroupedPosts[date].map((post) => (
              <ListElementPost key={post.ID} post={post} />
            ))}
          </List>
        </div>
      ))}

      {isLoading && <div className={style.loading}>Loading...</div>}
      {!isLoading && posts.length === 0 && (
        <div className={style.noPosts}>投稿はありません</div>
      )}
      {hasMore && !isLoading && <div ref={ref} style={{ height: "1px" }} />}
      {!hasMore && posts.length > 0 && (
        <div className={style.noPosts}>すべての投稿を読み込みました</div>
      )}
    </div>
  );
}

export default TimelineComponent;
