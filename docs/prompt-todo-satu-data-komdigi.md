# Prompt & TODO — Klon Portal Satu Data KOMDIGI (Go + Next.js)

> Referensi UI: https://data.komdigi.go.id/ ("Satu Data KOMDIGI") — portal open data bergaya CKAN
> Stack target: **Backend Golang** + **Frontend Next.js**
> Theme color portal asli: `#006792`

---

## Bagian 1 — PROMPT (untuk diberikan ke AI coding assistant / dokumentasi tim)

Salin blok di bawah ini sebagai prompt utama. Sudah dirancang lengkap agar bisa langsung dieksekusi bertahap.

```
Kamu adalah senior full-stack engineer. Bangun ulang portal open data "Satu Data KOMDIGI"
(referensi: data.komdigi.go.id) dengan backend Golang dan frontend Next.js.

KONTEKS PRODUK
Portal open data pemerintah yang menyajikan dataset dari berbagai unit kerja, dengan
fitur pencarian, filter, kategori, organisasi, visualisasi, infografis, dan unduh data.
Bergaya CKAN (dataset → resources). Warna utama #006792, tampilan bersih & aksesibel,
mobile-first (PWA-ready), Bahasa Indonesia.

HALAMAN YANG HARUS ADA (public)
1. Beranda (/) — hero + search bar, statistik ringkas (jumlah dataset/organisasi/kategori),
   daftar dataset terbaru/populer, kategori unggulan, organisasi, CTA.
2. Daftar Dataset (/dataset) — pencarian + faceted filter (kategori/grup, organisasi,
   format file, tag), sorting (relevansi/terbaru/A-Z), paginasi, jumlah hasil.
3. Detail Dataset (/dataset/[slug]) — judul, deskripsi, metadata (organisasi, lisensi,
   frekuensi update, tanggal), daftar resources (CSV/XLSX/JSON/PDF) dengan tombol unduh &
   preview, tag, dataset terkait.
4. Daftar Organisasi (/organisasi) — grid kartu organisasi + jumlah dataset per organisasi.
5. Detail Organisasi (/organisasi/[slug]) — profil + dataset milik organisasi tsb.
6. Daftar Kategori/Grup (/kategori) dan detail kategori (/kategori/[slug]).
7. Infografis/Artikel (/infografis) — grid kartu + detail (/infografis/[slug]).
8. Statistik/Visualisasi (/statistik) — chart agregat (opsional bila waktu cukup).
9. Tentang (/tentang), FAQ (/faq).
10. Login (/login) — untuk data non-publik / admin (auth JWT).

KOMPONEN UI
Navbar (logo, menu, search, tombol login), footer (kontak/link kementerian),
SearchBar dengan autocomplete, FacetFilter (checkbox + count), DatasetCard, ResourceItem
(ikon per format + tombol download/preview), Pagination, StatCard, Breadcrumb,
Badge/Tag, EmptyState, Skeleton loading, Toast.

BACKEND (Golang)
- Framework HTTP: Gin atau Fiber. Struktur clean/layered (handler → service → repository).
- Database: PostgreSQL (+ GORM atau sqlc). Full-text search via Postgres tsvector atau
  integrasi Meilisearch/Elasticsearch untuk pencarian dataset.
- Auth: JWT (access + refresh), role-based (public, editor, admin).
- Penyimpanan file resource: local disk atau S3-compatible (MinIO).
- Sediakan REST API JSON (lihat kontrak endpoint di bawah). Dokumentasi via Swagger/OpenAPI.
- Middleware: CORS, rate limit, logging (zap), recovery, request-id.
- Config via env (.env / viper). Migration via golang-migrate atau goose.

FRONTEND (Next.js)
- Next.js App Router + TypeScript + Tailwind CSS. Warna brand #006792.
- Data fetching: Server Components + fetch ke API Go; TanStack Query untuk interaktif
  (filter/paginasi client). SEO metadata per halaman, sitemap, PWA manifest.
- State filter tersimpan di URL query params (shareable).
- Aksesibilitas WCAG AA, responsive, dark-mode opsional.

KONTRAK API (REST, prefix /api/v1)
GET  /datasets?q=&group=&org=&format=&tag=&sort=&page=&limit=   -> list + pagination meta
GET  /datasets/{slug}                                           -> detail + resources
GET  /organizations                                             -> list + dataset_count
GET  /organizations/{slug}                                      -> detail + datasets
GET  /groups                                                    -> list kategori/grup
GET  /groups/{slug}                                             -> detail + datasets
GET  /articles (infografis)          /articles/{slug}
GET  /stats                                                     -> agregat untuk beranda
GET  /search/suggest?q=                                         -> autocomplete
POST /auth/login   POST /auth/refresh   POST /auth/logout
(admin, protected) CRUD /datasets, /resources, /organizations, /groups, /articles

MODEL DATA (inti)
Dataset(id, slug, title, notes, organization_id, license, update_frequency, tags[],
  metadata_created, metadata_modified, is_public, views)
Resource(id, dataset_id, name, format, url/file_path, size, description)
Organization(id, slug, name, description, image_url)
Group(id, slug, name, description, image_url)   // kategori
Article(id, slug, title, body, image_url, published_at)   // infografis
User(id, email, password_hash, role)
Tag(id, name, slug)

ATURAN KUALITAS
- Tulis test (unit service + handler) minimal untuk endpoint dataset & auth.
- Seed data contoh agar UI bisa dilihat tanpa data asli.
- Docker Compose: postgres + backend + frontend (+ minio/meilisearch bila dipakai).
- Kerjakan bertahap sesuai TODO; setelah tiap tahap, jalankan & verifikasi.
```

