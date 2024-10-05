import React from 'react'
import { MenuProvider } from '../components/MenuProvider'

export default function layout({children}: {children: React.ReactNode}) {
  return (
    <MenuProvider>
        {children}
    </MenuProvider>
  )
}
