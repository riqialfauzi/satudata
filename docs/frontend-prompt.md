# 🎯 Prompt — Frontend Satudata

## Konteks

Bangun **frontend untuk platform Satudata**, sebuah portal data statistik/publik. Backend sudah tersedia (Go + Gin, REST API `/api/v1`) dengan domain utama: **Releases** (dataset & article), **Standards**, **Auth (JWT)**, dan **Admin** (users, audit logs). Frontend terdiri dari dua bagian: **portal publik** (browse & baca data) dan **dashboard admin** (CRUD, upload, moderasi).

## Tech Stack

- **Next.js 14+** (App Router) + **TypeScript**
- **TailwindCSS** + **shadcn/ui** untuk komponen
- **TanStack Query** (React Query) untuk data fetching & caching
- **Zustand** untuk state global ringan (auth session)
- **React Hook Form** + **Zod** untuk form & validasi
- **Axios** (instance dengan interceptor untuk JWT + refresh token)
- **next-themes** (dark/light mode), **lucide-react** (ikon)

## Sumber Kebenaran: Kontrak API Backend

Semua response dibungkus `APIResponse` (`{ success, message, data, meta }`). Pagination lewat `meta` (`page, limit, total, total_pages`).

**Public**
- `GET /api/v1/public/releases?type=&year=&page=&limit=&q=` — list releases (dataset/article), berpaginasi
- `GET /api/v1/public/releases/:id` — detail release
- `GET /api/v1/public/releases/slug/:slug` — detail via slug (untuk URL SEO-friendly)
- `GET /api/v1/public/releases/stats` — statistik (total by type, by year)
- `GET /api/v1/public/standards` — list standar data
- `GET /health` — status service

**Auth**
- `POST /api/v1/auth/login` → `{ access_token, refresh_token, user }`
- `POST /api/v1/auth/register`
- `POST /api/v1/auth/refresh`
- `POST /api/v1/auth/logout`

**Protected** (butuh `Authorization: Bearer <token>`)
- `GET /api/v1/protected/profile`
- `POST /api/v1/protected/releases` — buat release (+ upload file dataset / cover image)
- `PUT /api/v1/protected/releases/:id`
- `DELETE /api/v1/protected/releases/:id` (Admin only)
- `POST /api/v1/protected/standards` · `PUT /api/v1/protected/standards/:id`

**Admin** (role = admin)
- `GET /api/v1/admin/users` · `PUT /api/v1/admin/users/:id/role`
- `GET /api/v1/admin/audit-logs`

## Model Data Inti (mirror dari backend)

```ts
type ReleaseType = "dataset" | "article";

interface Release {
  id: string;
  slug: string;
  title: string;
  type: ReleaseType;
  description: string;
  year: number;
  category?: string;
  file_url?: string;      // dataset CSV/JSON
  cover_image?: string;   // article
  published_at: string;
  created_at: string;
  updated_at: string;
}

interface Standard {
  id: string;
  year: number;
  title: string;
  doc_url: string;        // PDF standar data
  is_active: boolean;
}

interface User {
  id: string;
  name: string;
  email: string;
  role: "admin" | "editor" | "viewer";
}
```

## Yang Harus Dibangun

### A. Portal Publik
1. **Landing page** — hero, statistik ringkas (dari `/releases/stats`), release terbaru, CTA cari data.
2. **Halaman Releases** — grid/list dengan filter (type, year, kategori), pencarian, pagination, sort. State filter tersinkron ke URL query params.
3. **Detail Release** — dataset: preview metadata + tombol download; article: render konten + cover. Pakai slug untuk URL.
4. **Halaman Standards** — daftar standar per tahun, tandai yang aktif, tombol lihat/unduh PDF.
5. **Halaman Status/Health** — indikator sederhana dari `/health`.

### B. Dashboard Admin
6. **Login** — form email/password, simpan token, redirect ke dashboard.
7. **Layout dashboard** — sidebar, topbar (profil, logout), route guard berbasis role.
8. **Manajemen Releases** — tabel dengan search/filter/pagination, form create/edit (React Hook Form + Zod), **upload file** (dataset & cover) dengan progress, soft-delete (admin only) dengan konfirmasi.
9. **Manajemen Standards** — CRUD + upload PDF + toggle aktif.
10. **Manajemen Users** (admin) — daftar user, ubah role.
11. **Audit Logs** (admin) — tabel read-only, filter tanggal/action/user.
12. **Profil** — lihat info user.

## Requirement Non-Fungsional

- **Auth flow lengkap:** interceptor Axios menyisipkan bearer token; auto-refresh saat 401; redirect ke login saat refresh gagal.
- **Proteksi rute** berbasis role (public / authenticated / admin).
- **Loading & error states** konsisten (skeleton, empty state, toast error). Tangani bentuk `APIResponse` seragam.
- **Responsive** (mobile-first) + **dark mode**.
- **Aksesibilitas** dasar (label, focus, kontras) dan **SEO** (metadata dinamis, slug URL, Open Graph) untuk halaman publik.
- **Env config** via `.env.local` (`NEXT_PUBLIC_API_BASE_URL`).
- **Type-safe:** definisikan tipe API di `src/types`, layer service di `src/lib/api`.

## Struktur Folder yang Disarankan

```
src/
  app/                      # routes (App Router)
    (public)/               # landing, releases, standards
    (dashboard)/admin/      # protected admin area
    login/
  components/               # ui/ (shadcn), shared/
  lib/api/                  # axios instance + service per domain
  hooks/                    # useReleases, useAuth, dll (React Query)
  store/                    # zustand (auth)
  types/                    # kontrak API
  lib/                      # utils, validators (zod), constants
```

## Definition of Done

- Semua halaman & endpoint di atas terhubung ke API nyata (bukan mock).
- Auth + refresh + role guard berjalan.
- Filter/pagination tersinkron URL, cache lewat React Query.
- Responsive, dark mode, loading/error/empty states lengkap.
- Lint & type-check bersih; README setup (env, run, build).
