import { Signal } from "@preact/signals";
import { Inputs, useEffect, useRef } from "preact/hooks";
import { Observable, Observer, ReplaySubject } from "rxjs";

export function useObservable<T>(
  getObservable: () => Observable<T>,
  observer: Partial<Observer<T>> | ((value: T) => void) | undefined,
  inputs?: Inputs
) {
  useEffect(() => {
    const sub = getObservable().subscribe(observer as any);
    return () => sub.unsubscribe();
  }, inputs || []);
}

export function useObservableRef<T>(): [Observable<T>, (elem: T) => void] {
  interface State {
    obs: Observable<T>;
    refFn: (elem: T) => void;
  }

  const state = useRef<State>();
  if (!state.current) {
    const subj = new ReplaySubject<T>(1);
    state.current = {
      obs: subj,
      refFn: (elem) => subj.next(elem),
    };
  }

  return [state.current.obs, state.current.refFn];
}

export function observableFromSignal<T>(sig: Signal<T>): Observable<T> {
  return new Observable<T>((subscriber) => {
    const unsubscribe = sig.subscribe((o) => subscriber.next(o));
    return unsubscribe;
  });
}

export function observableFromIntersectionObserver(
  target: Element,
  options?: IntersectionObserverInit
) {
  return new Observable<IntersectionObserverEntry>((subscriber) => {
    const intObs = new IntersectionObserver((events) => {
      for (const event of events) {
        subscriber.next(event);
      }
    }, options);

    intObs.observe(target);

    return () => intObs.disconnect();
  });
}
