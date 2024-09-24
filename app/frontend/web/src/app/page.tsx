"use client"
import WordFadeIn from "@/components/magicui/word-fade-in";
import { Button } from "@/components/ui/button";
import { Carousel, CarouselContent, CarouselItem } from "@/components/ui/carousel";
import AutoScroll from "embla-carousel-auto-scroll";
import Image from "next/image";
import Link from "next/link";
import React from "react";
import { FcGoogle } from "react-icons/fc";
import { BiLogoGmail } from "react-icons/bi";
import { FaDiscord } from "react-icons/fa";
import { IoLogoGithub } from "react-icons/io";
import { FaSlack } from "react-icons/fa6";
import { PiMicrosoftOutlookLogo } from "react-icons/pi";

export default function Home() {
  const plugin = React.useRef(
    AutoScroll({ startDelay: 0 })
  )

  const services = [
    { name: "Gmail", icon: <BiLogoGmail className="mr-2 hidden md:block" /> },
    { name: "Discord", icon: <FaDiscord className="mr-2 hidden md:block" /> },
    { name: "Github", icon: <IoLogoGithub className="mr-2 hidden md:block" /> },
    { name: "Slack", icon: <FaSlack className="mr-2 hidden md:block" /> },
    { name: "Outlook", icon: <PiMicrosoftOutlookLogo className="mr-2 hidden md:block" /> },
  ]

  return (
    <div className="flex flex-1 w-full bg-gray-200 dark:bg-zinc-900 justify-center">
      <div className="flex flex-col items-center justify-start text-center text-black dark:text-white pt-20 gap-y-5 w-full">
        <div className="text-5xl font-bold mb-4">
          <WordFadeIn words="Connect and Automate Effortlessly" as="h1" />
        </div>
        <div className="text-xl font-bold max-w-2xl">
          <WordFadeIn as="p" words="Trigger empowers you to connect services seamlessly. Automate tasks and enhance productivity by turning your ideas into efficient workflows." />
        </div>
        <div className="max-w-md mx-auto flex flex-col md:flex-row gap-x-7">
          <Button className="w-full rounded-full text-lg py-6 px-12 mt-5 bg-orange-600 hover:bg-orange-700 text-white hover:text-white" variant="outline">
            Start with Email
          </Button>
          <Button className="w-full rounded-full border-black bg-white text-lg p-6 mt-5" variant="outline" asChild>
            <Link href="http://localhost:8000/api/auth/gmail/provider" className="flex items-center">
              <FcGoogle className="mr-2 text-2xl" /> Start with Google
            </Link>
          </Button>
        </div>
        <div className="flex flex-col w-1/2 mt-10">
          <Carousel
            opts={{
              loop: true
            }}
            plugins={[plugin.current]}
            onMouseEnter={plugin.current.stop}
            onMouseLeave={() => plugin.current.play(0)}
          >
            <CarouselContent className="flex">
              {services.map((item, index) => (
                <CarouselItem key={index} className="basis-1/2 md:basis-1/4">
                  <div className="p-1">
                    <span className="text-xl md:text-3xl font-semibold flex items-center justify-center text-muted-foreground">{item.icon}{item.name}</span>
                  </div>
                </CarouselItem>
              ))}
            </CarouselContent>
          </Carousel>
        </div>
      </div>

    </div>
  );
}
