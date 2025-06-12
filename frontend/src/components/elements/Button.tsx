import type { ButtonHTMLAttributes } from "react";

import styles from "../../styles/components/elements/Button.module.css";

interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: "primary" | "secondary" | "icon";
}

function Button({
  variant = "primary",
  children,
  className,
  ...props
}: ButtonProps) {
  const variantStyles = {
    primary: styles.primary,
    secondary: styles.secondary,
    icon: styles.icon,
  };

  return (
    <button
      className={`${styles.button} ${variantStyles[variant]} ${className}`}
      {...props}
    >
      {children}
    </button>
  );
}

export default Button;
