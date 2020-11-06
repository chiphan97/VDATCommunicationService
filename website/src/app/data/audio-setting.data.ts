import {SoundSetting} from '../model/sound-setting.model';

const getListSound = (): Array<SoundSetting> => {
  const audios: Array<SoundSetting> = new Array<SoundSetting>();

  const audio01 = new SoundSetting('Bell positive', 'bells-positive-sound.mp3');
  audios.push(audio01);

  const audio02 = new SoundSetting('Bubble pop swoosh', 'bubble-pop-swoosh.mp3');
  audios.push(audio02);

  const audio03 = new SoundSetting('Casual easy multimedia', 'casual-easy-multimedia-alert.mp3');
  audios.push(audio03);

  const audio04 = new SoundSetting('Percussion accept confirm', 'percussion-accept-confirm-sound.mp3');
  audios.push(audio04);

  const audio05 = new SoundSetting('Pop ding notification', 'pop-ding-notification-effect.mp3');
  audios.push(audio05);

  const audio06 = new SoundSetting('Smartphone mobile alert', 'smartphone-mobile-alert.mp3');
  audios.push(audio06);

  const audio07 = new SoundSetting('Xylo mallet complete sound percussion', 'xylo-mallet-complete-sound-4-percussion.mp3');
  audios.push(audio07);

  return audios;
};

const SoundSettingData = {
  getListSound
};

export default SoundSettingData;
