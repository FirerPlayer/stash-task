import { FormDrawer } from "@/components/forms/FormDrawer";
import ProtectedPage from "../Protected";
import { Plus, X } from "lucide-react";
import CreateTaskForm from "@/components/forms/CreateTaskForm";
import Cookies from "js-cookie";
import TaskPainel from "@/components/tasks/TaskPainel";
import { refreshTasks } from "../actions";
import { getCookie } from "cookies-next";

export default async function Page() {
  return (
    <>
      <ProtectedPage>
        <div className="w-full h-[calc(100svh-84px)] flex flex-col items-center md:justify-center">
          <TaskPainel
            subTitleComponent={
              <FormDrawer
                trigger={
                  <button className="btn !pl-2 !p-3">
                    <Plus className="w-8 h-8" />
                    New Task
                  </button>
                }
                content={
                  <div className="max-w-md w-full h-full mx-auto flex flex-col overflow-auto p-4 rounded-t-[10px]">
                    <CreateTaskForm />
                  </div>
                }
              />
            }
          />
        </div>
      </ProtectedPage>
    </>
  );
}
