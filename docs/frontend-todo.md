# 📋 Satudata Frontend — Todo

**Stack:** Next.js 14 (App Router) · TypeScript · TailwindCSS · shadcn/ui · TanStack Query · Zustand · React Hook Form + Zod · Axios

---

## ✅ PHASE 1: Project Setup & Foundation _(Day 1-2)_

### 1.1 Bootstrap
- [ ] Init `create-next-app` (TypeScript, App Router, Tailwind, ESLint)
- [ ] Setup `shadcn/ui` + tema (radius, warna brand)
- [ ] Install deps: `@tanstack/react-query`, `zustand`, `axios`, `react-hook-form`, `zod`, `next-themes`, `lucide-react`
- [ ] Setup Prettier + ESLint + `tsconfig` path alias (`@/*`)
- [ ] `.env.local` + `.env.example` (`NEXT_PUBLIC_API_BASE_URL`)

### 1.2 Struktur & Providers
- [ ] Struktur folder (`app`, `components`, `lib/api`, `hooks`, `store`, `types`)
- [ ] Root provider: React Query `QueryClientProvider` + `ThemeProvider`
- [ ] Global layout, font, dark/light toggle
- [ ] Komponen dasar: `Button`, `Input`, `Card`, `Badge`, `Skeleton`, `Toast`, `Dialog`, `Table`, `Pagination`

### 1.3 Tipe & API Layer
- [ ] Definisikan tipe di `src/types` (`Release`, `Standard`, `User`, `APIResponse`, `Meta`, `ReleaseType`)
- [ ] Axios instance + `baseURL` + interceptor (inject bearer, handle `APIResponse`)
- [ ] Helper unwrap `APIResponse` + normalisasi error → toast

---

## ✅ PHASE 2: Auth & Session _(Day 3)_

- [ ] Zustand store `useAuthStore` (user, tokens, actions)
- [ ] Service auth: `login`, `register`, `refresh`, `logout`, `getProfile`
- [ ] Interceptor: auto-refresh saat 401, retry request, redirect login saat gagal
- [ ] Persist token (cookie httpOnly-friendly / localStorage sesuai kebutuhan)
- [ ] Halaman `/login` (form RHF + Zod, error state)
- [ ] Route guard: `PublicRoute`, `ProtectedRoute`, `AdminRoute` (cek role)
- [ ] Hook `useAuth` (status, isAdmin, logout)

---

## ✅ PHASE 3: Portal Publik _(Day 4-6)_

### 3.1 Landing Page
- [ ] Hero + CTA cari data
- [ ] Statistik ringkas dari `GET /public/releases/stats` (total by type/year)
- [ ] Section "Release Terbaru" (fetch list, limit kecil)
- [ ] Footer + navigasi

### 3.2 Halaman Releases
- [ ] Hook `useReleases(filters)` → `GET /public/releases`
- [ ] Grid/list card release (dataset & article beda tampilan)
- [ ] Filter: `type`, `year`, `category` + search `q`
- [ ] Sinkron filter ↔ URL query params
- [ ] Pagination dari `meta`
- [ ] Loading skeleton + empty state + error state

### 3.3 Detail Release
- [ ] Route `/releases/[slug]` → `GET /public/releases/slug/:slug`
- [ ] Layout dataset: metadata + tombol **Download** (`file_url`)
- [ ] Layout article: cover + konten
- [ ] Metadata dinamis (title, description, Open Graph) untuk SEO

### 3.4 Standards & Status
- [ ] Halaman `/standards` → `GET /public/standards`, tandai aktif, lihat/unduh PDF
- [ ] Indikator status dari `GET /health`

---

## ✅ PHASE 4: Dashboard Admin — Layout & Releases _(Day 7-9)_

### 4.1 Layout Dashboard
- [ ] Route group `(dashboard)` dengan guard auth
- [ ] Sidebar navigasi + topbar (profil, logout, theme)
- [ ] Breadcrumb + page container reusable

### 4.2 Manajemen Releases
- [ ] Tabel releases (search, filter, sort, pagination)
- [ ] Form Create/Edit (RHF + Zod): title, type, year, category, description
- [ ] **Upload file** dataset & cover image (progress bar, validasi tipe/ukuran)
- [ ] Auto slug preview dari title
- [ ] `POST /protected/releases`, `PUT /protected/releases/:id`
- [ ] Soft-delete (admin only) + dialog konfirmasi
- [ ] Optimistic update / invalidate query setelah mutasi

---

## ✅ PHASE 5: Dashboard Admin — Standards, Users, Audit _(Day 10-11)_

### 5.1 Standards
- [ ] Tabel standards + CRUD (`POST/PUT /protected/standards/:id`)
- [ ] Upload PDF + toggle `is_active`

### 5.2 Users (Admin)
- [ ] `GET /admin/users` → tabel user
- [ ] Ubah role `PUT /admin/users/:id/role` (dropdown + konfirmasi)

### 5.3 Audit Logs (Admin)
- [ ] `GET /admin/audit-logs` → tabel read-only
- [ ] Filter tanggal / action / user + pagination

### 5.4 Profil
- [ ] Halaman profil dari `GET /protected/profile`

---

## ✅ PHASE 6: Polish, QA & Deploy _(Day 12-13)_

- [ ] Konsistensi loading/error/empty state di semua halaman
- [ ] Responsive check (mobile → desktop) + dark mode audit
- [ ] Aksesibilitas dasar (label, focus ring, kontras, alt text)
- [ ] SEO halaman publik (metadata, sitemap, OG image)
- [ ] Error boundary + halaman 404/500
- [ ] Lint & `tsc --noEmit` bersih
- [ ] README: setup env, run, build, struktur
- [ ] Build & deploy (Vercel / Docker), set env production

---

## 📊 Priority Matrix

| Priority | Fitur |
|----------|-------|
| **P0 (Must Have)** | Setup, API layer, Auth+guard, List & Detail Releases, Landing |
| **P1 (High)** | Dashboard layout, CRUD Releases + upload, Standards publik |
| **P2 (Medium)** | Standards CRUD, Users role mgmt, Audit logs, Profil |
| **P3 (Nice to Have)** | SEO/OG lanjutan, animasi, PWA, i18n |

## 🎯 Success Criteria

- [ ] Semua halaman terhubung ke API nyata (bukan mock)
- [ ] Auth + refresh token + role guard berjalan mulus
- [ ] Filter & pagination tersinkron URL, cache via React Query
- [ ] Responsive + dark mode + loading/error/empty states lengkap
- [ ] Lint & type-check lolos, README lengkap

## 📝 Notes

- Selaraskan tipe FE dengan struct backend; sesuaikan bila field API berbeda.
- Bungkus semua fetch dengan React Query (staleTime sesuai TTL cache backend).
- Simpan token secara aman; hindari expose di client bila memungkinkan.
- Commit tiap selesai satu phase.