---

## Bagian 2 — TODO / CHECKLIST IMPLEMENTASI

### Fase 0 — Setup & Fondasi
- [x] Buat repo monorepo/dua folder: `/backend` (Go) dan `/frontend` (Next.js).
- [x] Siapkan `docker-compose.yml`: postgres, backend, frontend (+ minio/meilisearch opsional).
- [x] Tentukan variabel env (`.env.example`) untuk DB, JWT secret, storage, CORS origin.
- [x] Setup linter/formatter: `golangci-lint` (Go), `eslint` + `prettier` (Next.js).

### Fase 1 — Backend Go: Skeleton
- [x] Init module Go, pilih framework (Gin/Fiber), struktur folder `cmd/`, `internal/handler`, `internal/service`, `internal/repository`, `internal/model`, `pkg/`.
- [x] Koneksi PostgreSQL (GORM/sqlc) + health-check endpoint `/healthz`.
- [x] Setup migration (golang-migrate/goose) untuk semua tabel model data.
- [x] Middleware: CORS, logger (zap), recovery, request-id, rate limit.
- [x] Config loader (viper) dari env.

### Fase 2 — Backend Go: Domain & API Publik
- [x] Model + repository: Release (dataset/article/infographic), Standard, User ✅. Organization, Group ✅ (baru).
- [x] `GET /api/v1/datasets` ✅ (via `/api/v1/public/releases` dengan filter, pagination, full-text search).
- [x] `GET /api/v1/datasets/{slug}` ✅ + dataset terkait ✅ (via `/api/v1/public/releases/slug/:slug` + `/related`).
- [x] `GET /api/v1/organizations` & `/{slug}` ✅ (baru).
- [x] `GET /api/v1/groups` & `/{slug}` ✅ (baru).
- [x] `GET /api/v1/articles` & `/{slug}` ✅ (covered by releases).
- [x] `GET /api/v1/stats` ✅ (via `/api/v1/public/releases/stats`).
- [x] `GET /api/v1/search/suggest?q=` ✅ (baru).
- [x] Pencarian: Postgres full-text (tsvector + GIN index) ✅ (baru).

### Fase 3 — Backend Go: Auth & Admin
- [x] `POST /auth/login`, `/refresh`, `/logout` (JWT access+refresh).
- [x] Middleware role-based (public/editor/admin).
- [x] CRUD terproteksi: datasets ✅, resources (upload file via MinIO) ✅ (baru), organizations ✅ (baru), groups ✅ (baru), articles ✅.
- [x] Validasi input + error handling konsisten (format error JSON standar).

