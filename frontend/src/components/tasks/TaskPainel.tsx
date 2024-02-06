import {
  completeTask,
  getUser,
  refreshTasks,
  uncompleteTask,
  deleteTask,
} from "@/app/actions";
import { Task, type ErrorResponse, User } from "@/model";
import { CalendarX } from "lucide-react";
import { ChangeEvent } from "react";
import TaskItem from "./TaskItem";
import { isValidDate } from "@/lib";

function TaskList({
  tasks,
  notFoundText,
}: {
  tasks: Task[];
  notFoundText?: string;
}) {
  return (
    <ul className="flex flex-col gap-2 max-h-96 md:max-h-[640px] overflow-auto">
      {tasks.length > 0 &&
        tasks.map((task) => {
          return (
            <TaskItem
              key={task.id}
              task={task}
              completed={isValidDate(task.completedAt)}
              completeTask={completeTask}
              uncompleteTask={uncompleteTask}
              deleteTask={deleteTask}
            />
          );
        })}
      {tasks.length === 0 && (
        <div
          className="rounded-lg flex flex-col items-center justify-center 
        gap-2 w-80 h-96 md:h-[640px] bg-primary-400 bg-primary-800/80"
        >
          <span className="font-semibold text-lg">
            {notFoundText ?? "No tasks found"}
          </span>
          <CalendarX className="w-10 h-10" />
        </div>
      )}
    </ul>
  );
}

export default async function TaskPainel({
  subTitleComponent,
}: {
  subTitleComponent?: JSX.Element;
}) {
  let respTasks = await refreshTasks();
  if (respTasks.error) {
    return (
      <div className="flex flex-col gap-2">
        <h1>Something went wrong</h1>
        <p>{respTasks.error.message}</p>
      </div>
    );
  }
  let respUser = await getUser();
  if (respUser.error) {
    return (
      <div className="flex flex-col gap-2">
        <h1>Something went wrong</h1>
        <p>{respUser.error.message}</p>
      </div>
    );
  }
  let tasks = respTasks.data as Task[];
  let user = respUser.data as User;
  if (!tasks) {
    tasks = [];
  }
  let completedTasks = tasks.filter((task) => isValidDate(task.completedAt));
  let uncompletedTasks = tasks.filter((task) => !isValidDate(task.completedAt));

  return (
    <div className="flex flex-col gap-2 bg-red-300a w-fit">
      <h1 className="capitalize font-semibold text-3xl">
        {user.username}'s Tasks{" "}
      </h1>
      {subTitleComponent}
      <div className="grid grid-cols-1 grid-flow-row md:grid-cols-2 md:grid-rows-[auto_1fr] gap-2">
        <h2 className="font-semibold text-xl">Uncompleted</h2>
        <h2 className="font-semibold text-xl hidden md:block">Completed</h2>
        <TaskList tasks={uncompletedTasks} />
        <h2 className="font-semibold text-xl block md:hidden">Completed</h2>
        <TaskList tasks={completedTasks} />
      </div>
    </div>
  );
}
