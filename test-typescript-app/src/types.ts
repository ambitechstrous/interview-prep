export interface ApiResponse<T> {
  data: T;
  status: number;
  ok: boolean;
}

// Replace with your actual data shape
export interface Post {
  id: number;
  title: string;
  body: string;
  userId: number;
}
