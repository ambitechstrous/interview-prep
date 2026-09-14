"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.api = void 0;
const BASE_URL = "https://localhost:5000/";
async function fetchJson(path) {
    const response = await fetch(`${BASE_URL}${path}`);
    if (!response.ok) {
        throw new Error(`HTTP ${response.status}: ${response.statusText}`);
    }
    const data = (await response.json());
    return { data, status: response.status, ok: response.ok };
}
exports.api = {
    get: (path) => fetchJson(path),
};
