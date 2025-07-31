import { useEffect, useState } from "react";
import { IoIosClose } from "react-icons/io";
import { useNavigate, useParams } from "react-router";

import { postService } from "../../../libs/ContentService";
import { formatDate, localDate, localTime } from "../../../libs/datetime";
import type { Post } from "../../../type/post";
import Button from "../Button";
import style from "./modal.module.css";

export default function Modal() {
  const navigate = useNavigate();
  const { id } = useParams();
  const [posts, setPosts] = useState<Post[]>([]);

  useEffect(() => {
    async function fetchPost() {
      if (!id) return;

      try {
        const post = await postService.fetchPostById(id);
        setPosts([post]);
      } catch {
        console.error("Failed to fetch post");
      }
    }

    fetchPost();
  }, [id]);

  function editPost() {
    const textarea = document.querySelector("textarea");
    const content = textarea?.value;
    if (!content) return;

    postService
      .updatePost(id || "", content)
      .then((updatedPost) => {
        setPosts([updatedPost]);
        navigate(-1);
      })
      .catch(() => {
        console.error("Failed to update post");
      });
    closeModal();
  }

  function deletePost() {
    postService
      .deletePost(id || "")
      .then(() => {
        setPosts((prevPosts) => prevPosts.filter((post) => post.ID !== id));
        navigate(-1);
      })
      .catch(() => {
        console.error("Failed to delete post");
      });
    closeModal();
  }

  function closeModal() {
    navigate(-1);
  }

  return (
    <div className={`${style.overlay} ${style.animation}`} onClick={closeModal}>
      <div className={style.modal} onClick={(e) => e.stopPropagation()}>
        <div className={style.header}>
          <Button variant="icon" onClick={closeModal}>
            <IoIosClose color="white" size="1.5rem" />
          </Button>
          <Button variant="primary" onClick={editPost}>
            編集
          </Button>
        </div>
        <textarea defaultValue={posts[0]?.Content} className={style.textarea} />
        {/* <p className={style.dueDate}>
          期限:{" "}
          {posts[0]?.due_date
            ? `${localDate(posts[0]?.due_date)} ${localTime(
                posts[0]?.due_date
              )}`
            : "未設定"}
        </p> */}
        <p>
          作成: {localDate(formatDate(posts[0]?.CreatedAt))}{" "}
          {localTime(formatDate(posts[0]?.CreatedAt))}
        </p>
        <p>
          最終更新: {localDate(formatDate(posts[0]?.UpdatedAt))}{" "}
          {localTime(formatDate(posts[0]?.UpdatedAt))}
        </p>
        <Button variant="primary" onClick={deletePost}>
          削除
        </Button>
      </div>
    </div>
  );
}
