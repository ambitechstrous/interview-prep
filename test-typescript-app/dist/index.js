"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const api_1 = require("./api");
async function main() {
    const { data: posts } = await api_1.api.get("/posts");
    console.log(`Fetched ${posts.length} posts`);
    console.log("First post:", posts[0]);
}
main().catch((err) => {
    console.error("Error:", err);
    process.exit(1);
});
