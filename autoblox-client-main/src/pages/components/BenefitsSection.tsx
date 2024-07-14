import React from "react";
import { m } from "framer-motion";

const variants = {
  visible: { opacity: 1, x: 0 },
  hidden: { opacity: 0, x: -25 },
};

function BenefitsSection() {
  return (
    <section className="bg-zinc-950/10 text-center px-16 py-16 gap-16 flex flex-col md:flex-row items-center justify-around">
      <m.div
        initial="hidden"
        whileInView="visible"
        viewport={{ once: true, margin: "-15px", amount: "all" }}
        transition={{ delay: 150 / 1000, duration: 400 / 1000 }}
        variants={variants}
        className="max-w-sm">
        <h4 className="text-3xl font-semibold">Safety First</h4>
        <p className="text-sm">
          AutoBlox mirrors human actions, ensuring safe gameplay on Roblox
          without the need for scripts or injectors.
        </p>
      </m.div>
      <m.div
        initial="hidden"
        whileInView="visible"
        viewport={{ once: true, margin: "-25px", amount: "all" }}
        transition={{ delay: 250 / 1000, duration: 400 / 1000 }}
        variants={variants}
        className="max-w-sm">
        <h4 className="text-3xl font-semibold">Swift Performance</h4>
        <p className="text-sm">
          AutoBlox delivers rapid farming and grinding performance, making your
          Roblox experience richer and more enjoyable.
        </p>
      </m.div>
      <m.div
        initial="hidden"
        whileInView="visible"
        viewport={{ once: true, margin: "-35px", amount: "all" }}
        transition={{
          delay: 350 / 1000,
          margin: "-50px",
          duration: 400 / 1000,
        }}
        variants={variants}
        className="max-w-sm">
        <h4 className="text-3xl font-semibold">Proven Reliability</h4>
        <p className="text-sm">
          AutoBlox has won the trust of Roblox users worldwide, offering a
          worry-free automation experience.
        </p>
      </m.div>
    </section>
  );
}

export default BenefitsSection;
