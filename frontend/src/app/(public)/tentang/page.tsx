import type { Metadata } from "next";
import { Breadcrumb } from "@/components/ui/breadcrumb";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

export const metadata: Metadata = {
  title: "Tentang - Satudata",
  description: "Pelajari lebih lanjut tentang Portal Data Terbuka Satudata, platform data statistik dan informasi publik Indonesia.",
  openGraph: {
    title: "Tentang - Satudata",
    description: "Portal data statistik dan informasi publik Indonesia.",
  },
};

export default function TentangPage() {
  return (
    <div className="container py-8">
      <Breadcrumb
        items={[{ label: "Tentang" }]}
        className="mb-6"
      />

      <div className="max-w-3xl mx-auto space-y-8">
        {/* Hero */}
        <div className="text-center space-y-4">
          <h1 className="text-4xl font-bold tracking-tight">Tentang Satudata</h1>
          <p className="text-lg text-muted-foreground">
            Portal data terbuka yang menyajikan dataset, artikel analitis, infografis, dan standar data dari berbagai unit kerja.
          </p>
        </div>

        {/* Visi & Misi */}
        <Card>
          <CardHeader>
            <CardTitle>Visi & Misi</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <div>
              <h3 className="font-semibold mb-2">Visi</h3>
              <p className="text-muted-foreground">
                Menjadi portal data terbuka yang terpercaya, transparan, dan mudah diakses untuk mendukung
                pengambilan keputusan berbasis data di Indonesia.
              </p>
            </div>
            <div>
              <h3 className="font-semibold mb-2">Misi</h3>
              <ul className="list-disc pl-5 space-y-1 text-muted-foreground">
                <li>Menyediakan akses terbuka terhadap data statistik dan informasi publik.</li>
                <li>Mendorong transparansi dan akuntabilitas melalui keterbukaan data.</li>
                <li>Mendukung inovasi dan penelitian berbasis data.</li>
                <li>Memfasilitasi kolaborasi antar unit kerja dalam berbagi data.</li>
              </ul>
            </div>
          </CardContent>
        </Card>

        {/* Apa yang Kami Sediakan */}
        <Card>
          <CardHeader>
            <CardTitle>Apa yang Kami Sediakan</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="grid md:grid-cols-2 gap-6">
              <div className="space-y-2">
                <h3 className="font-semibold">📊 Dataset</h3>
                <p className="text-sm text-muted-foreground">
                  Kumpulan data statistik dari berbagai sektor yang dapat diunduh dalam format CSV, JSON, XLSX, dan Parquet.
                </p>
              </div>
              <div className="space-y-2">
                <h3 className="font-semibold">📝 Artikel & Infografis</h3>
                <p className="text-sm text-muted-foreground">
                  Analisis mendalam dan visualisasi data dalam bentuk artikel dan infografis interaktif.
                </p>
              </div>
              <div className="space-y-2">
                <h3 className="font-semibold">📋 Standar Data</h3>
                <p className="text-sm text-muted-foreground">
                  Dokumen standar data yang menjadi acuan dalam pengelolaan data di lingkungan pemerintahan.
                </p>
              </div>
              <div className="space-y-2">
                <h3 className="font-semibold">🔍 Pencarian & Filter</h3>
                <p className="text-sm text-muted-foreground">
                  Fitur pencarian dan filter yang memudahkan Anda menemukan data yang dibutuhkan.
                </p>
              </div>
            </div>
          </CardContent>
        </Card>

        {/* Prinsip Data */}
        <Card>
          <CardHeader>
            <CardTitle>Prinsip Data Terbuka</CardTitle>
          </CardHeader>
          <CardContent>
            <ol className="list-decimal pl-5 space-y-3 text-muted-foreground">
              <li><strong className="text-foreground">Lengkap</strong> — Data dipublikasikan secara utuh tanpa redaksi.</li>
              <li><strong className="text-foreground">Primer</strong> — Data bersumber langsung dari instansi pemilik data.</li>
              <li><strong className="text-foreground">Tepat Waktu</strong> — Data diperbarui secara berkala sesuai frekuensi yang ditetapkan.</li>
              <li><strong className="text-foreground">Mudah Diakses</strong> — Data tersedia dalam format yang mudah diunduh dan diolah.</li>
              <li><strong className="text-foreground">Terproses</strong> — Data dapat diproses dengan mesin untuk analisis lebih lanjut.</li>
              <li><strong className="text-foreground">Non-Diskriminatif</strong> — Data tersedia untuk siapa saja tanpa persyaratan khusus.</li>
              <li><strong className="text-foreground">Non-Proprieter</strong> — Data dalam format yang tidak terikat pada vendor tertentu.</li>
              <li><strong className="text-foreground">Tanpa Lisensi</strong> — Data tidak dibatasi oleh hak paten atau lisensi yang membatasi penggunaan.</li>
            </ol>
          </CardContent>
        </Card>

        {/* Kontak */}
        <Card>
          <CardHeader>
            <CardTitle>Hubungi Kami</CardTitle>
          </CardHeader>
          <CardContent className="space-y-2 text-muted-foreground">
            <p>Email: <a href="mailto:info@satudata.go.id" className="text-primary hover:underline">info@satudata.go.id</a></p>
            <p>Alamat: Kantor Kementerian Komunikasi dan Informatika RI, Jakarta Pusat</p>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
