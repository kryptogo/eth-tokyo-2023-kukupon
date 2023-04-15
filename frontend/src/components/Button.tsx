import React from 'react';

interface IButton extends React.HTMLAttributes<HTMLButtonElement> {
  text: string;
  status?: 'normal' | 'disabled' | 'loading';
  size?: 'small' | 'medium' | 'large';
}

const Button: React.FC<IButton> = ({
  text,
  status = 'normal',
  size = 'medium',
  ...rest
}) => {
  return (
    <button
      className={`bg-primary flex items-center justify-center text-white text-[24px] rounded-[10px] py-2 px-[30px] shadow-primary w-fit ${
        status === 'disabled' ? 'opacity-50' : ''
      }
      ${size === 'small' ? 'text-[16px] py-1 px-[20px]' : ''}
      `}
      {...rest}
    >
      {text}
    </button>
  );
};

export default Button;
