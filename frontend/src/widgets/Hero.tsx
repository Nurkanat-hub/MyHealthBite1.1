import React from "react";
import bread from "../shared/assets/7894ec5b30af3e14192c17ad00557e9d3f437bb5.png";

export const Hero = () => {
  return (
    <section className=" text-white flex justify-between ">
      <div className="max-h-112 flex flex-col md:flex-row items-center justify-between gap-10 bg-[#B0CC0D] p-32 rounded-3xl">
        <div className="">
          <h1 className="text-7xl font-bold mb-4 uppercase	">
            Your Personal Cook
          </h1>
          <p className="text-xl mb-6 text-black">
            The balanced diet for every day
          </p>
          <button className="bg-[#FFA800] hover:bg-[#ff9500] text-white font-semibold py-3 px-6 rounded-full ">
            Make your choice
          </button>
        </div>
        <img
          src={bread}
          alt="Avocado toast"
          className="max-h-160 translate-y-8"
        />
      </div>
    </section>
  );
};
