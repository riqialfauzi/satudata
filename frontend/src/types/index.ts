// ============ Enums ============

export type ReleaseType = "dataset" | "article" | "infographic";
export type ReleaseStatus = "draft" | "published" | "archived";
export type StandardStatus = "active" | "archived" | "draft";
export type UserRole = "admin" | "editor" | "viewer";

// ============ API Response ============

export interface Meta {
  page: number;
  limit: number;
  total: number;
  total_pages: number;
}

export interface APIResponse<T> {
  success: boolean;
  message?: string;
  data?: T;
  meta?: Meta;
  errors?: string[];
}

// ============ Domain Models ============

export interface Release {
  id: string;
  title: string;
  slug: string;
  description?: string;
  release_type: ReleaseType;
  status: ReleaseStatus;
  year: number;
  cover_image_url?: string;
  tags?: string[];
  view_count: number;
  published_at?: string;
  created_by?: string;
  created_at: string;
  updated_at: string;
  dataset_metadata?: DatasetMetadata;
  article_metadata?: ArticleMetadata;
}

export interface DatasetMetadata {
  file_url: string;
  file_format: string;
  file_size: number;
  row_count?: number;
  column_count?: number;
  data_source?: string;
  data_period_start?: string;
  data_period_end?: string;
  update_frequency?: string;
  is_geospatial: boolean;
}

export interface ArticleMetadata {
  content: string;
  excerpt?: string;
  author_name?: string;
  reading_time_minutes?: number;
  is_featured: boolean;
  category?: string;
}

export interface Standard {
  id: string;
  title: string;
  description?: string;
  year: number;
  file_url?: string;
  file_size?: number;
  status: StandardStatus;
  version: string;
  is_current: boolean;
  created_at: string;
  updated_at: string;
}

export interface User {
  id: string;
  email: string;
  full_name: string;
  role: UserRole;
  is_active?: boolean;
  last_login_at?: string;
  created_at?: string;
}

export interface TokenResponse {
  access_token: string;
  refresh_token: string;
  token_type: string;
  expires_in: number;
  user: User;
}

export interface ReleaseStats {
  total: number;
  by_type: Record<string, number>;
  by_year: Record<string, number>;
}

// ============ Request Types ============

export interface LoginRequest {
  email: string;
  password: string;
}

export interface RegisterRequest {
  email: string;
  password: string;
  full_name: string;
}

export interface CreateReleaseRequest {
  title: string;
  description?: string;
  release_type: ReleaseType;
  year: number;
  cover_image_url?: string;
  tags?: string[];
  file_url?: string;
  file_format?: string;
  file_size?: number;
  data_source?: string;
  content?: string;
  excerpt?: string;
  author_name?: string;
  category?: string;
  is_featured?: boolean;
}

export interface UpdateReleaseRequest {
  title?: string;
  description?: string;
  status?: ReleaseStatus;
  year?: number;
  cover_image_url?: string;
  tags?: string[];
}

export interface CreateStandardRequest {
  title: string;
  description?: string;
  year: number;
  file_url?: string;
  file_size?: number;
  version?: string;
  is_current?: boolean;
}

export interface UpdateStandardRequest {
  title?: string;
  description?: string;
  status?: StandardStatus;
  file_url?: string;
  file_size?: number;
  version?: string;
  is_current?: boolean;
}

export interface ReleaseFilter {
  type?: string;
  year?: number;
  q?: string;
  page?: number;
  limit?: number;
  sort_by?: string;
  sort_dir?: string;
}
