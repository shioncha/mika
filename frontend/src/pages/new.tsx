import { useState } from "react";
import { useNavigate } from "react-router";

import Button from "../components/elements/Button";
import { postService } from "../libs/ContentService";
import style from "../styles/pages/new.module.css";

function NewPage() {
  const [postContent, setPostContent] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [isSubmitting, setIsSubmitting] = useState(false);
  const navigate = useNavigate();

  async function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    setIsSubmitting(true);
    setError(null);
    if (!postContent.trim()) {
      setError("Content cannot be empty.");
      setIsSubmitting(false);
      return;
    }
    try {
      await postService.createPost(postContent);
      navigate("/");
    } catch {
      console.error("Failed to create post");
      setError("Failed to create post. Please try again.");
    } finally {
      setPostContent("");
      setIsSubmitting(false);
    }
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
        <Button type="submit" disabled={isSubmitting}>
          {isSubmitting ? "Creating..." : "Create"}
        </Button>
      </form>
    </div>
  );
}

export default NewPage;
