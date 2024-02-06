"use client";
import { useContext, useEffect, useState } from "react";
import { useForm, SubmitHandler } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { API_PROTECTED, createTask } from "@/schemes";
import { ErrorResponse } from "@/model";
import { toast } from "sonner";
import { getCookie } from "cookies-next";

type Inputs = z.infer<typeof createTask>;

export default function CreateTaskForm() {
  const [data, setData] = useState<Inputs>();
  const [priority, setPriority] = useState(3);

  const {
    register,
    handleSubmit,
    watch,
    reset,
    formState: { errors },
  } = useForm<Inputs>({
    resolver: zodResolver(createTask),
  });

  const processForm: SubmitHandler<Inputs> = async (data) => {
    const res = await fetch(API_PROTECTED + "/tasks", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + getCookie("token"),
      },
      body: JSON.stringify({ ...data, priority }),
    });

    if (!res.ok) {
      let err: ErrorResponse = await res.json();
      console.log(err.message);
      toast.error("Error creating task");
      return;
    }

    reset();
    toast.success("Task created");
  };

  return (
    <section className="flex flex-col gap-6">
      <h1 className="text-3xl font-bold">New Task</h1>
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
        <button className="btn">Create task</button>
      </form>
    </section>
  );
}
