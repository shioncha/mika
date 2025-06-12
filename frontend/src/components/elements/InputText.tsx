import { type InputHTMLAttributes,useState } from "react";
import { FiEye, FiEyeOff } from "react-icons/fi";

import styles from "../../styles/components/elements/InputText.module.css";

interface InputTextProps extends InputHTMLAttributes<HTMLInputElement> {
  hasPasswordMask?: boolean;
}

function InputText({
  hasPasswordMask = false,
  className,
  type,
  ...props
}: InputTextProps) {
  const [unmasking, setUnmasking] = useState(false);

  const inputText = (
    <input
      type={unmasking ? "text" : type}
      className={`${styles.input} ${className}`}
      {...props}
    />
  );

  if (hasPasswordMask) {
    return (
      <div className={styles.maskInput}>
        {inputText}
        <button
          className={styles.maskButton}
          type="button"
          onClick={() => setUnmasking((s) => !s)}
        >
          {unmasking ? <FiEye /> : <FiEyeOff />}
        </button>
      </div>
    );
  }

  return inputText;
}
export default InputText;
