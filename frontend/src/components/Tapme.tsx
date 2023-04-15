import Lottie from 'react-lottie';
import animationData from '../tap_me.json';

export default function TapME() {
  const options = {
    animationData: animationData,
    segments: [10, 20],
    loop: false,
  };
  return (
    <div className="yoyo">
      <Lottie options={options} height={70} width={70} />
    </div>
  );
}
