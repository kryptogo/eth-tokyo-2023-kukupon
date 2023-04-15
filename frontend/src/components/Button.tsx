import React from 'react';

interface IButton extends React.HTMLAttributes<HTMLButtonElement> {
  text: string;
  status?: 'normal' | 'disabled' | 'loading';
}

const Button: React.FC<IButton> = ({ text, status = 'normal', ...rest }) => {
  return (
    <button
      className="bg-primary flex items-center justify-center text-white text-[24px] rounded-[10px] py-2 px-[30px] shadow-primary w-fit"
      {...rest}
    >
      {text}
    </button>
  );
};

export default Button;
