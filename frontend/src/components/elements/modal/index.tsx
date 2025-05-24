import { useContext, useEffect, useState } from "react";
import { IoIosClose } from "react-icons/io";
import { useNavigate, useParams } from "react-router";

import { AuthContext } from "../../..//hooks/auth_context";
import { PostAPI } from "../../../libs/api";
import { localDate, localTime } from "../../../libs/datetime";
import type { Post } from "../../../type/post";
import style from "./modal.module.css";

export default function Modal() {
  const navigate = useNavigate();
  const { id } = useParams();
  const { token } = useContext(AuthContext);
  const [posts, setPosts] = useState<Post[]>([]);

  useEffect(() => {
    if (token == null) {
      return;
    }
    async function fetchPost() {
      await PostAPI({
        method: "GET",
        id,
        token,
        setPosts,
      });
    }
    fetchPost();
  }, [id, token]);

  function editPost() {
    const textarea = document.querySelector("textarea");
    const content = textarea?.value;
    if (!content) return;

    PostAPI({
      method: "PUT",
      id: id,
      body: { content, has_checkbox: false, is_checked: false, due_date: "" },
      token,
    });
    closeModal();
  }

  function deletePost() {
    PostAPI({
      method: "DELETE",
      id: id,
      token,
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
          <IoIosClose size={24} className={style.close} onClick={closeModal} />
          <button onClick={editPost} className={style.edit}>
            編集
          </button>
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
          作成: {localDate(posts[0]?.CreatedAt)}{" "}
          {localTime(posts[0]?.CreatedAt)}
        </p>
        <p>
          最終更新: {localDate(posts[0]?.UpdatedAt)}{" "}
          {localTime(posts[0]?.UpdatedAt)}
        </p>
        <button onClick={deletePost} className={style.button}>
          削除
        </button>
      </div>
    </div>
  );
}
