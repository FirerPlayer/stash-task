"use client";
import { Plus, X } from "lucide-react";
import { Drawer } from "vaul";
import CreateTaskForm from "./CreateTaskForm";
import { useState } from "react";

type Props = {
  trigger: React.ReactNode;
  content: React.ReactNode;
};

export function FormDrawer({ trigger, content }: Props) {
  return (
    <Drawer.Root direction="right">
      <Drawer.Trigger asChild>{trigger}</Drawer.Trigger>
      <Drawer.Portal>
        <Drawer.Overlay className="fixed inset-0 bg-black/40" />
        <Drawer.Content className="flex flex-col fixed top-0 right-0 h-[100svh] bg-white dark:bg-slate-700 z-[99]">
          <Drawer.Close asChild>
            <button className="absolute right-0 top-2">
              <X className="w-8 h-8" />
            </button>
          </Drawer.Close>
          <div className="max-w-md w-full h-full mx-auto flex flex-col overflow-auto p-4 rounded-t-[10px]">
            {content}
          </div>
        </Drawer.Content>
      </Drawer.Portal>
    </Drawer.Root>
  );
}
