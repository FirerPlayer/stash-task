import type { Metadata } from "next";
import { Ubuntu } from "next/font/google";
import { Toaster } from 'sonner';
import "./globals.css";
import { ProviderTheme } from "@/components/ThemeProvider";
import { Header } from "@/components/layout/Header";
import { CookiesProvider } from "react-cookie";
import 'notyf/notyf.min.css';
const ubuntu = Ubuntu({ weight: "400", subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Stash Task",
  description: "Seu gerenciador de tarefas",
};
export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="pt-br" suppressHydrationWarning>
      <body className={ubuntu.className + " text-gray-700 dark:text-gray-50 bg-slate-200 dark:bg-slate-900"}>
        <ProviderTheme>
          <Header />
          <section>
            {children}
          </section>
        </ProviderTheme>
        <Toaster closeButton duration={4000} position="top-left" className="z-[9999]" />
      </body>
    </html>
  );
}
