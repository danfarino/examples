import "preact/debug";
import * as Preact from "preact";
import { computed, signal, useSignal } from "@preact/signals";
import { debounceTime, map, NEVER, of, scan, switchMap } from "rxjs";
import {
  observableFromIntersectionObserver,
  observableFromSignal,
  useObservable,
  useObservableRef,
} from "./helpers";
import styled from "styled-components";
import { mousePos } from "mouse";

const count = signal(0);

const countMsg = computed(() => `Count is ${count}`);

const BoxDiv = styled.div<{ highlight?: boolean }>`
  border: 1px solid #ccc;
  margin: 1em 0em;
  padding: 1em;
  ${(p) => p.highlight && "background-color: LightGoldenRodYellow"}
`;

function Thing(props: { name: string; highlight?: boolean }) {
  const str = computed(() => `thing ${props.name}: count=${count}`);

  return (
    <BoxDiv highlight={props.highlight}>
      <div>{props.name}</div>
      <div>{str}</div>
    </BoxDiv>
  );
}

type ArrowProps = { radians: number };

const Arrow = styled.div.attrs<ArrowProps>((p: ArrowProps) => ({
  style: {
    transform: `rotate(${p.radians}rad)`,
  },
}))`
  display: table;
  margin-left: auto;
  margin-right: auto;
  font-size: 2em;
  &::before {
    content: "➤";
  }
`;

function App() {
  const delayed = useSignal("none");
  const mouseUpdates = useSignal(0);
  const radians = useSignal(0);
  const [arrowElemObs, arrowElemRef] = useObservableRef<Element>();
  const showArrow = useSignal(true);
  const arrowInViewport = useSignal("");

  useObservable(
    () => observableFromSignal(mousePos).pipe(scan((n) => n + 1, 0)),
    (n) => (mouseUpdates.value = n)
  );

  useObservable(
    () =>
      observableFromSignal(countMsg).pipe(
        debounceTime(1000),
        map((s) => `${s} at ${new Date()}`)
      ),
    (s) => (delayed.value = s)
  );

  useObservable(
    () =>
      arrowElemObs.pipe(
        switchMap((arrowElem) => {
          if (!arrowElem) {
            return NEVER;
          }

          return observableFromSignal(mousePos).pipe(
            map((p) => {
              const rect = arrowElem.getBoundingClientRect();
              const dx = p.x - (rect.width / 2 + rect.x);
              const dy = p.y - (rect.height / 2 + rect.y);

              return Math.atan2(dy, dx);
            })
          );
        })
      ),
    (r) => {
      radians.value = r;
    }
  );

  useObservable(
    () =>
      arrowElemObs.pipe(
        switchMap((arrowElem) =>
          arrowElem
            ? observableFromIntersectionObserver(arrowElem, { threshold: 0.5 })
            : of(null)
        )
      ),
    (e) => {
      if (e) {
        arrowInViewport.value = String(e.isIntersecting);
      } else {
        arrowInViewport.value = "(hidden)";
      }
    }
  );

  return (
    <div>
      <div>click count: {count}</div>
      <pre>{delayed}</pre>
      <div>
        <button onClick={() => count.value++}>Clicky</button>
      </div>
      <div>mouse pos: {JSON.stringify(mousePos.value)}</div>
      <div>mouse updates: {mouseUpdates}</div>
      <div>
        <label style={{ userSelect: "none" }}>
          <input
            type="checkbox"
            checked={showArrow.value}
            onInput={(e) => (showArrow.value = e.currentTarget.checked)}
          />
          Show arrow?
        </label>
      </div>
      <div>Arrow in viewport? {arrowInViewport}</div>
      <textarea
        style={{
          width: "100%",
          height: "4em",
          padding: "0.5em",
          boxSizing: "border-box",
        }}
        placeholder="(resize me to make the arrow leave the viewport)"
      ></textarea>
      {showArrow.value && <Arrow ref={arrowElemRef} radians={radians.value} />}
      <Thing name="foo" />
      <Thing name="bar" highlight={true} />
    </div>
  );
}

Preact.render(<App />, document.getElementById("root")!);
