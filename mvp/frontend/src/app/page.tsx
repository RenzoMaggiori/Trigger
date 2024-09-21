import { Navbar } from "@/components/Navbar";
import { Button } from "@/components/ui/button";
import { FcGoogle } from "react-icons/fc";

export default function Home() {
  return (
    <div className="flex flex-1 w-full bg-gray-200 dark:bg-zinc-900 justify-center">
      <div className="flex flex-col items-center justify-start text-center text-black dark:text-white pt-20 gap-y-5">
        <h1 className="text-5xl font-bold mb-4">
          Connect and Automate Effortlessly
        </h1>
        <p className="text-xl font-bold max-w-2xl">
          Trigger empowers you to connect services seamlessly. Automate tasks and enhance productivity by turning your ideas into efficient workflows.
        </p>
          <Button className="rounded-full border-black bg-white text-lg p-6" variant="outline"><FcGoogle className="mr-2"/> Sign in with Google</Button>
      </div>

    </div>
  );
}
