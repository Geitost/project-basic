function Footer() {
  return (
    <footer className="flex flex-col gap-4 sm:flex-row items-center justify-between py-6 px-8 bg-zinc-950/20">
      <div className="text-2xl font-medium">AutoBlox</div>
      <div className="flex gap-4">
        <a
          href={`${process.env.NEXT_PUBLIC_API_URL}/discord`}
          target="_blank"
          className="text-gray-300 hover:underline">
          Discord
        </a>
      </div>
    </footer>
  );
}
export default Footer;
