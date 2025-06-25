import type { Post } from "../type/post";

interface GroupedPostsProps {
  [key: string]: Post[];
}

function formatDate(dateString: string): Date {
  if (!dateString) {
    return new Date();
  }
  try {
    const normalized = dateString.replace(
      /(\d{4}-\d{2}-\d{2}) (\d{2}:\d{2}:\d{2})\.(\d{3})\d{3} \+0000 UTC/,
      "$1T$2.$3Z"
    );
    return new Date(normalized);
  } catch {
    return new Date();
  }
}

function localDate(date: Date): string {
  return date.toLocaleString("ja-JP", {
    year: "numeric",
    month: "2-digit",
    day: "2-digit",
  });
}

function localTime(date: Date): string {
  return date.toLocaleString("ja-JP", {
    hour: "numeric",
    minute: "numeric",
    hour12: false,
  });
}

function groupByDate(data: Post[]): GroupedPostsProps {
  if (!data || !Array.isArray(data)) {
    return {};
  }

  return data.reduce((grouped: GroupedPostsProps, item) => {
    const date = localDate(formatDate(item.CreatedAt));
    grouped[date] = grouped[date] || [];
    grouped[date].push(item);
    return grouped;
  }, {});
}

export { formatDate, groupByDate, localDate, localTime };
