import WordFadeIn from "@/components/magicui/word-fade-in";
import { Navbar } from "@/components/Navbar";
import { Button } from "@/components/ui/button";
import { FcGoogle } from "react-icons/fc";

export default function Home() {
  return (
    <div className="flex flex-1 w-full bg-gray-200 dark:bg-zinc-900 justify-center">
      <div className="flex flex-col items-center justify-start text-center text-black dark:text-white pt-20 gap-y-5">
        <div className="text-5xl font-bold mb-4">
          <WordFadeIn words="Connect and Automate Effortlessly" as="h1" />
        </div>
        <div className="text-xl font-bold max-w-2xl">
          <WordFadeIn as="p" words="Trigger empowers you to connect services seamlessly. Automate tasks and enhance productivity by turning your ideas into efficient workflows." />
        </div>
          <Button className="rounded-full border-black bg-white text-lg p-6 mt-5" variant="outline"><FcGoogle className="mr-2"/> Sign in with Google</Button>
      </div>

    </div>
  );
}
