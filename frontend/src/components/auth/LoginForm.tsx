"use client"
import { ErrorResponse } from '@/model';
import { API, userLogin } from '@/schemes';
import { zodResolver } from '@hookform/resolvers/zod';
import { useRouter } from 'next/navigation';
import { SubmitHandler, useForm } from 'react-hook-form';
import { toast } from 'sonner';
import { z } from 'zod';
import Cookies from "js-cookie";
import { setCookie } from 'cookies-next';

type Inputs = z.infer<typeof userLogin>

export default function LoginForm() {
  const {
    register,
    handleSubmit,
    watch,
    reset,
    formState: { errors }
  } = useForm<Inputs>({
    resolver: zodResolver(userLogin)
  })
  const router = useRouter()


  const processForm: SubmitHandler<Inputs> = async data => {
    const res = await fetch(API + '/users/login', {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),

    })

    if (!res.ok) {
      let err: ErrorResponse = await res.json()
      console.log(err.message)
      toast.error("Invalid credentials")
      return
    }
    const token = (await res.json()).token

    reset()
    toast.success("Login successful")
    router.push('/app')
    setCookie("token", token, { maxAge: 60 * 60 * 24 * 30 })
  }


  return (
    <section className='relative w-full h-[calc(100svh-84px)]'>
      <div className="absolute flex flex-col gap-3 items-center h-fit w-full 
      top-1/2 -translate-y-1/2 left-1/2 -translate-x-1/2 md:p-0 px-8">
        <h1 className='text-3xl font-bold mb-5'>Login</h1>
        <form
          onSubmit={handleSubmit(processForm)}
          className='flex flex-1 flex-col gap-4 w-full max-w-md '
        >
          <input
            id='email'
            type='email'
            placeholder='Email'
            className='input'
            {...register('email')}
          />
          {errors.email?.message && (
            <p className='text-sm text-red-400 font-semibold'>{errors.email?.message}</p>
          )}

          <input
            placeholder='Password'
            id='password'
            className='input'
            type='password'
            {...register('password')}
          />
          {errors.password?.message && (
            <p className='text-sm text-red-400 font-semibold'>
              {errors.password.message}
            </p>
          )}
          <button className='btn flex items-center justify-center py-4'>Login</button>
        </form>
      </div>
    </section>
  )
}
