import { useState, useCallback } from "react";
import { postService } from "../libs/ContentService";
import type { Post } from "../type/post";

export function useInfinitePosts() {
  const [posts, setPosts] = useState<Post[]>([]);
  const [nextCursor, setNextCursor] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [hasMore, setHasMore] = useState(true);

  const loadMorePosts = useCallback(async () => {
    if (isLoading || !hasMore) return;

    setIsLoading(true);
    try {
      const response = await postService.fetchPosts(
        20,
        nextCursor ?? undefined
      );
      setPosts((prevPosts: Post[]) => [...prevPosts, ...response.posts]);

      if (response.next_cursor) {
        setNextCursor(response.next_cursor);
      } else {
        setHasMore(false);
      }
    } catch (error) {
      console.error("Failed to load more posts:", error);
    } finally {
      setIsLoading(false);
    }
  }, [isLoading, hasMore, nextCursor]);

  return { posts, isLoading, hasMore, loadMorePosts };
}
