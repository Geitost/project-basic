import CountUp from "react-countup";
import { m } from "framer-motion";
import { useEffect, useState } from "react";

function StatsSection() {
  const [users, setUsers] = useState(0);
  const [keys, setKeys] = useState(0);

  useEffect(() => {
    async function updateStats() {
      try {
        const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/stats`);
        const data = await res.json();
        if (data?.success) {
          setUsers(data?.keyCount);
          setKeys(data?.checkoutCount);
        }
      } catch (err) {
        console.error(err);
      }
    }

    updateStats();
  }, []);

  return (
    <section className="flex flex-col gap-16 md:flex-row items-center py-16 justify-evenly">
      <m.div
        initial="hidden"
        whileInView="visible"
        viewport={{ once: true, amount: "all" }}
        transition={{ delay: 50 / 1000, duration: 600 / 1000 }}
        variants={{
          visible: { opacity: 1, y: 0 },
          hidden: { opacity: 0, y: 25 },
        }}
        className="flex items-center gap-4 text-primary">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          strokeWidth={1.5}
          stroke="currentColor"
          className="w-20 h-20 sm:w-24 sm:h-24">
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            d="M15.182 15.182a4.5 4.5 0 01-6.364 0M21 12a9 9 0 11-18 0 9 9 0 0118 0zM9.75 9.75c0 .414-.168.75-.375.75S9 10.164 9 9.75 9.168 9 9.375 9s.375.336.375.75zm-.375 0h.008v.015h-.008V9.75zm5.625 0c0 .414-.168.75-.375.75s-.375-.336-.375-.75.168-.75.375-.75.375.336.375.75zm-.375 0h.008v.015h-.008V9.75z"
          />
        </svg>
        <div>
          <h3 className="font-semibold text-6xl sm:text-7xl">
            <CountUp end={users} duration={5} delay={50} enableScrollSpy />+
          </h3>
          <span className="text-gray-300 text-lg sm:text-xl font-medium">
            Happy Users
          </span>
        </div>
      </m.div>
      <m.div
        initial="hidden"
        whileInView="visible"
        viewport={{ once: true, margin: "-25px" }}
        transition={{ delay: 600 / 2 / 1000, duration: 1000 / 1000 }}
        variants={{
          visible: { opacity: 1, y: 0 },
          hidden: { opacity: 0, y: 25 },
        }}
        className="flex items-center gap-4 text-primary">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          strokeWidth={1.5}
          stroke="currentColor"
          className="w-20 h-20 sm:w-24 sm:h-24">
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            d="M15.75 5.25a3 3 0 013 3m3 0a6 6 0 01-7.029 5.912c-.563-.097-1.159.026-1.563.43L10.5 17.25H8.25v2.25H6v2.25H2.25v-2.818c0-.597.237-1.17.659-1.591l6.499-6.499c.404-.404.527-1 .43-1.563A6 6 0 1121.75 8.25z"
          />
        </svg>

        <div>
          <h3 className="font-semibold text-6xl sm:text-7xl">
            <CountUp end={keys} duration={5} delay={600 / 2} enableScrollSpy />+
          </h3>
          <span className="text-gray-300 text-lg sm:text-xl font-medium">
            Keys Bought
          </span>
        </div>
      </m.div>
    </section>
  );
}

export default StatsSection;
