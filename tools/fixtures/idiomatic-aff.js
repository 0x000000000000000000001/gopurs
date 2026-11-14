let calls = 0;
let started = false;
let stopped = false;

const canceler = (_reason, _onError, onSuccess) => onSuccess();

export const hello = name => (onError, onSuccess) => {
  calls++;
  if (name === 'error') onError(new Error('expected error'));
  else onSuccess(`hello ${name}`);
  return canceler;
};
export const notify = (_onError, onSuccess) => { onSuccess(); return canceler; };
export const noError = (_onError, onSuccess) => { onSuccess('plain result'); return canceler; };
export const fire = notify;
export const readCalls = () => calls;
export const label = () => 'effect without Unit';
export const runCallback = callback => () => callback('typed');
export const pending = (_onError, _onSuccess) => {
  started = true;
  return (_reason, _cancelError, onCanceled) => { stopped = true; onCanceled(); };
};
export const isStarted = () => started;
export const isStopped = () => stopped;
