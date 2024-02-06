"use client";
import { useContext, useEffect, useState } from "react";
import { useForm, SubmitHandler } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { API_PROTECTED, updateTask } from "@/schemes";
import { DefaultActionResponse, ErrorResponse, Task } from "@/model";
import { toast } from "sonner";
import { getCookie } from "cookies-next";
import { useRouter } from "next/navigation";
import { Eye, Trash, Upload } from "lucide-react";
import { isValidDate } from "@/lib";

type Inputs = z.infer<typeof updateTask>;

export default function TaskViewForm({
  task,
  deleteTask,
}: {
  task: Task;
  deleteTask: (taskID: string) => Promise<DefaultActionResponse>;
}) {
  // const [data, setData] = useState<Inputs>();
  const [priority, setPriority] = useState(3);
  const router = useRouter();

  const {
    register,
    handleSubmit,
    watch,
    reset,
    formState: { errors },
  } = useForm<Inputs>({
    resolver: zodResolver(updateTask),
    defaultValues: task,
  });

  const processForm: SubmitHandler<Inputs> = async (data) => {
    const res = await fetch(API_PROTECTED + "/tasks/" + task.id, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + getCookie("token"),
      },
      body: JSON.stringify({ ...data, priority }),
    });

    if (!res.ok) {
      let err: ErrorResponse = await res.json();
      console.log(err.message);
      toast.error("Error updating task");
      return;
    }

    reset();
    toast.success("Task updated");
  };

  return (
    <section className="flex flex-col gap-6 min-w-[128px]">
      <h1 className="text-3xl font-bold">View Task</h1>
      <form
        onSubmit={handleSubmit(processForm)}
        className="flex flex-1 flex-col gap-4 w-full"
      >
        <label>
          <h4 className="mb-2 w-fit text-base">Title</h4>
          <input placeholder="Title" className="input" {...register("title")} />
        </label>
        {errors.title?.message && (
          <p className="text-sm text-red-400 font-semibold">
            {errors.title.message}
          </p>
        )}

        <label>
          <h4 className="mb-2 w-fit text-base">Description</h4>
          <input
            placeholder="Description"
            className="input"
            {...register("description")}
          />
        </label>
        {errors.description?.message && (
          <p className="text-sm text-red-400 font-semibold">
            {errors.description.message}
          </p>
        )}

        <label>
          <h4 className="mb-2 w-fit text-base">Priority</h4>
          <select
            name="priority"
            id="priority"
            defaultValue={3}
            className="input"
            onChange={(e) => {
              setPriority(parseInt(e.target.value));
            }}
          >
            <option value={1}>High</option>
            <option value={2}>Medium</option>
            <option value={3}>Low</option>
          </select>
        </label>
        {errors.priority?.message && (
          <p className="text-sm text-red-400 font-semibold">
            {errors.priority.message}
          </p>
        )}
        {/* <label> */}
        <h4 className="mb-2 w-fit text-base">
          {task.completedAt && isValidDate(task.completedAt) ? (
            <p className="text-base">
              <span className="mb-2">Completed at:</span> <br />
              {new Date(task.completedAt).toLocaleString()}
            </p>
          ) : (
            "Not completed"
          )}
        </h4>
        {/* <input
            type="checkbox"
            className="min-w-8 min-h-8 p-0 z-10"
            {...register("completed")}
          />
        </label>
        {errors.completed?.message && (
          <p className="text-sm text-red-400 font-semibold">
            {errors.completed.message}
          </p>
        )} */}
        <button
          onClick={async () => {
            const res = await deleteTask(task.id);
            if (res.error) {
              toast.error("Error deleting task");
              return;
            }
            router.refresh();
            toast.success("Task deleted");
          }}
          className="btn !bg-red-300 !text-red-500"
        >
          <Trash className="w-6 h-6" />
          Delete task
        </button>
        <button className="btn">
          <Upload className="w-6 h-6" />
          Update task
        </button>
      </form>
    </section>
  );
}
