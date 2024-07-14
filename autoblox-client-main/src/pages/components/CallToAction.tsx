import { m, useInView } from "framer-motion";

/* eslint-disable react/no-unescaped-entities */
function CallToAction() {
  return (
    <m.div
      initial="hidden"
      whileInView="visible"
      viewport={{ once: true, margin: "-50px" }}
      transition={{ delay: 150 / 1000, duration: 400 / 1000 }}
      variants={{
        visible: { opacity: 1, y: 0 },
        hidden: { opacity: 0, y: 25 },
      }}
      className="mt-32 px-12">
      <div className="relative max-w-4xl mx-auto">
        <div className="absolute rounded-full w-24 h-24 bg-red-500/60 blur-3xl left-8 top-8" />
        <div className="absolute rounded-full w-24 h-24 bg-indigo-500/60 blur-3xl right-8 bottom-8" />
        <div className="py-12 px-12 bg-[#151518]/75 rounded-lg backdrop-blur-3xl flex flex-col justify-center items-center cta-container">
          <h3 className="text-xl md:text-2xl font-semibold">
            What Are You Waiting For?
          </h3>
          <p className="text-gray-400 mt-2 text-sm md:text-base">
            It's about time you start farming 💸
          </p>
          <div className="flex flex-col gap-4 mt-6 w-full max-w-xl">
            <a
              href={`${process.env.NEXT_PUBLIC_API_URL}/download`}
              target="_blank"
              className="flex-1 py-3 px-3 bg-primary hover:opacity-80 text-white font-semibold rounded-md flex justify-center items-center gap-4 transition-opacity text-sm md:text-base">
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
              href={`${process.env.NEXT_PUBLIC_API_URL}/discord`}
              target="_blank"
              className="flex-1 py-3 px-3 bg-indigo-500 hover:opacity-80 text-white font-semibold rounded-md flex justify-center items-center gap-4 transition-opacity text-sm md:text-base">
              Join The Community
              <svg
                fill="currentColor"
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 127.14 96.36"
                className="w-5 h-5">
                <g id="图层_2" data-name="图层 2">
                  <g id="Discord_Logos" data-name="Discord Logos">
                    <g
                      id="Discord_Logo_-_Large_-_White"
                      data-name="Discord Logo - Large - White">
                      <path d="M107.7,8.07A105.15,105.15,0,0,0,81.47,0a72.06,72.06,0,0,0-3.36,6.83A97.68,97.68,0,0,0,49,6.83,72.37,72.37,0,0,0,45.64,0,105.89,105.89,0,0,0,19.39,8.09C2.79,32.65-1.71,56.6.54,80.21h0A105.73,105.73,0,0,0,32.71,96.36,77.7,77.7,0,0,0,39.6,85.25a68.42,68.42,0,0,1-10.85-5.18c.91-.66,1.8-1.34,2.66-2a75.57,75.57,0,0,0,64.32,0c.87.71,1.76,1.39,2.66,2a68.68,68.68,0,0,1-10.87,5.19,77,77,0,0,0,6.89,11.1A105.25,105.25,0,0,0,126.6,80.22h0C129.24,52.84,122.09,29.11,107.7,8.07ZM42.45,65.69C36.18,65.69,31,60,31,53s5-12.74,11.43-12.74S54,46,53.89,53,48.84,65.69,42.45,65.69Zm42.24,0C78.41,65.69,73.25,60,73.25,53s5-12.74,11.44-12.74S96.23,46,96.12,53,91.08,65.69,84.69,65.69Z" />
                    </g>
                  </g>
                </g>
              </svg>
            </a>
          </div>
        </div>
      </div>
    </m.div>
  );
}

export default CallToAction;
