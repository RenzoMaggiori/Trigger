"use client"
import React from 'react'
import { MenuProvider } from '../components/MenuProvider'

export default function layout({ children }: { children: React.ReactNode }) {
  const [mounted, setMounted] = React.useState(false);

  React.useEffect(() => {
    setMounted(true);
  }, []);

  if (!mounted) return null;
  return (
    <MenuProvider>
      {children}
    </MenuProvider>
  )
}
