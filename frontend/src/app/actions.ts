"use server"
import { DefaultActionResponse, ErrorResponse, Task, User } from "@/model"
import { API_PROTECTED } from "@/schemes"
import { revalidatePath, revalidateTag } from "next/cache"
import { cookies } from "next/headers"


export const refreshTasks:
  () => Promise<DefaultActionResponse<Task[]>> =
  async () => {
    let token = cookies().get('token')?.value as string
    console.log(API_PROTECTED + '/tasks/user')
    let res = await fetch(API_PROTECTED + '/tasks/user',
    {
      headers: {
        "Authorization": "Bearer " + token
      }
    }
    )
    if (!res.ok) {
      let err: ErrorResponse = await res.json()
      console.log(err.message)
      // toast.error("Error loading tasks")
      return { error: err }
    }
    let tasks: Task[] = await res.json()
    
    // revalidateTag('tasks')
    return { data: tasks }
  }

export const getUser: () => Promise<DefaultActionResponse<User>>
  = async () => {
    let token = cookies().get('token')?.value as string
    let res = await fetch(API_PROTECTED + '/users',
      {
        headers: {
          "Authorization": "Bearer " + token
        }
      }
    )
    if (!res.ok) {
      let err: ErrorResponse = await res.json()
      console.log(err.message)
      // toast.error("Error loading tasks")
      return { error: err }
    }
    let user: User = await res.json()
    // console.log(user)
    return { data: user }

  }

export const completeTask: (taskID: string) => Promise<DefaultActionResponse>
  = async (taskID) => {
    let token = cookies().get('token')?.value as string
    let res = await fetch(API_PROTECTED + '/tasks/complete/' + taskID,
      {
        method: 'PATCH',
        headers: {
          "Authorization": "Bearer " + token
        },
        // next: {revalidate: 1000 }
      }
    )
    if (!res.ok) {
      let err: ErrorResponse = await res.json()
      console.log(err.message)
      // toast.error("Error loading tasks")
      return { error: err }
    }

    // console.log(tasks)
    // revalidatePath('/app')
    return {}
  }

export const uncompleteTask: (taskID: string) => Promise<DefaultActionResponse>
  = async (taskID) => {
    let token = cookies().get('token')?.value as string
    let res = await fetch(API_PROTECTED + '/tasks/uncomplete/' + taskID,
      {
        method: 'PATCH',
        headers: {
          "Authorization": "Bearer " + token
        },
        // next: {revalidate: 1000 }
      }
    )
    if (!res.ok) {
      let err: ErrorResponse = await res.json()
      console.log(err.message)
      // toast.error("Error loading tasks")
      return { error: err }
    }

    // console.log(tasks)
    // revalidatePath('/app')
    return {}

  }

  export const deleteTask: (taskID: string) => Promise<DefaultActionResponse>
    = async (taskID) => {
      let token = cookies().get('token')?.value as string
      let res = await fetch(API_PROTECTED + '/tasks/' + taskID,
        {
          method: 'DELETE',
          headers: {
            "Authorization": "Bearer " + token
          },
          // next: {revalidate: 1000 }
        }
      )
      if (!res.ok) {
        let err: ErrorResponse = await res.json()
        console.log(err.message)
        // toast.error("Error loading tasks")
        return { error: err }
      }
      return {}
    }
