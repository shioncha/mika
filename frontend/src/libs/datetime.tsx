import type { Post } from "../type/post";

interface GroupedPostsProps {
  [key: string]: Post[];
}

function localDate(dateString: string): string {
  const date = new Date(dateString);
  return date.toLocaleString("ja-JP", {
    year: "numeric",
    month: "2-digit",
    day: "2-digit",
  });
}

function localTime(dateString: string): string {
  const date = new Date(dateString);
  return date.toLocaleString("ja-JP", {
    hour: "numeric",
    minute: "numeric",
    hour12: false,
  });
}

function groupByDate(data: Post[]): GroupedPostsProps {
  const grouped: GroupedPostsProps = {};

  if (!data || data.length === 0) {
    return grouped;
  }

  data.forEach((item) => {
    const date = localDate(item.CreatedAt);
    if (!grouped[date]) {
      grouped[date] = [];
    }
    grouped[date].push(item);
  });

  return grouped;
}

export { groupByDate, localDate, localTime };