### Fase 4 — Backend Go: Kualitas
- [x] Unit test service ✅ (auth, release, standard). Handler tests ⚠️ (masih pending).
- [x] Seed data contoh ✅ (2 users, 3 standards, 4 releases). + Organisasi ✅ (baru), Grup ✅ (baru).
- [x] Dokumentasi OpenAPI/Swagger di `/swagger`.

### Fase 5 — Frontend Next.js: Fondasi
- [x] Init Next.js (App Router) + TypeScript + Tailwind; set warna brand `#006792`.
- [x] Layout global: Navbar, Footer, container, tema/typography.
- [x] API client (fetch wrapper) + tipe TypeScript dari kontrak API; setup TanStack Query.
- [x] Komponen dasar: Button, Card, Badge, Pagination, Skeleton, Toast ✅. Breadcrumb ✅ (baru), SearchBar ✅ (baru), EmptyState ✅ (baru).
- [x] PWA manifest ✅ + metadata SEO ✅ (baru).

### Fase 6 — Frontend Next.js: Halaman Publik
- [x] Beranda: hero + CTA, StatCard (dari `/stats`), dataset terbaru ✅, organisasi ✅, kategori ✅ (baru).
- [x] Daftar Dataset: SearchBar ✅ + FacetFilter ✅ (organisasi/kategori), pagination ✅, URL query params ✅.
- [x] Detail Dataset: metadata, resource (format + download), tag ✅, dataset terkait ✅ (baru).
- [x] Daftar Organisasi ✅ (terintegrasi di beranda & filter).
- [x] Daftar Kategori/Grup ✅ (terintegrasi di beranda & filter).
- [x] Infografis: grid + detail (via releases).
- [x] Tentang ✅, FAQ ✅.
- [x] Halaman Login ✅ (form + token, guard route admin).

### Fase 7 — Integrasi & Polish
- [x] Sambungkan semua halaman ke API Go ✅; loading/error/empty state ⚠️ (sebagian).
- [x] Responsif (mobile-first) ✅; aksesibilitas ⚠️ (masih perlu polish).
- [ ] Preview resource (CSV/JSON) sederhana ⏳ (masih pending).
- [ ] Statistik/visualisasi chart ⏳ (masih pending).
- [x] SEO: sitemap.xml ✅, robots.txt ✅, metadata per halaman ✅, Open Graph ✅ (baru).

### Fase 8 — Verifikasi & Deploy
- [x] Docker Compose end-to-end ✅ (setup sudah ada).
- [ ] Tes lintas peramban + Lighthouse ⏳ (masih pending).
- [x] Pipeline CI (lint + test + build) — GitHub Actions ✅ (baru).
- [x] Dokumentasi README ✅ (cara run lokal, env, arsitektur).

---

## Catatan & Asumsi
- Portal asli berbasis **CKAN** (pola `dataset` → `resource`, filter `groups`/`org`). Struktur data di atas mengikuti pola tersebut agar mudah migrasi/kompatibel.
- Halaman utama portal dirender di sisi klien, jadi rute di prompt/TODO adalah rekonstruksi berdasarkan peta navigasi publik (`/opendata/dataset`, `/opendata/organization`, `/opendata/about`, `/article`, `/faq`, `/login`). Sesuaikan slug rute sesuai preferensimu.
- Alternatif praktis: karena backend asli CKAN, kamu bisa memilih membungkus **CKAN API** langsung dan hanya membangun frontend Next.js. TODO di atas mengasumsikan kamu ingin backend Go sendiri.

Sumber referensi UI: [Beranda](https://data.komdigi.go.id/opendata), [Dataset](https://data.komdigi.go.id/opendata/dataset), [Organisasi](https://data.komdigi.go.id/opendata/organization), [Tentang](https://data.komdigi.go.id/opendata/about), [Infografis](https://data.komdigi.go.id/article), [FAQ](https://data.komdigi.go.id/faq), [Login](https://data.komdigi.go.id/login)
