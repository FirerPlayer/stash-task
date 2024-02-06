"use client";

import { isValidDate } from "@/lib";
import { DefaultActionResponse, Task } from "@/model";
import { Eye, Trash } from "lucide-react";
import { useRouter } from "next/navigation";
import { ChangeEvent, useState } from "react";
import { FormDrawer } from "../forms/FormDrawer";
import { toast } from "sonner";
import TaskViewForm from "../forms/TaskViewForm";

export default function TaskItem({
  key,
  task,
  completed,
  completeTask,
  uncompleteTask,
  deleteTask,
}: // deleteTask,
{
  key: string;
  task: Task;
  completed: boolean;
  completeTask: (taskID: string) => Promise<DefaultActionResponse>;
  uncompleteTask: (taskID: string) => Promise<DefaultActionResponse>;
  deleteTask: (taskID: string) => Promise<DefaultActionResponse>;
}) {
  const [checkbox, setCheckbox] = useState(isValidDate(task.completedAt));
  const router = useRouter();
  const checkHandler = async (e: ChangeEvent<HTMLInputElement>) => {
    setCheckbox(e.target.checked);
    if (e.target.checked) {
      const res = await completeTask(task.id);
      if (res.error) {
        console.log(res.error);
        toast.error("Error on completing task");
      }
      router.refresh();
      toast.success("Task completed");
      return;
    }

    const res = await uncompleteTask(task.id);
    if (res.error) {
      console.log(res.error);
      toast.error("Error on uncompleting task");
    }
    router.refresh();
    toast.success("Task uncompleted");
    return;
  };

  return (
    <li
      key={key}
      className={
        (completed ? "!bg-primary-300/40 dark:!bg-primary-700/80" : "") +
        ` relative flex gap-4 items-center min-h-20 w-80 rounded-lg p-3
    bg-slate-300 dark:bg-slate-700 overflow-hidden`
      }
    >
      <input
        onChange={checkHandler}
        // onClick={(e) => e.stopPropagation()}
        type="checkbox"
        checked={checkbox}
        className="min-w-8 min-h-8 p-0 z-10"
      />
      <div className="flex flex-col gap-1 w-full max-w-60 z-10">
        <h3 className="text-lg font-semibold w-full overflow-text">
          {task.title}
        </h3>
        <p className="text-sm w-full overflow-text">{task.description}</p>
      </div>
      {/* <div
        hidden={!completed}
        className=" absolute w-full h-full  inset-0 z-auto"
      ></div> */}
      <FormDrawer
        trigger={
          <button className="btn !bg-transparent hover:!bg-slate-200/30">
            <Eye className="w-6 h-6" />
          </button>
        }
        content={<TaskViewForm deleteTask={deleteTask} task={task} />}
      />
    </li>
  );
}
