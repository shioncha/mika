import type { Dispatch, SetStateAction } from "react";

import type { Post } from "../type/post";
import type { Tag } from "../type/tag";

interface PostAPIProps {
  method: string;
  body?: object;
  id?: string;
  token: string | null;
  setPosts?: Dispatch<SetStateAction<Post[]>>;
  tag?: string;
}

interface TagsAPIProps {
  method: string;
  token: string | null;
  setTags: Dispatch<SetStateAction<Tag[]>>;
}

async function PostsAPI({
  method,
  body,
  token,
  setPosts,
}: PostAPIProps): Promise<Post[]> {
  const myHeaders = new Headers();
  myHeaders.append("Authorization", `Bearer ${token}`);
  const myBody = JSON.stringify(body);
  const requestOptions: RequestInit = {
    method,
    headers: myHeaders,
  };
  if (method != "GET") requestOptions.body = myBody;

  fetch("/api/posts", requestOptions)
    .then((response) => response.json())
    .then((result) => (setPosts != undefined ? setPosts(result) : []))
    .catch((error) => console.error(error));
  return [];
}

async function PostAPI({
  method,
  id,
  body,
  token,
  setPosts,
}: PostAPIProps) {
  const myHeaders = new Headers();
  myHeaders.append("Authorization", `Bearer ${token}`);
  const myBody = JSON.stringify(body);
  const requestOptions: RequestInit = {
    method,
    headers: myHeaders,
  };
  if (method != "GET") requestOptions.body = myBody;

  fetch(`/api/posts/${id}`, requestOptions)
    .then((response) => response.json())
    .then((result) => (setPosts != undefined ? setPosts([result]) : null))
    .catch((error) => console.error(error));
}

async function TagsAPI({
  method,
  token,
  setTags,
}: TagsAPIProps): Promise<Tag[]> {
  const headers = new Headers();
  headers.append("Authorization", `Bearer ${token}`);
  const requestOptions: RequestInit = {
    method,
    headers,
  };

  fetch("/api/tags", requestOptions)
    .then((response) => response.json())
    .then((result) => setTags(result))
    .catch((error) => console.error(error));
  return [];
}

async function TagsPostsAPI({
  method,
  tag,
  body,
  token,
  setPosts,
}: PostAPIProps): Promise<Post[]> {
  const myHeaders = new Headers();
  myHeaders.append("Authorization", `Bearer ${token}`);
  const myBody = JSON.stringify(body);
  const requestOptions: RequestInit = {
    method,
    headers: myHeaders,
  };
  if (method != "GET") requestOptions.body = myBody;

  fetch(`/api/tags/${tag}/posts`, requestOptions)
    .then((response) => response.json())
    .then((result) => (setPosts != undefined ? setPosts(result) : []))
    .catch((error) => console.error(error));
  return [];
}

export { PostAPI, PostsAPI, TagsAPI, TagsPostsAPI };
