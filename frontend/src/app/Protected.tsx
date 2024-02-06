"use client";

import { useEffect } from "react";
import { useRouter } from "next/navigation";
import { getCookie } from 'cookies-next'

export default function ProtectedPage({ children }: { children: React.ReactNode }) {
  const router = useRouter();
  let token: string | undefined
  useEffect(() => {
    token = getCookie('token')

    if (!token) {
      router.replace("/login"); // If no token is found, redirect to login page
      return;
    }
  }, [router]);

  return children;
}
