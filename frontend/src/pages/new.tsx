import { useContext, useState } from "react";

import Button from "../components/elements/Button";
import { AuthContext } from "../hooks/auth_context";
import { PostsAPI } from "../libs/api";
import style from "../styles/pages/new.module.css";
import type { Post } from "../type/post";

function NewPage() {
  const { token } = useContext(AuthContext);
  const [postContent, setPostContent] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [, setPosts] = useState<Post[]>([]);

  function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    if (!postContent.trim()) {
      setError("Content cannot be empty.");
      return;
    }
    PostsAPI({
      method: "POST",
      body: { content: postContent },
      token,
      setPosts,
    });
  }

  return (
    <div>
      <h1>New Post</h1>
      {error && <p className={style.error}>{error}</p>}
      <form onSubmit={handleSubmit} className={style.form}>
        <input
          type="text"
          placeholder="Content"
          onChange={(e) => setPostContent(e.target.value)}
          value={postContent}
          className={style.input}
        />
        <Button type="submit">Create</Button>
      </form>
    </div>
  );
}

export default NewPage;
