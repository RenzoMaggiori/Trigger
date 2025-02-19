"use client";

import React from "react";
import { FcGoogle } from "react-icons/fc";
import { FiGithub } from "react-icons/fi";
import { FaSpotify } from "react-icons/fa";
import { z } from "zod";
import { useMutation } from "@tanstack/react-query";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";

import { Button } from "@/components/ui/button";
import {
  Form,
  FormField,
  FormItem,
  FormLabel,
  FormControl,
  FormMessage,
} from "@/components/ui/form";
import { Card, CardContent, CardHeader } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { login } from "@/app/auth/lib/auth";
import { Providers } from "@/app/auth/components/providers";

const formSchema = z.object({
  email: z.string(),
  password: z.string(),
});

export function Login() {
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  const mutation = useMutation({
    mutationFn: (values: z.infer<typeof formSchema>) =>
      login(values.email, values.password),
    onSuccess: () => {
      form.reset({});
    },
  });

  const onSubmit = (values: z.infer<typeof formSchema>) => {
    mutation.mutate(values);
  };

  const inputs = [
    {
      name: "email",
      label: "Email",
      type: "email",
      placeholder: "john.doe@example.com",
    },
    {
      name: "password",
      label: "Password",
      type: "password",
      placeholder: "********",
    },
  ] as const;

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)}>
        <Card>
          <CardHeader className="text-xl font-bold">Log in</CardHeader>
          <CardContent className="space-y-5 text-xl">
            {/* Continue with Credentials */}
            {inputs.map(({ name, label, type, placeholder }, index) => (
              <FormField
                key={index}
                control={form.control}
                name={name}
                render={({ field }) => (
                  <FormItem className="space-y-1">
                    <FormLabel>{label}</FormLabel>
                    <FormControl>
                      <Input placeholder={placeholder} type={type} {...field} />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
            ))}
            <Button
              className="flex w-full items-center justify-center rounded-full bg-orange-600 hover:bg-orange-700"
              type="submit"
            >
              Log in
            </Button>
            <p className="flex items-center justify-center font-bold text-lg py-2">
              or
            </p>
            {/* Continue with Provider */}
            <Providers
              providers={[
                {
                  name: "google",
                  text: "Continue with google",
                  icon: <FcGoogle className="mr-2" />,
                },
                {
                  name: "github",
                  text: "Continue with Github",
                  className:
                    "bg-zinc-800 text-white hover:bg-zinc-950 hover:text-white",
                  icon: <FiGithub className="mr-2" />,
                },
                {
                  name: "spotify",
                  text: "Continue with Spotify",
                  className:
                    "bg-green-600 text-white hover:bg-green-700 hover:text-white",
                  icon: <FaSpotify className="mr-2" />,
                },
              ]}
            />
          </CardContent>
        </Card>
      </form>
    </Form>
  );
}
