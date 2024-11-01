"use client";
import WordFadeIn from "@/components/magicui/word-fade-in";
import { Button } from "@/components/ui/button";
import {
  Carousel,
  CarouselContent,
  CarouselItem,
} from "@/components/ui/carousel";
import AutoScroll from "embla-carousel-auto-scroll";
import Link from "next/link";
import React from "react";
import { FcGoogle } from "react-icons/fc";
import { BiLogoGmail } from "react-icons/bi";
import { FaDiscord } from "react-icons/fa";
import { IoLogoGithub } from "react-icons/io";
import { FaSpotify } from "react-icons/fa";
import { FaTwitch } from "react-icons/fa";

import { Card, CardHeader } from "@/components/ui/card";
import { Footer } from "@/components/ui/footer";
import { env } from "@/lib/env";

export default function Home() {
  const plugin = React.useRef(AutoScroll({ startDelay: 0 }));
  const slogan = [
    "Connect and Automate Effortlessly",
    "Trigger empowers you to connect services seamlessly. Automate tasks and enhance productivity by turning your ideas into efficient workflows.",
  ];
  const services = [
    { name: "Gmail", icon: <BiLogoGmail className="mr-2 hidden md:block" /> },
    { name: "Discord", icon: <FaDiscord className="mr-2 hidden md:block" /> },
    { name: "Github", icon: <IoLogoGithub className="mr-2 hidden md:block" /> },
    { name: "Spotify", icon: <FaSpotify className="mr-2 hidden md:block" /> },
    { name: "Twitch", icon: <FaTwitch className="mr-2 hidden md:block" /> },
  ];

  // TODO: Add href to start with google
  return (
    <div className="flex flex-1 flex-col w-full justify-center">
      <div className="flex flex-col items-center justify-start text-center text-black dark:text-white pt-20 gap-y-5 w-full">
        <div className="text-5xl font-bold mb-4">
          <WordFadeIn words={slogan[0]} as="h1" />
        </div>
        <div className="text-xl font-bold max-w-2xl">
          <WordFadeIn as="p" words={slogan[1]} />
        </div>
        <div className="max-w-md mx-auto flex flex-col md:flex-row gap-x-7">
          <Button
            className="w-full rounded-full text-lg py-6 px-12 mt-5 bg-orange-600 hover:bg-orange-700 text-white hover:text-white"
            variant="outline"
            asChild
          >
            <Link href="/auth?type=register">Start with Email</Link>
          </Button>
          <Button
            className="w-full rounded-full border-black bg-white text-lg p-6 mt-5"
            variant="outline"
            asChild
          >
            <Link
              href={`${env.NEXT_PUBLIC_AUTH_SERVICE_URL}/api/oauth2/login?provider=google&redirect=${env.NEXT_PUBLIC_WEB_URL}/home`}
              className="flex items-center"
            >
              <FcGoogle className="mr-2 text-2xl" /> Start with Google
            </Link>
          </Button>
        </div>
        <div className="flex flex-col w-1/2 mt-10">
          <Carousel
            opts={{
              loop: true,
            }}
            plugins={[plugin.current]}
            onMouseEnter={plugin.current.stop}
            onMouseLeave={() => plugin.current.play(0)}
          >
            <CarouselContent className="flex">
              {services.concat(services).map((item, index) => (
                <CarouselItem key={index} className="basis-1/2 md:basis-1/4">
                  <div className="p-1">
                    <span className="text-xl md:text-3xl font-semibold flex items-center justify-center text-muted-foreground">
                      {item.icon}
                      {item.name}
                    </span>
                  </div>
                </CarouselItem>
              ))}
            </CarouselContent>
          </Carousel>
        </div>
        <Card className="flex w-2/3 mt-10 items-center justify-center my-6">
          <CardHeader className="w-full h-full p-0">
            <video
              autoPlay
              muted
              loop
              className="w-full h-full object-cover rounded-md"
            >
              <source src="/video_placeholder.mov" />
            </video>
          </CardHeader>
        </Card>
      </div>
      <Footer />
    </div>
  );
}
