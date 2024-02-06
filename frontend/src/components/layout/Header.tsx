"use client";
import { LogOut } from "lucide-react";
import ThemeSwitcher from "../ThemeSwitcher";
import { useRouter } from "next/navigation";
import { deleteCookie } from "cookies-next";

export function Header() {
  const router = useRouter();
  return (
    <div
      className="p-2 md:p-3 w-full h-fit bg-primary-200 z-[99]
    dark:bg-primary-800/80 shadow-md flex items-center gap-4"
    >
      <div
        className="text-2xl md:text-3xl font-bold bg-primary-50 w-fit h-fit 
     p-2 pr-1 rounded-2xl shadow-md text-orange-500 hover:scale-105 transition-transform duration-300
      border-4 border-orange-500"
        style={{ letterSpacing: "4px" }}
      >
        <span className="select-none">ST</span>
      </div>
      <h1 className="text-2xl font-bold ">Stash Task</h1>
      <div className="ml-auto flex items-center gap-3">
        <ThemeSwitcher />
        <button
          onClick={async () => {
            deleteCookie("token");
            router.replace("/login");
          }}
        >
          <LogOut className="w-10 h-10" />
        </button>
      </div>
    </div>
  );
}
