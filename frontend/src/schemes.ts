import { z } from "zod"

export const API_PROTECTED = "http://localhost:8080/api/private"
export const API = "http://localhost:8080/api"

export const createTask = z.object({
  title: z.string().min(1, "Title is required"),
  description: z.string().optional(),
  priority: z.number().optional(),
})
export const updateTask = z.object({
  title: z.string().min(1, "Title is required"),
  description: z.string().optional(),
  priority: z.number().optional(),
  completed: z.boolean().optional(),
})



const MB_BYTES = 1000000; // Number of bytes in a megabyte.
const MAX_SIZE = 5 * MB_BYTES;

// This is the list of mime types you will accept with the schema
const ACCEPTED_MIME_TYPES = ["image/gif", "image/jpeg", "image/png"];
export const createUser = z.object({
  email: z.string(),
  avatar: z.instanceof(File).superRefine((f, ctx) => {
    // First, add an issue if the mime type is wrong.
    if (!ACCEPTED_MIME_TYPES.includes(f.type)) {
      ctx.addIssue({
        code: z.ZodIssueCode.custom,
        message: `File must be one of [${ACCEPTED_MIME_TYPES.join(
          ", "
        )}] but was ${f.type}`
      });
    }
    // Next add an issue if the file size is too large.
    if (f.size > MAX_SIZE) {
      ctx.addIssue({
        code: z.ZodIssueCode.too_big,
        type: "array",
        message: `The file must not be larger than ${MAX_SIZE} bytes: ${f.size}`,
        maximum: MAX_SIZE,
        inclusive: true
      });
    }
  }).optional(),
  username: z.string(),
  bio: z.string().optional()
})

export const userLogin = z.object({
  email: z.string().email().min(1, "Email is required"),
  password: z.string().min(1, "Password is required")
})

export const userRegister = z.object({
  email: z.string().email().min(1, "Email is required"),
  username: z.string().min(1, "Username is required"),
  password: z.string().min(1, "Password is required"),
  bio: z.string().optional()
})