import styles from "../../styles/components/elements/List.module.css";

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
