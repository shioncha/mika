import styles from "./List.module.css";

function List({
  children,
  className,
  ...props
}: React.HTMLAttributes<HTMLDivElement>) {
  return (
    <div className={`${styles.list} ${className}`} {...props}>
      {children}
    </div>
  );
}

export default List;
