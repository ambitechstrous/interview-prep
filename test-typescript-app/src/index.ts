import { api } from "./api";
import { Post } from "./types";

async function main() {
  const { data: posts } = await api.get<Post[]>("/posts");
  console.log(`Fetched ${posts.length} posts`);
  console.log("First post:", posts[0]);
}

main().catch((err) => {
  console.error("Error:", err);
  process.exit(1);
});
