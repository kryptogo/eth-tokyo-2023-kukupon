import React from "react";

import Button from "./Button";

interface IEventCard {
  title: string;
  description: string[];
  backgroundImage: string;
}

const EventCard: React.FC<IEventCard> = ({
  title,
  description,
  backgroundImage,
}) => {
  return (
    <div
      className="w-[280px] h-[394px] border border-white rounded-[10px] flex items-center justify-center"
      style={{
        backgroundImage: `url(${backgroundImage})`,
      }}
    >
      <div className="flex flex-col items-center gap-[120px]">
        <div className="flex flex-col items-center gap-5">
          <h2 className="text-[42px] font-bold">{title}</h2>
          <ul>
            {description.map((item, index) => (
              <li key={index}>{item}</li>
            ))}
          </ul>
        </div>
        <Button text="Verify" />
      </div>
    </div>
  );
};

export default EventCard;
