import Lottie from 'react-lottie';
import animationData from '../logo.json';

export function LottieComponent() {
  const options = {
    animationData: animationData,
    loop: false, // set loop to false
  };
  return (
    <div>
      <Lottie options={options} height={187} width={400} />
    </div>
  );
}
