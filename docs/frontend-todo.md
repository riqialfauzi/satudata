# 📋 Satudata Frontend — Todo

**Stack:** Next.js 14 (App Router) · TypeScript · TailwindCSS · shadcn/ui · TanStack Query · Zustand · React Hook Form + Zod · Axios

---

## ✅ PHASE 1: Project Setup & Foundation ✅ **COMPLETED**

### 1.1 Bootstrap
- [x] Init `create-next-app` (TypeScript, App Router, Tailwind, ESLint)
- [x] Setup `shadcn/ui` + tema (radius, warna brand)
- [x] Install deps: `@tanstack/react-query`, `zustand`, `axios`, `react-hook-form`, `zod`, `next-themes`, `lucide-react`
- [x] Setup Prettier + ESLint + `tsconfig` path alias (`@/*`)
- [x] `.env.local` + `.env.example` (`NEXT_PUBLIC_API_BASE_URL`)

### 1.2 Struktur & Providers
- [x] Struktur folder (`app`, `components`, `lib/api`, `hooks`, `store`, `types`)
- [x] Root provider: React Query `QueryClientProvider` + `ThemeProvider`
- [x] Global layout, font, dark/light toggle
- [x] Komponen dasar: `Button`, `Input`, `Card`, `Badge`, `Skeleton`, `Dialog`, `Table`, `Pagination`, `Sonner`

### 1.3 Tipe & API Layer
- [x] Definisikan tipe di `src/types` (`Release`, `Standard`, `User`, `APIResponse`, `Meta`, `ReleaseType`)
- [x] Axios instance + `baseURL` + interceptor (inject bearer, handle `APIResponse`)
- [x] Helper unwrap `APIResponse` + normalisasi error

---

## ✅ PHASE 2: Auth & Session ✅ **COMPLETED**

- [x] Zustand store `useAuthStore` (user, tokens, actions, hydrate)
- [x] Service auth: `login`, `register`, `refresh`, `logout`, `getProfile`
- [x] Interceptor: auto-refresh saat 401, retry request, redirect login saat gagal
- [x] Persist token (localStorage)
- [x] Halaman `/login` (form RHF + Zod, error state, show/hide password)
- [x] Route guard: `PublicRoute`, `ProtectedRoute`, `AdminRoute` (cek role)
- [x] Hook `useAuth` (status, isAdmin, login, register, logout)

---

## ✅ PHASE 3: Portal Publik ✅ **COMPLETED**

### 3.1 Landing Page
- [x] Hero + CTA cari data
- [x] Statistik ringkas dari `GET /public/releases/stats` (total by type/year)
- [x] Section "Release Terbaru" (fetch list, limit kecil)
- [x] Footer + navigasi (Header, Footer, ThemeToggle)

### 3.2 Halaman Releases
- [x] Hook `useReleases(filters)` → `GET /public/releases`
- [x] Grid/list card release (dataset & article beda tampilan)
- [x] Filter: `type`, `year` + search `q`
- [x] Sinkron filter ↔ URL query params
- [x] Pagination dari `meta`
- [x] Loading skeleton + empty state + error state

### 3.3 Detail Release
- [x] Route `/releases/[slug]` → `GET /public/releases/slug/:slug`
- [x] Layout dataset: metadata + tombol **Download** (`file_url`)
- [x] Layout article: konten HTML
- [ ] Metadata dinamis (title, description, Open Graph) untuk SEO _(pending)_

### 3.4 Standards & Status
- [x] Halaman `/standards` → `GET /public/standards`, tandai aktif, lihat/unduh PDF
- [ ] Indikator status dari `GET /health` _(pending)_

---

## ✅ PHASE 4: Dashboard Admin — Layout & Releases ✅ **COMPLETED**

### 4.1 Layout Dashboard
- [x] Route group `(dashboard)` dengan guard auth
- [x] Sidebar navigasi + topbar (profil, logout, theme)
- [x] Page container reusable

### 4.2 Manajemen Releases
- [x] Tabel releases (search, pagination)
- [x] Form Create/Edit (RHF + Zod): title, type, year, description, tags
- [ ] **Upload file** dataset & cover image (progress bar, validasi tipe/ukuran) _(pending)_
- [ ] Auto slug preview dari title _(pending)_
- [x] `POST /protected/releases`, `PUT /protected/releases/:id`
- [x] Soft-delete (admin only) + dialog konfirmasi
- [x] Invalidate query setelah mutasi

---

## ✅ PHASE 5: Dashboard Admin — Standards, Users, Audit ✅ **COMPLETED**

### 5.1 Standards
- [x] List standards + create (dialog form)
- [ ] Upload PDF + toggle `is_active` _(pending)_

### 5.2 Users (Admin)
- [x] `GET /admin/users` → tabel user
- [x] Ubah role `PUT /admin/users/:id/role` (dropdown)

### 5.3 Audit Logs (Admin)
- [x] `GET /admin/audit-logs` → tabel read-only
- [ ] Filter tanggal / action / user + pagination _(pending)_

### 5.4 Profil
- [x] Halaman profil dari `GET /protected/profile`

---

## ✅ PHASE 6: Polish, QA & Deploy ✅ **COMPLETED**

- [x] Konsistensi loading/error/empty state di semua halaman
- [x] Responsive (mobile hamburger menu + responsive grid)
- [x] Aksesibilitas dasar (label, aria-label pada icon buttons)
- [ ] SEO halaman publik (metadata, sitemap, OG image) _(pending)_
- [ ] Error boundary + halaman 404/500 _(pending)_
- [x] Lint & `tsc --noEmit` bersih
- [ ] README: setup env, run, build, struktur _(pending)_
- [ ] Build & deploy (Vercel / Docker), set env production _(pending)_

---

## 📊 Priority Matrix

| Priority | Fitur |
|----------|-------|
| **P0 (Must Have)** | Setup, API layer, Auth+guard, List & Detail Releases, Landing |
| **P1 (High)** | Dashboard layout, CRUD Releases + upload, Standards publik |
| **P2 (Medium)** | Standards CRUD, Users role mgmt, Audit logs, Profil |
| **P3 (Nice to Have)** | SEO/OG lanjutan, animasi, PWA, i18n |

## 🎯 Success Criteria

- [x] Semua halaman terhubung ke API nyata (bukan mock) — 14 routes
- [x] Auth + refresh token + role guard berjalan mulus
- [x] Filter & pagination tersinkron URL, cache via React Query
- [x] Responsive + dark mode + loading/error/empty states lengkap
- [x] Lint & type-check lolos (`tsc --noEmit` + `next lint` ✅)
- [ ] README: setup env, run, build, struktur _(pending)

## 📝 Notes

- Selaraskan tipe FE dengan struct backend; sesuaikan bila field API berbeda.
- Bungkus semua fetch dengan React Query (staleTime sesuai TTL cache backend).
- Simpan token secara aman; hindari expose di client bila memungkinkan.
- Commit tiap selesai satu phase.
