"use client";

import { useState } from "react";
import { ChevronDown } from "lucide-react";
import { Breadcrumb } from "@/components/ui/breadcrumb";
import { cn } from "@/lib/utils";

const faqs = [
  {
    question: "Apa itu Satudata?",
    answer: "Satudata adalah portal data terbuka yang menyediakan akses ke dataset statistik, artikel analitis, infografis, dan standar data dari berbagai unit kerja di Indonesia. Portal ini bertujuan untuk mendukung transparansi dan pengambilan keputusan berbasis data.",
  },
  {
    question: "Apakah data di Satudata gratis?",
    answer: "Ya, seluruh data yang tersedia di portal Satudata dapat diakses dan diunduh secara gratis oleh masyarakat umum tanpa biaya.",
  },
  {
    question: "Bagaimana cara mencari dataset?",
    answer: "Anda dapat menggunakan fitur pencarian pada halaman beranda atau halaman dataset. Gunakan kata kunci untuk mencari dataset tertentu, atau gunakan filter berdasarkan kategori, organisasi, format file, dan tahun untuk mempersempit hasil pencarian.",
  },
  {
    question: "Format file apa saja yang tersedia?",
    answer: "Dataset tersedia dalam berbagai format termasuk CSV, JSON, XLSX, dan Parquet. Artikel dan infografis tersedia dalam format HTML. Dokumen standar data umumnya tersedia dalam format PDF.",
  },
  {
    question: "Bagaimana cara mengunduh dataset?",
    answer: "Buka halaman detail dataset yang Anda inginkan, lalu klik tombol 'Download' pada resource yang tersedia. File akan langsung terunduh ke perangkat Anda.",
  },
  {
    question: "Apakah saya perlu mendaftar untuk mengakses data?",
    answer: "Tidak. Seluruh data publik dapat diakses tanpa perlu mendaftar atau login. Akun pengguna hanya diperlukan untuk kontributor data dan admin.",
  },
  {
    question: "Seberapa sering data diperbarui?",
    answer: "Frekuensi pembaruan data bervariasi tergantung pada jenis datanya. Beberapa dataset diperbarui secara bulanan, triwulanan, atau tahunan. Informasi frekuensi pembaruan tersedia di halaman detail masing-masing dataset.",
  },
  {
    question: "Apa yang dimaksud dengan Standar Data?",
    answer: "Standar Data adalah dokumen yang menjadi acuan dalam pengelolaan data, mencakup definisi, struktur, format, dan metadata yang harus dipenuhi oleh setiap unit kerja dalam menyampaikan datanya.",
  },
  {
    question: "Bagaimana jika saya menemukan kesalahan data?",
    answer: "Silakan hubungi tim Satudata melalui email info@satudata.go.id untuk melaporkan kesalahan data. Tim kami akan memverifikasi dan melakukan koreksi yang diperlukan.",
  },
  {
    question: "Bagaimana cara menjadi kontributor data?",
    answer: "Untuk menjadi kontributor data, silakan menghubungi admin Satudata melalui email info@satudata.go.id. Anda akan diberikan akses sebagai editor untuk mengelola dataset dari unit kerja Anda.",
  },
];

function FaqItem({ question, answer }: { question: string; answer: string }) {
  const [open, setOpen] = useState(false);

  return (
    <div className="border-b border-border last:border-0">
      <button
        onClick={() => setOpen(!open)}
        className="flex w-full items-center justify-between py-4 text-left transition-colors hover:text-foreground/80"
      >
        <span className="font-medium pr-4">{question}</span>
        <ChevronDown
          className={cn(
            "h-5 w-5 shrink-0 text-muted-foreground transition-transform duration-200",
            open && "rotate-180"
          )}
        />
      </button>
      <div
        className={cn(
          "overflow-hidden transition-all duration-200",
          open ? "pb-4" : "h-0"
        )}
      >
        <p className="text-muted-foreground">{answer}</p>
      </div>
    </div>
  );
}

export default function FAQPage() {
  return (
    <div className="container py-8">
      <Breadcrumb
        items={[{ label: "FAQ" }]}
        className="mb-6"
      />

      <div className="max-w-3xl mx-auto">
        <div className="text-center mb-8">
          <h1 className="text-4xl font-bold tracking-tight mb-4">
            Frequently Asked Questions
          </h1>
          <p className="text-lg text-muted-foreground">
            Pertanyaan yang sering diajukan tentang portal Satudata
          </p>
        </div>

        <div className="bg-card rounded-lg border border-border px-6">
          {faqs.map((faq, index) => (
            <FaqItem key={index} question={faq.question} answer={faq.answer} />
          ))}
        </div>

        <div className="mt-8 text-center text-sm text-muted-foreground">
          Masih punya pertanyaan?{" "}
          <a href="mailto:info@satudata.go.id" className="text-primary hover:underline">
            Hubungi kami
          </a>
        </div>
      </div>
    </div>
  );
}
