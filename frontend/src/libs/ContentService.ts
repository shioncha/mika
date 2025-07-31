import type { Post } from "../type/post";
import type { Tag } from "../type/tag";
import type { User } from "../type/user";
import apiClient from "./api";

// Posts関連のAPIサービス
export const postService = {
  /**
   * すべての投稿を取得
   */
  async fetchPosts(): Promise<Post[]> {
    // apiClientを使い、GETリクエストを送信
    // AuthorizationヘッダーなどはapiClientが自動で付与します
    const { data } = await apiClient.get<Post[]>("/posts");
    return data;
  },

  /**
   * 新しい投稿を作成
   * @param content 投稿内容
   */
  async createPost(content: string): Promise<Post> {
    const { data } = await apiClient.post<Post>("/posts", { content });
    return data;
  },

  /**
   * 指定したIDの投稿を1件取得
   * @param id 投稿のID
   */
  async fetchPostById(id: string): Promise<Post> {
    const { data } = await apiClient.get<Post>(`/posts/${id}`);
    return data;
  },

  /**
   * 指定したIDの投稿を更新
   * @param id 投稿のID
   * @param content 更新する内容
   */
  async updatePost(id: string, content: string): Promise<Post> {
    const { data } = await apiClient.put<Post>(`/posts/${id}`, { content });
    return data;
  },

  /**
   * 指定したIDの投稿を削除
   * @param id 投稿のID
   */
  async deletePost(id: string): Promise<void> {
    await apiClient.delete(`/posts/${id}`);
  },
};

// Tags関連のAPIサービス
export const tagService = {
  /**
   * タグの一覧を取得
   */
  async fetchTags(): Promise<Tag[]> {
    const { data } = await apiClient.get<Tag[]>("/tags");
    return data;
  },

  /**
   * 指定したタグを持つすべての投稿を取得
   * @param tagName タグ名
   */
  async fetchPostsByTag(tagName: string): Promise<Post[]> {
    const { data } = await apiClient.get<Post[]>(`/tags/${tagName}/posts`);
    return data;
  },
};

// Users関連のAPIサービス
export const userService = {
  /**
   * 現在ログインしているユーザーの情報を取得
   */
  async fetchMe(): Promise<User> {
    const { data } = await apiClient.get<User>("/users/me");
    return data;
  },
};
