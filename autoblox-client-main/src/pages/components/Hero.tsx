import Image from "next/image";
import { Typewriter } from "react-simple-typewriter";

function Hero() {
  return (
    <main
      style={{
        backgroundImage: 'url("/background.webp")',
        backgroundSize: "cover",
        backgroundPosition: "center",
      }}
      className="w-full h-[700px] relative text-center">
      <div className="absolute top-0 left-0 pt-24 px-8 flex flex-col items-center justify-evenly w-full h-full bg-zinc-900/20 backdrop-blur-lg">
        <h2 className="text-white text-4xl lg:text-5xl font-bold">
          <Typewriter
            cursor
            cursorBlinking
            loop
            words={[
              "Trusted Roblox Autofarms",
              "Best Roblox Automation Tools",
              "Fastest Roblox Autofarms",
            ]}
          />
        </h2>
        <Image
          src="/product.png"
          alt="AutoBlox App Image"
          className="rounded-lg w-full max-w-xl border-2 border-primary drop-shadow-2xl"
          width={400}
          height={250}
        />
        <div className="flex flex-col gap-4 w-full max-w-xl sm:flex-row -mt-6">
          <a
            href={`${process.env.NEXT_PUBLIC_API_URL}/download`}
            target="_blank"
            className="flex-1 py-2 px-8 bg-primary hover:bg-primary-400 text-white font-medium rounded-lg flex justify-center items-center gap-4 transition-colors">
            Download Now
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 20 20"
              fill="currentColor"
              className="w-5 h-5">
              <path d="M10.75 2.75a.75.75 0 00-1.5 0v8.614L6.295 8.235a.75.75 0 10-1.09 1.03l4.25 4.5a.75.75 0 001.09 0l4.25-4.5a.75.75 0 00-1.09-1.03l-2.955 3.129V2.75z" />
              <path d="M3.5 12.75a.75.75 0 00-1.5 0v2.5A2.75 2.75 0 004.75 18h10.5A2.75 2.75 0 0018 15.25v-2.5a.75.75 0 00-1.5 0v2.5c0 .69-.56 1.25-1.25 1.25H4.75c-.69 0-1.25-.56-1.25-1.25v-2.5z" />
            </svg>
          </a>
          <a
            href={`${process.env.NEXT_PUBLIC_API_URL}/key`}
            target="_blank"
            className="flex-1 py-2 px-8 bg-primary/40 hover:bg-primary/80 border border-primary text-white font-medium rounded-lg flex justify-center items-center gap-4 transition-colors">
            Get Key
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 20 20"
              fill="currentColor"
              className="w-5 h-5">
              <path
                fillRule="evenodd"
                d="M8 7a5 5 0 113.61 4.804l-1.903 1.903A1 1 0 019 14H8v1a1 1 0 01-1 1H6v1a1 1 0 01-1 1H3a1 1 0 01-1-1v-2a1 1 0 01.293-.707L8.196 8.39A5.002 5.002 0 018 7zm5-3a.75.75 0 000 1.5A1.5 1.5 0 0114.5 7 .75.75 0 0016 7a3 3 0 00-3-3z"
                clipRule="evenodd"
              />
            </svg>
          </a>
        </div>
      </div>
    </main>
  );
}

export default Hero;
